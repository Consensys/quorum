pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";
/** @title Role manager contract
  * @notice This contract holds implementation logic for all role management
    functionality. This can be called only by the implementation
    contract only. there are few view functions exposed as public and
    can be called directly. these are invoked by quorum for populating
    permissions data in cache
  */
contract RoleManager {
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

    event RoleCreated(string _roleId, string _orgId, uint256 _baseAccess,
        bool _isVoter, bool _isAdmin);
    event RoleRevoked(string _roleId, string _orgId);

    /** @notice confirms that the caller is the address of implementation
         contract
     */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /** @notice constructor. sets the permissions upgradable address
      */
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice function to add a new role definition to an organization
      * @param _roleId - unique identifier for the role being added
      * @param _orgId - org id to which the role belongs
      * @param _baseAccess - 0-ReadOnly, 1-Transact, 2-ContractDeply, 3- Full
      * @param _isVoter - bool to indicate if voter role or not
      * @param _isAdmin - bool to indicate if admin role or not
      * @dev base access can have any of the following values:
            0 - Read only
            1 - Transact only
            2 - Contract deploy. can transact as well
            3 - Full access
      */
    function addRole(string memory _roleId, string memory _orgId, uint256 _baseAccess,
        bool _isVoter, bool _isAdmin) public onlyImplementation {
        // check if the role access passed is valid
        require(_baseAccess < 4, "invalid access value");
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
      * @return role id
      * @return org id
      * @return access type
      * @return bool to indicate if the role is a voter role
      * @return bool to indicate if the role is active
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
      * @return role id
      * @return org id
      * @return access type
      * @return bool to indicate if the role is a voter role
      * @return bool to indicate if the role is active
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

    /** @notice returns the role index based on role id and org id
      * @param _roleId - role id
      * @param _orgId - org id
      * @return role index
      */
    function _getRoleIndex(string memory _roleId, string memory _orgId)
    internal view returns (uint256) {
        return roleIndex[keccak256(abi.encode(_roleId, _orgId))] - 1;
    }
}
