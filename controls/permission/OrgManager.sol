pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

contract OrgManager {
    string private adminOrgId;
    PermissionsUpgradable private permUpgradable;
    // checks if first time network boot up has happened or not
    bool private networkBoot = false;
    //    enum OrgStatus {0- NotInList, 1- Proposed, 2- Approved, 3- PendingSuspension, 4- Suspended, 5- RevokeSuspension}
    struct OrgDetails {
        string orgId;
        uint status;
    }
    OrgDetails [] private orgList;
    mapping(bytes32 => uint) private OrgIndex;
    uint private orgNum = 0;

    // events related to Master Org add
    event OrgApproved(string _orgId);
    event OrgPendingApproval(string _orgId, uint _type);
    event OrgSuspended(string _orgId);
    event OrgSuspensionRevoked(string _orgId);

    event Dummy(string _msg);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    modifier orgNotExists(string memory _orgId) {
        require(checkOrgExists(_orgId) == false, "Org already exists");
        _;
    }

    modifier orgExists(string memory _orgId) {
        require(checkOrgExists(_orgId) == true, "Org does not exists");
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    function getImpl() public view returns (address) {
        return permUpgradable.getPermImpl();
    }

    function addAdminOrg(string calldata _orgId) external
    onlyImpl
    {
        addNewOrg(_orgId, 2);
        emit OrgApproved(_orgId);
    }

    function addNewOrg(string memory _orgId, uint _status) internal
    {
        orgNum++;
        OrgIndex[keccak256(abi.encodePacked(_orgId))] = orgNum;
        uint id = orgList.length++;
        orgList[id].orgId = _orgId;
        orgList[id].status = _status;
    }

    function getNumberOfOrgs() public view returns (uint)
    {
        return orgList.length;
    }

    // Org related functions
    // returns the org index for the org list
    function getOrgIndex(string memory _orgId) public view returns (uint)
    {
        return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
    }

    function getOrgStatus(string memory _orgId) public view returns (uint)
    {
        return orgList[OrgIndex[keccak256(abi.encodePacked(_orgId))]].status;
    }

    // function for adding a new master org
    function addOrg(string calldata _orgId) external
    onlyImpl
    orgNotExists(_orgId)
    {
        addNewOrg(_orgId, 1);
        emit OrgPendingApproval(_orgId, 1);
    }

    function updateOrg(string calldata _orgId, uint _status) external
    onlyImpl
    orgExists(_orgId)
    {
        if (_status == 3) {
            suspendOrg(_orgId);
        }
        else {
            revokeOrgSuspension(_orgId);
        }
    }

    function approveOrgStatusUpdate(string calldata _orgId, uint _status) external
    onlyImpl
    orgExists(_orgId)
    {
        if (_status == 3) {
            approveOrgSuspension(_orgId);
        }
        else {
            approveOrgRevokeSuspension(_orgId);
        }
    }


    // function for adding a new master org
    function suspendOrg(string memory _orgId) internal
    {
        require(checkOrgStatus(_orgId, 2) == true, "Org not in approved state");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 3;
        emit OrgPendingApproval(_orgId, 3);
    }

    function revokeOrgSuspension(string memory _orgId) internal

    {
        require(checkOrgStatus(_orgId, 4) == true, "Org not in suspended state");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 5;
        emit OrgPendingApproval(_orgId, 5);
    }

    function approveOrg(string calldata _orgId) external
    onlyImpl
    {
        require(checkOrgStatus(_orgId, 1) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgApproved(_orgId);
    }

    function approveOrgSuspension(string memory _orgId) internal
    {
        require(checkOrgStatus(_orgId, 3) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 4;
        emit OrgSuspended(_orgId);
    }

    function approveOrgRevokeSuspension(string memory _orgId) internal
    {
        require(checkOrgStatus(_orgId, 5) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgSuspensionRevoked(_orgId);
    }

    function checkOrgStatus(string memory _orgId, uint _orgStatus) public view returns (bool){
        uint id = getOrgIndex(_orgId);
        return ((OrgIndex[keccak256(abi.encodePacked(_orgId))] != 0) && orgList[id].status == _orgStatus);
    }

    // function to check if morg exists
    function checkOrgExists(string memory _orgId) public view returns (bool)
    {
        return (!(OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0));
    }

    // returns org and master org details based on org index
    function getOrgInfo(uint _orgIndex) external view returns (string memory, uint)
    {
        return (orgList[_orgIndex].orgId, orgList[_orgIndex].status);
    }
}
