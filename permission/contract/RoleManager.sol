pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

// TODO: check code comments
contract RoleManager {
    PermissionsUpgradable private permUpgradable;

    struct RoleDetails {
        string roleId;
        string orgId;
        uint baseAccess;
        bool isVoter;
        bool isAdmin;
        bool active;
    }

    RoleDetails[] private roleList;
    mapping(bytes32 => uint) private roleIndex;
    uint private numberOfRoles;

    event RoleCreated(string _roleId, string _orgId, uint _baseAccess, bool _isVoter, bool _isAdmin);
    event RoleRevoked(string _roleId, string _orgId);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    function roleExists(string memory _roleId, string memory _orgId, string memory _ultParent) public view returns (bool)
    {
        uint id;
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0) {
            id = getRoleIndex(_roleId, _orgId);
            return roleList[id].active;
        }
        else if (roleIndex[keccak256(abi.encodePacked(_roleId, _ultParent))] != 0) {
            id = getRoleIndex(_roleId, _ultParent);
            return roleList[id].active;
        }
        return false;
    }

    function getRoleDetails(string calldata _roleId, string calldata _orgId) external view returns (string memory roleId, string memory orgId, uint accessType, bool voter, bool active)
    {
        if (!(roleExists(_roleId, _orgId, ""))) {
            return (_roleId, "", 0, false, false);
        }
        uint rIndex = getRoleIndex(_roleId, _orgId);
        return (roleList[rIndex].roleId, roleList[rIndex].orgId, roleList[rIndex].baseAccess, roleList[rIndex].isVoter, roleList[rIndex].active);
    }

    function getRoleDetailsFromIndex(uint rIndex) external view returns (string memory roleId, string memory orgId, uint accessType, bool voter, bool admin, bool active)
    {
        return (roleList[rIndex].roleId, roleList[rIndex].orgId, roleList[rIndex].baseAccess, roleList[rIndex].isVoter, roleList[rIndex].isAdmin, roleList[rIndex].active);
    }

    // Get number of Role
    function getNumberOfRoles() external view returns (uint)
    {
        return roleList.length;
    }

    function addRole(string memory _roleId, string memory _orgId, uint _baseAccess, bool _voter, bool _admin) public
    {
        // Check if account already exists
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] == 0) {
            numberOfRoles ++;
            roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] = numberOfRoles;
            roleList.push(RoleDetails(_roleId, _orgId, _baseAccess, _voter, _admin, true));
            emit RoleCreated(_roleId, _orgId, _baseAccess, _voter, _admin);
        }
    }

    function removeRole(string calldata _roleId, string calldata _orgId) external {
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0) {
            uint rIndex = getRoleIndex(_roleId, _orgId);
            roleList[rIndex].active = false;
            emit RoleRevoked(_roleId, _orgId);
        }
    }
    // Returns the account index based on account id
    function getRoleIndex(string memory _roleId, string memory _orgId) internal view returns (uint)
    {
        return roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] - 1;
    }


    function isFullAccessRole(string calldata _roleId, string calldata _orgId, string calldata _ultParent) external view returns (bool){
        if (!(roleExists(_roleId, _orgId, _ultParent))) {
            return false;
        }
        uint rIndex;
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0) {
            rIndex = getRoleIndex(_roleId, _orgId);
        }
        else {
            rIndex = getRoleIndex(_roleId, _ultParent);
        }
        return (roleList[rIndex].active && roleList[rIndex].baseAccess == 3);
    }

    function isVoterRole(string calldata _roleId, string calldata _orgId, string calldata _ultParent) external view returns (bool){
        if (!(roleExists(_roleId, _orgId, _ultParent))) {
            return false;
        }
        uint rIndex;
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0) {
            rIndex = getRoleIndex(_roleId, _orgId);
        }
        else {
            rIndex = getRoleIndex(_roleId, _ultParent);
        }
        return (roleList[rIndex].active && roleList[rIndex].isVoter);
    }

    function isAdminRole(string calldata _roleId, string calldata _orgId, string calldata _ultParent) external view returns (bool){
        if (!(roleExists(_roleId, _orgId, _ultParent))) {
            return false;
        }
        uint rIndex;
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0) {
            rIndex = getRoleIndex(_roleId, _orgId);
        }
        else {
            rIndex = getRoleIndex(_roleId, _ultParent);
        }
        return (roleList[rIndex].active && roleList[rIndex].isAdmin);
    }

}
