pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

contract OrgManager {
    string private adminOrgId;
    PermissionsUpgradable private permUpgradable;
    // checks if first time network boot up has happened or not
    bool private networkBoot = false;
    uint private DEPTH_LIMIT = 4;
    uint private BREADTH_LIMIT = 4;
    //    enum OrgStatus {0- NotInList, 1- Proposed, 2- Approved, 3- PendingSuspension, 4- Suspended, 5- RevokeSuspension}
    struct OrgDetails {
        string orgId;
        uint status;
        string parentId;
        string fullOrgId;
        string ultParent;
        uint pindex;
        uint level;
        uint [] subOrgIndexList;
    }

    OrgDetails [] private orgList;
    mapping(bytes32 => uint) private OrgIndex;
    uint private orgNum = 0;

    // events related to Master Org add
    event OrgApproved(string _orgId, string _porgId, string _ultParent, uint _level, uint _status);
    event OrgPendingApproval(string _orgId, string _porgId, string _ultParent, uint _level, uint _status);
    event OrgSuspended(string _orgId, string _porgId, string _ultParent, uint _level);
    event OrgSuspensionRevoked(string _orgId, string _porgId, string _ultParent, uint _level);

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

    function setUpOrg(string calldata _orgId, uint _breadth, uint _depth) external
    onlyImpl
    {
        addNewOrg("", _orgId, 1, 2);
        DEPTH_LIMIT = _depth;
        BREADTH_LIMIT = _breadth;
    }

    function addNewOrg(string memory _pOrg, string memory _orgId, uint _level, uint _status) internal
    {
        bytes32 pid = "";
        bytes32 oid = "";
        uint parentIndex = 0;

        if (_level == 1) {//root
            oid = keccak256(abi.encodePacked(_orgId));
        } else {
            pid = keccak256(abi.encodePacked(_pOrg));
            oid = keccak256(abi.encodePacked(_pOrg, ".", _orgId));
        }
        orgNum++;
        OrgIndex[oid] = orgNum;
        uint id = orgList.length++;
        if (_level == 1) {
            orgList[id].level = _level;
            orgList[id].pindex = 0;
            orgList[id].fullOrgId = _orgId;
            orgList[id].ultParent = _orgId;
        } else {
            parentIndex = OrgIndex[pid] - 1;

            require(orgList[parentIndex].subOrgIndexList.length < BREADTH_LIMIT, "breadth level exceeded");
            require(orgList[parentIndex].level < DEPTH_LIMIT, "depth level exceeded");

            orgList[id].level = orgList[parentIndex].level + 1;
            orgList[id].pindex = parentIndex;
            orgList[id].ultParent = orgList[parentIndex].ultParent;
            uint subOrgId = orgList[parentIndex].subOrgIndexList.length++;
            orgList[parentIndex].subOrgIndexList[subOrgId] = id;
            orgList[id].fullOrgId = string(abi.encodePacked(_pOrg, ".", _orgId));
        }
        orgList[id].orgId = _orgId;
        orgList[id].parentId = _pOrg;
        orgList[id].status = _status;
        if (_status == 1) {
            emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level, 1);
        }
        else {
            emit OrgApproved(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level, 2);
        }
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
        addNewOrg("", _orgId, 1, 1);
    }

    // function for adding a new master org
    function addSubOrg(string calldata _pOrg, string calldata _orgId) external
    onlyImpl
    orgNotExists(string(abi.encodePacked(_pOrg, ".", _orgId)))
    {
        addNewOrg(_pOrg, _orgId, 2, 2);
    }

    function updateOrg(string calldata _orgId, uint _status) external
    onlyImpl
    orgExists(_orgId)
    returns (uint)
    {
        require((_status == 3 || _status == 5), "Operation not allowed");
        uint id = getOrgIndex(_orgId);
        require(orgList[id].level == 1, "not a master org. operation not allowed");

        uint reqStatus;
        uint pendingOp;
        if (_status == 3) {
            reqStatus = 2;
            pendingOp = 2;
        }
        else if (_status == 5) {
            reqStatus = 4;
            pendingOp = 3;
        }
        require(checkOrgStatus(_orgId, reqStatus) == true, "Operation not allowed");
        if (_status == 3) {
            suspendOrg(_orgId);
        }
        else {
            revokeOrgSuspension(_orgId);
        }
        return pendingOp;
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
        emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level, 3);
    }

    function revokeOrgSuspension(string memory _orgId) internal

    {
        require(checkOrgStatus(_orgId, 4) == true, "Org not in suspended state");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 5;
        emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level, 5);
    }

    function approveOrg(string calldata _orgId) external
    onlyImpl
    {
        require(checkOrgStatus(_orgId, 1) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgApproved(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level, 2);
    }

    function approveOrgSuspension(string memory _orgId) internal
    {
        require(checkOrgStatus(_orgId, 3) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 4;
        emit OrgSuspended(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level);
    }

    function approveOrgRevokeSuspension(string memory _orgId) internal
    {
        require(checkOrgStatus(_orgId, 5) == true, "Nothing to approve");
        uint id = getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgSuspensionRevoked(orgList[id].orgId, orgList[id].parentId, orgList[id].ultParent, orgList[id].level);
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

    // function to check if morg exists
    function checkNodeExists(string memory _pOrg, string memory _orgId) public view returns (bool)
    {
        return (!(OrgIndex[keccak256(abi.encodePacked(_pOrg, _orgId))] == 0));
    }

    // returns org and master org details based on org index
    function getOrgInfo(uint _orgIndex) external view returns (string memory, string memory, string memory, uint, uint)
    {
        return (orgList[_orgIndex].orgId, orgList[_orgIndex].parentId, orgList[_orgIndex].ultParent, orgList[_orgIndex].level, orgList[_orgIndex].status);
    }

    function getSubOrgInfo(uint _orgIndex) external view returns (uint[] memory)
    {
        return orgList[_orgIndex].subOrgIndexList;
    }

    function getSubOrgIndexLength(uint _orgIndex) external view returns (uint)
    {
        return orgList[_orgIndex].subOrgIndexList.length;
    }

    function getSubOrgIndexLength(uint _orgIndex, uint _subOrgIndex) external view returns (uint)
    {
        return orgList[_orgIndex].subOrgIndexList[_subOrgIndex];
    }

    function getUltimateParent(string calldata _orgId) external view returns (string memory)
    {
        return orgList[getOrgIndex(_orgId)].ultParent;
    }
}
