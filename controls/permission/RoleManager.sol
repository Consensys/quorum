pragma solidity ^0.5.3;
import "./PermissionsUpgradable.sol";


contract RoleManager {
    PermissionsImplUpgradeable private permUpgradable;
    struct RoleDetails {
        string roleId;
        string orgId;
        uint baseAccess;
        bool isVoter;
        bool active;
    }

    RoleDetails[] private roleList;
    mapping(bytes32 => uint) private roleIndex;
    uint private numberOfRoles;

    event RoleCreated(string _roleId, string _orgId);
    event RoleRevoked(string _roleId, string _orgId);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsImplUpgradeable(_permUpgradable);
    }

    function roleExists(string memory _roleId, string memory _orgId) public view returns(bool)
    {
        return (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] != 0);
    }

    function getRoleDetails(string calldata _roleId, string calldata _orgId) external view returns (string memory roleId, string memory orgId, uint accessType, bool voter, bool active)
    {
        if (!(roleExists(_roleId, _orgId))){
            return (_roleId, "", 0, false, false);
        }
        uint rIndex = getRoleIndex(_roleId, _orgId);
        return (roleList[rIndex].roleId, roleList[rIndex].orgId, roleList[rIndex].baseAccess, roleList[rIndex].isVoter, roleList[rIndex].active);
    }

    // Get number of Role
    function getNumberOfRoles() external view returns (uint)
    {
        return roleList.length;
    }

    function addRole(string memory _roleId, string memory _orgId, uint _baseAccess, bool _voter) public
    {
        // Check if account already exists
        if (roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] == 0) {
            numberOfRoles ++;
            roleIndex[keccak256(abi.encodePacked(_roleId, _orgId))] = numberOfRoles;
            roleList.push(RoleDetails(_roleId, _orgId, _baseAccess, _voter, true));
            emit RoleCreated(_roleId, _orgId);
        }
    }

    function removeRole(string calldata _roleId, string calldata _orgId) external{
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


    function isFullAccessRole(string calldata _roleId, string calldata _orgId) external view returns (bool){
        if (!(roleExists(_roleId, _orgId))){
            return false;
        }
        uint rIndex = getRoleIndex(_roleId, _orgId);
        return (roleList[rIndex].active && roleList[rIndex].baseAccess == 3);
    }

    function isVoterRole(string calldata _roleId, string calldata _orgId) external view returns (bool){
        if (!(roleExists(_roleId, _orgId))){
            return false;
        }
        uint rIndex = getRoleIndex(_roleId, _orgId);
        return (roleList[rIndex].active && roleList[rIndex].isVoter);
    }

}
