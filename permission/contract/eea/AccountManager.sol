pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

/** @title Account manager contract
  * @notice This contract holds implementation logic for all account management
    functionality. This can be called only by the implementation contract only.
    there are few view functions exposed as public and can be called directly.
    these are invoked by quorum for populating permissions data in cache
  * @dev account status is denoted by a fixed integer value. The values are
    as below:
        0 - Not in list
        1 - Account pending approval
        2 - Active
        3 - Inactive
        4 - Suspended
        5 - Blacklisted
        6 - Revoked
        7 - Recovery Initiated for blacklisted accounts and pending approval
            from network admins
     Once the account is blacklisted no further activity on the account is
     possible.
     When adding a new org admin account to an existing org, the existing org
     admin account will be in revoked status and can be assigned a new role
     later
  */
contract AccountManager {
    PermissionsUpgradable private permUpgradable;
    struct AccountAccessDetails {
        address account;
        string orgId;
        string role;
        uint status;
        bool orgAdmin;
    }

    AccountAccessDetails[] private accountAccessList;
    mapping(address => uint) private accountIndex;
    uint private numAccounts;

    string private adminRole;
    string private orgAdminRole;

    mapping(bytes32 => address) private orgAdminIndex;

    // account permission events
    event AccountAccessModified(address _account, string _orgId, string _roleId, bool _orgAdmin, uint _status);
    event AccountAccessRevoked(address _account, string _orgId, string _roleId, bool _orgAdmin);
    event AccountStatusChanged(address _account, string _orgId, uint _status);

    /** @notice confirms that the caller is the address of implementation
        contract
      */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /** @notice checks if the account is exists and belongs to the org id passed
      * @param _orgId - org id
      * @param _account - account id
      */
    modifier accountExists(string memory _orgId, address _account) {
        require((accountIndex[_account]) != 0, "account does not exists");
        require(keccak256(abi.encode(accountAccessList[_getAccountIndex(_account)].orgId)) == keccak256(abi.encode(_orgId)), "account in different org");
        _;
    }

    /// @notice constructor. sets the permissions upgradable address
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }


    /** @notice returns the account details for a given account
      * @param _account account id
      * @return account id
      * @return org id of the account
      * @return role linked to the account
      * @return status of the account
      * @return bool indicating if the account is an org admin
      */
    function getAccountDetails(address _account) external view returns (address,
        string memory, string memory, uint, bool){
        if (accountIndex[_account] == 0) {
            return (_account, "NONE", "", 0, false);
        }
        uint aIndex = _getAccountIndex(_account);
        return (accountAccessList[aIndex].account, accountAccessList[aIndex].orgId,
        accountAccessList[aIndex].role, accountAccessList[aIndex].status,
        accountAccessList[aIndex].orgAdmin);
    }

    /** @notice returns the account details for a given account if account is valid/active
      * @param _account account id
      * @return org id of the account
      * @return role linked to the account
      */
    function getAccountOrgRole(address _account) external view
    returns (string memory, string memory){
        if (accountIndex[_account] == 0) {
            return ("NONE", "");
        }
        uint aIndex = _getAccountIndex(_account);
        return (accountAccessList[aIndex].orgId, accountAccessList[aIndex].role);
    }

    /** @notice returns the account details a given account index
      * @param  _aIndex account index
      * @return account id
      * @return org id of the account
      * @return role linked to the account
      * @return status of the account
      * @return bool indicating if the account is an org admin
      */
    function getAccountDetailsFromIndex(uint _aIndex) external view returns
    (address, string memory, string memory, uint, bool) {
        return (accountAccessList[_aIndex].account,
        accountAccessList[_aIndex].orgId, accountAccessList[_aIndex].role,
        accountAccessList[_aIndex].status, accountAccessList[_aIndex].orgAdmin);
    }

    /** @notice returns the total number of accounts
      * @return total number accounts
      */
    function getNumberOfAccounts() external view returns (uint) {
        return accountAccessList.length;
    }

    /** @notice this is called at the time of network initialization to set
        the default values of network admin and org admin roles
      */
    function setDefaults(string calldata _nwAdminRole, string calldata _oAdminRole)
    external onlyImplementation {
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
    }

    /** @notice this function is called to assign the org admin or network
        admin roles only to the passed account
      * @param _account - account id
      * @param _orgId - org to which it belongs
      * @param _roleId - role id to be assigned
      * @param _status - account status to be assigned
      */
    function assignAdminRole(address _account, string calldata _orgId,
        string calldata _roleId, uint _status) external onlyImplementation {
        require(((keccak256(abi.encode(_roleId)) == keccak256(abi.encode(orgAdminRole))) ||
        (keccak256(abi.encode(_roleId)) == keccak256(abi.encode(adminRole)))),
            "can be called to assign admin roles only");

        _setAccountRole(_account, _orgId, _roleId, _status, true);

    }

    /** @notice this function is called to assign the any role to the passed
        account.
      * @param _account - account id
      * @param _orgId - org to which it belongs
      * @param _roleId - role id to be assigned
      * @param _adminRole - indicates of the role is an admin role
      */
    function assignAccountRole(address _account, string calldata _orgId,
        string calldata _roleId, bool _adminRole) external onlyImplementation {
        require(((keccak256(abi.encode(_roleId)) != keccak256(abi.encode(adminRole)))
        && (keccak256(abi.encode(abi.encode(_roleId))) != keccak256(abi.encode(orgAdminRole)))),
            "cannot be called fro assigning org admin and network admin roles");
        _setAccountRole(_account, _orgId, _roleId, 2, _adminRole);
    }

    /** @notice this function removes existing admin account. will be called at
        the time of adding a new account as org admin account. at org
        level there can be one org admin account only
      * @param _orgId - org id
      * @return bool to indicate if voter update is required or not
      * @return _adminRole - indicates of the role is an admin role
      */
    function removeExistingAdmin(string calldata _orgId) external
    onlyImplementation
    returns (bool voterUpdate, address account) {
        // change the status of existing org admin to revoked
        if (orgAdminExists(_orgId)) {
            uint id = _getAccountIndex(orgAdminIndex[keccak256(abi.encode(_orgId))]);
            accountAccessList[id].status = 6;
            accountAccessList[id].orgAdmin = false;
            emit AccountAccessModified(accountAccessList[id].account,
                accountAccessList[id].orgId, accountAccessList[id].role,
                accountAccessList[id].orgAdmin, accountAccessList[id].status);
            return ((keccak256(abi.encode(accountAccessList[id].role)) == keccak256(abi.encode(adminRole))),
            accountAccessList[id].account);
        }
        return (false, address(0));
    }

    /** @notice function to add an account as network admin or org admin.
      * @param _orgId - org id
      * @param _account - account id
      * @return bool to indicate if voter update is required or not
      */
    function addNewAdmin(string calldata _orgId, address _account) external
    onlyImplementation
    returns (bool voterUpdate) {
        // check of the account role is org admin role and status is pending
        // approval. if yes update the status to approved
        string memory role = getAccountRole(_account);
        uint status = getAccountStatus(_account);
        uint id = _getAccountIndex(_account);
        if ((keccak256(abi.encode(role)) == keccak256(abi.encode(orgAdminRole))) &&
            (status == 1)) {
            orgAdminIndex[keccak256(abi.encode(_orgId))] = _account;
        }
        accountAccessList[id].status = 2;
        accountAccessList[id].orgAdmin = true;
        emit AccountAccessModified(_account, accountAccessList[id].orgId, accountAccessList[id].role,
            accountAccessList[id].orgAdmin, accountAccessList[id].status);
        return (keccak256(abi.encode(accountAccessList[id].role)) == keccak256(abi.encode(adminRole)));
    }

    /** @notice updates the account status to the passed status value
      * @param _orgId - org id
      * @param _account - account id
      * @param _action - new status of the account
      * @dev the following actions are allowed
            1 - Suspend the account
            2 - Reactivate a suspended account
            3 - Blacklist an account
            4 - Initiate recovery for black listed account
            5 - Complete recovery of black listed account and update status to active
      */
    function updateAccountStatus(string calldata _orgId, address _account, uint _action) external
    onlyImplementation
    accountExists(_orgId, _account) {
        require((_action > 0 && _action < 6), "invalid status change request");

        // check if the account is org admin. if yes then do not allow any status change
        require(checkOrgAdmin(_account, _orgId, "") != true, "status change not possible for org admin accounts");
        uint newStatus;
        if (_action == 1) {
            // for suspending an account current status should be active
            require(accountAccessList[_getAccountIndex(_account)].status == 2,
                "account is not in active status. operation cannot be done");
            newStatus = 4;
        }
        else if (_action == 2) {
            // for reactivating a suspended account, current status should be suspended
            require(accountAccessList[_getAccountIndex(_account)].status == 4,
                "account is not in suspended status. operation cannot be done");
            newStatus = 2;
        }
        else if (_action == 3) {
            require(accountAccessList[_getAccountIndex(_account)].status != 5,
                "account is already blacklisted. operation cannot be done");
            newStatus = 5;
        }
        else if (_action == 4) {
            require(accountAccessList[_getAccountIndex(_account)].status == 5,
                "account is not blacklisted. operation cannot be done");
            newStatus = 7;
        }
        else if (_action == 5) {
            require(accountAccessList[_getAccountIndex(_account)].status == 7, "account recovery not initiated. operation cannot be done");
            newStatus = 2;
        }

        accountAccessList[_getAccountIndex(_account)].status = newStatus;
        emit AccountStatusChanged(_account, _orgId, newStatus);
    }

    /** @notice checks if the passed account exists and if exists does it
        belong to the passed organization.
      * @param _account - account id
      * @param _orgId - org id
      * @return bool true if the account does not exists or exists and belongs
      * @return passed org
      */
    function validateAccount(address _account, string calldata _orgId) external
    view returns (bool){
        if (accountIndex[_account] == 0) {
            return true;
        }
        uint256 id = _getAccountIndex(_account);
        return (keccak256(abi.encode(accountAccessList[id].orgId)) == keccak256(abi.encode(_orgId)));
    }

    /** @notice checks if org admin account exists for the passed org id
      * @param _orgId - org id
      * @return true if the org admin account exists and is approved
      */
    function orgAdminExists(string memory _orgId) public view returns (bool) {
        if (orgAdminIndex[keccak256(abi.encode(_orgId))] != address(0)) {
            address adminAcct = orgAdminIndex[keccak256(abi.encode(_orgId))];
            return getAccountStatus(adminAcct) == 2;
        }
        return false;

    }

    /** @notice returns the role id linked to the passed account
      * @param _account account id
      * @return role id
      */
    function getAccountRole(address _account) public view returns (string memory) {
        if (accountIndex[_account] == 0) {
            return "NONE";
        }
        uint256 acctIndex = _getAccountIndex(_account);
        if (accountAccessList[acctIndex].status != 0) {
            return accountAccessList[acctIndex].role;
        }
        else {
            return "NONE";
        }
    }

    /** @notice returns the account status for a given account
      * @param _account account id
      * @return account status
      */
    function getAccountStatus(address _account) public view returns (uint256) {
        if (accountIndex[_account] == 0) {
            return 0;
        }
        uint256 aIndex = _getAccountIndex(_account);
        return (accountAccessList[aIndex].status);
    }


    /** @notice checks if the account is a org admin for the passed org or
        for the ultimate parent organization
      * @param _account account id
      * @param _orgId org id
      * @param _ultParent master org id or
      */
    function checkOrgAdmin(address _account, string memory _orgId,
        string memory _ultParent) public view returns (bool) {
        // check if the account role is network admin. If yes return success
        if (keccak256(abi.encode(getAccountRole(_account))) == keccak256(abi.encode(adminRole))) {
            // check of the orgid is network admin org. then return true
            uint256 id = _getAccountIndex(_account);
            return ((keccak256(abi.encode(accountAccessList[id].orgId)) == keccak256(abi.encode(_orgId)))
            || (keccak256(abi.encode(accountAccessList[id].orgId)) == keccak256(abi.encode(_ultParent))));
        }
        return ((orgAdminIndex[keccak256(abi.encode(_orgId))] == _account) || (orgAdminIndex[keccak256(abi.encode(_ultParent))] == _account));
    }

    /** @notice returns the index for a given account id
      * @param _account account id
      * @return account index
      */
    function _getAccountIndex(address _account) internal view returns (uint256) {
        return accountIndex[_account] - 1;
    }

    /** @notice sets the account role to the passed role id and sets the status
      * @param _account account id
      * @param _orgId org id
      * @param _status status to be set
      * @param _oAdmin bool to indicate if account is org admin
      */
    function _setAccountRole(address _account, string memory _orgId,
        string memory _roleId, uint256 _status, bool _oAdmin) internal onlyImplementation {
        // Check if account already exists
        uint256 aIndex = _getAccountIndex(_account);
        if (accountIndex[_account] != 0) {
            accountAccessList[aIndex].role = _roleId;
            accountAccessList[aIndex].status = _status;
            accountAccessList[aIndex].orgAdmin = _oAdmin;
        }
        else {
            numAccounts ++;
            accountIndex[_account] = numAccounts;
            accountAccessList.push(AccountAccessDetails(_account, _orgId,
                _roleId, _status, _oAdmin));
        }
        emit AccountAccessModified(_account, _orgId, _roleId, _oAdmin, _status);
    }
}
