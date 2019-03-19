pragma solidity ^0.5.3;
import "./PermissionsUpgradable.sol";

contract AccountManager {
    PermissionsUpgradable private permUpgradable;
//    enum AccountStatus {0-NotInList, 1-PendingApproval, 2-Active, 3-Inactive}
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

    mapping(bytes32 => bool) private orgAdminIndex;

    // account permission events
    event AccountAccessModified(address _address, string _roleId);
    event AccountAccessRevoked(address _address, string _roleId);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    // Get account details given index

    function orgAdminExists(string memory _orgId) public view returns (bool)
    {
        return orgAdminIndex[keccak256(abi.encodePacked(_orgId))];

    }

    function getAccountStatus(address _acct) internal view returns (uint)
    {
        if (accountIndex[_acct] == 0) {
            return 0;
        }
        uint aIndex = getAcctIndex(_acct);
        return (acctAccessList[aIndex].status);
    }

    function getAccountDetails(address _acct) external view returns (address, string memory, string memory, uint, bool)
    {
        if (accountIndex[_acct] == 0) {
            return (_acct, "NONE", "", 0, false);
        }
        uint aIndex = getAcctIndex(_acct);
        return (acctAccessList[aIndex].acctId, acctAccessList[aIndex].orgId, acctAccessList[aIndex].role, acctAccessList[aIndex].status, acctAccessList[aIndex].orgAdmin);
    }

    // Get number of accounts
    function getNumberOfAccounts() external view returns (uint)
    {
        return acctAccessList.length;
    }

    function setDefaults(string calldata _nwAdminRole, string calldata _oAdminRole) external
    {
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
    }

    function setAccountRole(address _address, string memory _orgId, string memory _roleId, uint _status, bool _oAdmin) internal
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
        if (_oAdmin) {
            orgAdminIndex[keccak256(abi.encodePacked(_orgId))] = true;
        }
        emit AccountAccessModified(_address, _roleId);
    }

    function addNWAdminAccount(address _address, string calldata _orgId) external
    {
        setAccountRole(_address, _orgId, adminRole, 2, true);
    }

    function assignAccountRole(address _address, string calldata _orgId, string calldata _roleId) external
    {
        bool oAdminRole = false;
        uint status = 2;
        // if the role id is ORGADMIN then check if already an orgadmin exists
        if ((keccak256(abi.encodePacked(_roleId)) == keccak256(abi.encodePacked(orgAdminRole))) ||
            (keccak256(abi.encodePacked(_roleId)) == keccak256(abi.encodePacked(adminRole)))) {
            if (orgAdminIndex[keccak256(abi.encodePacked(_orgId))]) {
                return;
            }
            else {
                oAdminRole = true;
                status = 1;
            }
        }
        setAccountRole(_address, _orgId, _roleId, status, oAdminRole);
    }

    function approveOrgAdminAccount(address _address) external
    {
        // check of the account role is ORGADMIN and status is pending approval
        // if yes update the status to approved
        string memory role = getAccountRole(_address);
        uint status = getAccountStatus(_address);

        if ((keccak256(abi.encodePacked(role)) == keccak256(abi.encodePacked(orgAdminRole))) &&
            (status == 1)) {
            uint aIndex = getAcctIndex(_address);
            acctAccessList[aIndex].status = 2;
            emit AccountAccessModified(_address, acctAccessList[aIndex].role);
        }

    }

    function revokeAccountRole(address _address) external
    {
        // Check if account already exists
        uint aIndex = getAcctIndex(_address);
        if (accountIndex[_address] != 0) {
            acctAccessList[aIndex].status = 3;
            emit AccountAccessRevoked(_address, acctAccessList[aIndex].role);
        }
    }

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

    function checkOrgAdmin(address _acct, string calldata _orgId) external view returns (bool)
    {
        if (accountIndex[_acct] == 0) {
            return false;
        }
        uint acctIndex = getAcctIndex(_acct);
        return ((acctAccessList[acctIndex].orgAdmin) &&
                (keccak256(abi.encodePacked(acctAccessList[acctIndex].orgId)) == keccak256(abi.encodePacked(_orgId))));
    }

    // this function checks if account access can be modified. Account access can be modified for a new account
    // or if the call is from the orgadmin of the same org.
    function valAcctAccessChange(address _acct, string calldata _orgId) external view returns (bool)
    {
        if (accountIndex[_acct] == 0) {
            return true;
        }
        uint acctIndex = getAcctIndex(_acct);
        return ((keccak256(abi.encodePacked(acctAccessList[acctIndex].orgId)) == keccak256(abi.encodePacked(_orgId))));
    }
    // Returns the account index based on account id
    function getAcctIndex(address _acct) internal view returns (uint)
    {
        return accountIndex[_acct] - 1;
    }

}
