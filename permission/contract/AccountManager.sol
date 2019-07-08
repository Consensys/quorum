pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

contract AccountManager {
    PermissionsUpgradable private permUpgradable;
    //    enum AccountStatus {0-NotInList, 1-PendingApproval, 2-Active, 3-Inactive, 4-Suspended, 5-Blacklisted, 6-Revoked}
    struct AccountAccessDetails {
        address acctId;
        string orgId;
        string role;
        uint status;
        bool orgAdmin;
    }

    AccountAccessDetails[] private acctAccessList;
    mapping(address => uint) private accountIndex;
    uint private numberOfAccts;

    string private adminRole;
    string private orgAdminRole;

    mapping(bytes32 => address) private orgAdminIndex;

    // account permission events
    event AccountAccessModified(address _address, string _orgId, string _roleId, bool _orgAdmin, uint _status);
    event AccountAccessRevoked(address _address, string _orgId, string _roleId, bool _orgAdmin);
    event AccountStatusChanged(address _address, string _orgId, uint _status);

    // checks if the caller is implementation contracts
    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    // checks if the account is existing and part of the org
    modifier accountExists(string memory _orgId, address _account)
    {
        require((accountIndex[_account]) != 0, "account does not exists");
        // if account exists it should belong to the same orgAdminIndex
        require(keccak256(abi.encodePacked(acctAccessList[getAcctIndex(_account)].orgId)) == keccak256(abi.encodePacked(_orgId)), "account in different org");
        _;
    }

    // constructor. sets the upgradable address
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    // checks if the org is already having an org admin account
    function orgAdminExists(string memory _orgId) public view returns (bool)
    {
        if (orgAdminIndex[keccak256(abi.encodePacked(_orgId))] != address(0)) {
            address adminAcct = orgAdminIndex[keccak256(abi.encodePacked(_orgId))];
            return getAccountStatus(adminAcct) == 2;
        }
        return false;

    }

    // returns the status of input account. Returns 0 if the account is not
    // existing
    function getAccountStatus(address _acct) internal view returns (uint)
    {
        if (accountIndex[_acct] == 0) {
            return 0;
        }
        uint aIndex = getAcctIndex(_acct);
        return (acctAccessList[aIndex].status);
    }

    // Gets account details for a given account
    function getAccountDetails(address _acct) external view returns (address, string memory, string memory, uint, bool)
    {
        if (accountIndex[_acct] == 0) {
            return (_acct, "NONE", "", 0, false);
        }
        uint aIndex = getAcctIndex(_acct);
        return (acctAccessList[aIndex].acctId, acctAccessList[aIndex].orgId, acctAccessList[aIndex].role, acctAccessList[aIndex].status, acctAccessList[aIndex].orgAdmin);
    }

    // Gets account details given index
    function getAccountDetailsFromIndex(uint aIndex) external view returns (address, string memory, string memory, uint, bool)
    {
        return (acctAccessList[aIndex].acctId, acctAccessList[aIndex].orgId, acctAccessList[aIndex].role, acctAccessList[aIndex].status, acctAccessList[aIndex].orgAdmin);
    }

    // Get number of accounts
    function getNumberOfAccounts() external view returns (uint)
    {
        return acctAccessList.length;
    }

    // sets the default values for network admin and org admin roles
    function setDefaults(string calldata _nwAdminRole, string calldata _oAdminRole) external
    onlyImpl
    {
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
    }

    // associates an account with a role and organization
    function setAccountRole(address _address, string memory _orgId, string memory _roleId, uint _status, bool _oAdmin) internal
    onlyImpl
    {
        // Check if account already exists
        uint aIndex = getAcctIndex(_address);
        if (accountIndex[_address] != 0) {
            acctAccessList[aIndex].role = _roleId;
            acctAccessList[aIndex].status = _status;
            acctAccessList[aIndex].orgAdmin = _oAdmin;
        }
        else {
            numberOfAccts ++;
            accountIndex[_address] = numberOfAccts;
            acctAccessList.push(AccountAccessDetails(_address, _orgId, _roleId, _status, _oAdmin));
        }
        emit AccountAccessModified(_address, _orgId, _roleId, _oAdmin, _status);
    }

    // this function can be only called for assigning org admin to network amdin roles and can be invoked by
    // network admins only
    function assignAdminRole(address _address, string calldata _orgId, string calldata _roleId, uint _status) external
    onlyImpl
    {
        require(((keccak256(abi.encodePacked(_roleId)) == keccak256(abi.encodePacked(orgAdminRole))) ||
        (keccak256(abi.encodePacked(_roleId)) == keccak256(abi.encodePacked(adminRole)))), "can be called to assign admin roles only");

        setAccountRole(_address, _orgId, _roleId, _status, true);

    }

    // this function can be only called for assigning any roles to accounts can be called by
    // org admins only
    function assignAccountRole(address _address, string calldata _orgId, string calldata _roleId, bool _adminRole) external
    onlyImpl
    {
        require(((keccak256(abi.encodePacked(_roleId)) != keccak256(abi.encodePacked(adminRole))) && (keccak256(abi.encodePacked(abi.encodePacked(_roleId))) != keccak256(abi.encodePacked(orgAdminRole)))), "cannot be called fro assigning org admin and network admin roles");
        setAccountRole(_address, _orgId, _roleId, 2, _adminRole);
    }

    // this function removes an existing org admin from the admin role
    function removeExistingAdmin(string calldata _orgId) external
    onlyImpl
    returns (bool voterUpdate, address acct)
    {
        // change the status of existing org admin to revoked
        if (orgAdminExists(_orgId)) {
            uint id = getAcctIndex(orgAdminIndex[keccak256(abi.encodePacked(_orgId))]);
            acctAccessList[id].status = 6;
            acctAccessList[id].orgAdmin = false;
            emit AccountAccessModified(acctAccessList[id].acctId, acctAccessList[id].orgId, acctAccessList[id].role, acctAccessList[id].orgAdmin, acctAccessList[id].status);
            return ((keccak256(abi.encodePacked(acctAccessList[id].role)) == keccak256(abi.encodePacked(adminRole))), acctAccessList[id].acctId);
        }
        return (false, address(0));
    }

    // this function associates a new account with org or network admin role
    function addNewAdmin(string calldata _orgId, address _address) external
    onlyImpl
    returns (bool voterUpdate)
    {
        // check of the account role is ORGADMIN and status is pending approval
        // if yes update the status to approved
        string memory role = getAccountRole(_address);
        uint status = getAccountStatus(_address);
        uint id = getAcctIndex(_address);
        if ((keccak256(abi.encodePacked(role)) == keccak256(abi.encodePacked(orgAdminRole))) &&
            (status == 1)) {
            orgAdminIndex[keccak256(abi.encodePacked(_orgId))] = _address;
        }
        acctAccessList[id].status = 2;
        acctAccessList[id].orgAdmin = true;
        emit AccountAccessModified(_address, acctAccessList[id].orgId, acctAccessList[id].role, acctAccessList[id].orgAdmin, acctAccessList[id].status);
        return (keccak256(abi.encodePacked(acctAccessList[id].role)) == keccak256(abi.encodePacked(adminRole)));
    }

    // this function can be called for updating the account status suspending or blaclisting an account
    // and for revoking suspension of an account
    function updateAccountStatus(string calldata _orgId, address _account, uint _status) external
    onlyImpl
    accountExists(_orgId, _account)
    {
        // changing node status to integer (0-NotInList, 1-PendingApproval, 2-Active, 3-Suspended, 4-Blacklisted, 5-Revoked)
        // operations that can be done 1-Suspend account, 2-Unsuspend Account, 3-Blacklist account
        require((_status == 1 || _status == 2 || _status == 3), "invalid operation");
        // check if the account is org admin. if yes then do not allow any status change
        require(checkOrgAdmin(_account, _orgId, "") != true, "cannot perform the operation on org admin account");
        uint newStat;
        if (_status == 1) {
            // account current status should be active
            require(acctAccessList[getAcctIndex(_account)].status == 2, "account should be active");
            newStat = 4;
        }
        else if (_status == 2) {
            // account current status should be suspended
            require(acctAccessList[getAcctIndex(_account)].status == 4, "account should be suspended");
            newStat = 2;
        }
        else if (_status == 3) {
            require(acctAccessList[getAcctIndex(_account)].status != 5, "account already blacklisted");
            newStat = 5;
        }
        acctAccessList[getAcctIndex(_account)].status = newStat;
        emit AccountStatusChanged(_account, _orgId, newStat);
    }

    // returns the account role
    function getAccountRole(address _acct) public view returns (string memory)
    {
        if (accountIndex[_acct] == 0) {
            return "NONE";
        }
        uint acctIndex = getAcctIndex(_acct);
        if (acctAccessList[acctIndex].status != 0) {
            return acctAccessList[acctIndex].role;
        }
        else {
            return "NONE";
        }
    }

    // checks if the account is a org admin for the passed organization or for the ultimate
    // parent organization
    function checkOrgAdmin(address _acct, string memory _orgId, string memory _ultParent) public view returns (bool)
    {
        // check if the account role is network admin. If yes return success
        if (keccak256(abi.encodePacked(getAccountRole(_acct))) == keccak256(abi.encodePacked(adminRole))) {
            // check of the orgid is network admin org. then return true
            uint id = getAcctIndex(_acct);
            return ((keccak256(abi.encodePacked(acctAccessList[id].orgId)) == keccak256(abi.encodePacked(_orgId)))
            || (keccak256(abi.encodePacked(acctAccessList[id].orgId)) == keccak256(abi.encodePacked(_ultParent))));
        }
        return ((orgAdminIndex[keccak256(abi.encodePacked(_orgId))] == _acct) || (orgAdminIndex[keccak256(abi.encodePacked(_ultParent))] == _acct));
    }

    // this function checks if account access can be modified. Account access can be modified for a new account
    // or if the call is from the orgadmin of the same org.
    function validateAccount(address _acct, string calldata _orgId) external view returns (bool)
    {
        if (accountIndex[_acct] == 0) {
            return true;
        }
        // check if the acount is part of this org else return false
        uint id = getAcctIndex(_acct);
        return (keccak256(abi.encodePacked(acctAccessList[id].orgId)) == keccak256(abi.encodePacked(_orgId)));
    }
    // Returns the account index based on account id
    function getAcctIndex(address _acct) internal view returns (uint)
    {
        return accountIndex[_acct] - 1;
    }

}
