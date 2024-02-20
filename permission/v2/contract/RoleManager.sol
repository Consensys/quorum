pragma solidity ^0.8.17;

import "./PermissionsUpgradable.sol";
import "./openzeppelin-v5/Initializable.sol";
/** @title Role manager contract
  * @notice This contract holds implementation logic for all role management
    functionality. This can be called only by the implementation
    contract only. there are few view functions exposed as public and
    can be called directly. these are invoked by quorum for populating
    permissions data in cache
  */
contract RoleManager is Initializable {
    PermissionsUpgradable private permUpgradable;

    struct RoleDetails {
        string roleId;
        string orgId;
        uint256 baseAccess;
        bool isVoter;
        bool isAdmin;
        bool active;
    }

    RoleDetails[] private roleList;
    mapping(bytes32 => uint256) private roleIndex;
    uint256 private numberOfRoles;
    uint256 private defaultAccessForUnconfiguredAccounts;

    event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess,
        bool _isVoter, bool _isAdmin);
    event RoleRevoked(string _roleId, string _orgId);

    /** notice: confirms that the caller is the address of implementation
         contract
     */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    // @notice initialized only once. sets the permissions upgradable address
    function initialize(address _permUpgradable) public initializer {
        require(_permUpgradable != address(0x0), "Cannot set to empty address");
        
        permUpgradable = PermissionsUpgradable(_permUpgradable);
        defaultAccessForUnconfiguredAccounts = 5;
    }  

    /** @notice function to add a new role definition to an organization
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @param _baseAccess - can be from 0 to 7
      * @param _isVoter - bool to indicate if voter role or not
      * @param _isAdmin - bool to indicate if admin role or not
      * @dev base access can have any of the following values:
            0 - Read only
            1 - value transfer
            2 - contract deploy
            3 - full access
            4 - contract call
            5 - value transfer and contract call
            6 - value transfer and contract deploy
            7 - contract call and deploy
      */
    function addRole(string memory _roleId, string memory _orgId, uint256 _baseAccess,
        bool _isVoter, bool _isAdmin) public onlyImplementation {
        require(_baseAccess < 8, "invalid access value");
        // Check if account already exists
        require(roleIndex[keccak256(abi.encode(_roleId, _orgId))] == 0, "role exists for the org");
        numberOfRoles ++;
        roleIndex[keccak256(abi.encode(_roleId, _orgId))] = numberOfRoles;
        roleList.push(RoleDetails(_roleId, _orgId, _baseAccess, _isVoter, _isAdmin, true));
        emit RoleCreated(_roleId, _orgId, _baseAccess, _isVoter, _isAdmin);
    }

    /** @notice function to remove an existing role definition from an organization
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      */
    function removeRole(string calldata _roleId, string calldata _orgId) external
    onlyImplementation {
        require(roleIndex[keccak256(abi.encode(_roleId, _orgId))] != 0, "role does not exist");
        uint256 rIndex = _getRoleIndex(_roleId, _orgId);
        roleList[rIndex].active = false;
        emit RoleRevoked(_roleId, _orgId);
    }

    /** @notice checks if the role is a voter role or not
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @param _ultParent - master org id
      * @return true or false
      * @dev checks for the role existence in the passed org and master org
      */
    function isVoterRole(string calldata _roleId, string calldata _orgId,
        string calldata _ultParent) external view onlyImplementation returns (bool){
        if (!(roleExists(_roleId, _orgId, _ultParent))) {
            return false;
        }
        uint256 rIndex;
        if (roleIndex[keccak256(abi.encode(_roleId, _orgId))] != 0) {
            rIndex = _getRoleIndex(_roleId, _orgId);
        }
        else {
            rIndex = _getRoleIndex(_roleId, _ultParent);
        }
        return (roleList[rIndex].active && roleList[rIndex].isVoter);
    }

    /** @notice checks if the role is an admin role or not
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @param _ultParent - master org id
      * @return true or false
      * @dev checks for the role existence in the passed org and master org
      */
    function isAdminRole(string calldata _roleId, string calldata _orgId,
        string calldata _ultParent) external view onlyImplementation returns (bool){
        if (!(roleExists(_roleId, _orgId, _ultParent))) {
            return false;
        }
        uint256 rIndex;
        if (roleIndex[keccak256(abi.encode(_roleId, _orgId))] != 0) {
            rIndex = _getRoleIndex(_roleId, _orgId);
        }
        else {
            rIndex = _getRoleIndex(_roleId, _ultParent);
        }
        return (roleList[rIndex].active && roleList[rIndex].isAdmin);
    }

    /** @notice returns the role details for a passed role id and org
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @return roleId
      * @return orgId
      * @return accessType
      * @return voter - bool to indicate if the role is a voter role
      * @return admin
      * @return active - bool to indicate if the role is active
      */
    function getRoleDetails(string calldata _roleId, string calldata _orgId)
    external view returns (string memory roleId, string memory orgId,
        uint256 accessType, bool voter, bool admin, bool active) {
        if (!(roleExists(_roleId, _orgId, ""))) {
            return (_roleId, "", 0, false, false, false);
        }
        uint256 rIndex = _getRoleIndex(_roleId, _orgId);
        return (roleList[rIndex].roleId, roleList[rIndex].orgId,
        roleList[rIndex].baseAccess, roleList[rIndex].isVoter,
        roleList[rIndex].isAdmin, roleList[rIndex].active);
    }

    /** @notice returns the role details for a passed role index
      * @param _rIndex - unique identifier for the role being added
      * @return roleId
      * @return orgId
      * @return accessType
      * @return voter - bool to indicate if the role is a voter role
      * @return admin
      * @return active - bool to indicate if the role is active
      */
    function getRoleDetailsFromIndex(uint256 _rIndex) external view returns
    (string memory roleId, string memory orgId, uint256 accessType,
        bool voter, bool admin, bool active) {
        return (roleList[_rIndex].roleId, roleList[_rIndex].orgId,
        roleList[_rIndex].baseAccess, roleList[_rIndex].isVoter,
        roleList[_rIndex].isAdmin, roleList[_rIndex].active);
    }

    /** @notice returns the total number of roles in the network
      * @return total number of roles
      */
    function getNumberOfRoles() external view returns (uint256) {
        return roleList.length;
    }

    /** @notice checks if the role exists for the given org or master org
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @param _ultParent - master org id
      * @return true or false
      */
    function roleExists(string memory _roleId, string memory _orgId,
        string memory _ultParent) public view returns (bool) {
        uint256 id;
        if (roleIndex[keccak256(abi.encode(_roleId, _orgId))] != 0) {
            id = _getRoleIndex(_roleId, _orgId);
            return roleList[id].active;
        }
        else if (roleIndex[keccak256(abi.encode(_roleId, _ultParent))] != 0) {
            id = _getRoleIndex(_roleId, _ultParent);
            return roleList[id].active;
        }
        return false;
    }

    function roleAccess(string memory _roleId, string memory _orgId,
        string memory _ultParent) public view returns (uint256) {
        uint256 id;
        if (roleIndex[keccak256(abi.encode(_roleId, _orgId))] != 0) {
            id = _getRoleIndex(_roleId, _orgId);
            return roleList[id].baseAccess;
        }
        else if (roleIndex[keccak256(abi.encode(_roleId, _ultParent))] != 0) {
            id = _getRoleIndex(_roleId, _ultParent);
            return roleList[id].baseAccess;
        }
        return 0;
    }

    function transactionAllowed(string calldata _roleId, string calldata _orgId,
        string calldata _ultParent, uint256 _typeOfTxn) external view returns (bool) {
        uint256 access = roleAccess(_roleId, _orgId, _ultParent);

        return isTransactionAllowedBasedOnRoleAccess(access, _typeOfTxn);
    }

    function isTransactionAllowedBasedOnRoleAccess(uint256 access, uint256 _typeOfTxn) public pure returns (bool) {

        if (access == 3) {
            return true;
        }

        /** typeOfTxn
            1 - value transfer
            2 - contract deploy
            3 - contract call **/

        if (_typeOfTxn == 1 && (access == 1 || access == 5 || access == 6)){
            return true;
        }
        if (_typeOfTxn == 2 && (access == 2 || access == 6 || access == 7)){
            return true;
        }
        if (_typeOfTxn == 3 && (access == 4 || access == 5 || access == 7)){
            return true;
        }

        return false;
    }

    /** @notice returns the role index based on role id and org id
      * @param _roleId - role id
      * @param _orgId - org id
      * @return role index
      */
    function _getRoleIndex(string memory _roleId, string memory _orgId)
    internal view returns (uint256) {
        uint256 _roleIndex = roleIndex[keccak256(abi.encode(_roleId, _orgId))];
        if (_roleIndex == 0){
            return type(uint256).max;
        }
        return _roleIndex - 1;
    }

    /** @notice function to set the default access level for unconfigured account. 
            Unconfigured account does not have role and org membership but is assigned
            a default access level of 5 (transfer value and/or call contract) 
      * @param _accessLevel - set the default access level for unconfigured account.
      */
    function setAccessLevelForUnconfiguredAccount(uint256 _accessLevel) external 
        onlyImplementation
    {
        require(_accessLevel >= 0 && _accessLevel <= 7, "accessLevel value should be between 0 to 7");
        defaultAccessForUnconfiguredAccounts = _accessLevel;
    }

    /** @notice get the default access level for unconfigured account. */
    function getAccessLevelForUnconfiguredAccount() external view returns (uint256)
    {
        return defaultAccessForUnconfiguredAccounts;
    }
}
