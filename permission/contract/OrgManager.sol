pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";
/** @title Organization manager contract
  * @notice This contract holds implementation logic for all org management
    functionality. This can be called only by the implementation
    contract only. there are few view functions exposed as public and
    can be called directly. these are invoked by quorum for populating
    permissions data in cache
  * @dev the status of the organization is denoted by a set of integer
    values. These are as below:
        0 - Not in list,
        1 - Org proposed for approval by network admins
        2 - Org in Approved status
        3 - Org proposed for suspension and pending approval by network admins
        4 - Org in Suspended,
     Once the node is blacklisted no further activity on the node is
     possible.
  */
contract OrgManager {
    string private adminOrgId;
    PermissionsUpgradable private permUpgradable;
    // checks if first time network boot up has happened or not
    bool private networkBoot = false;

    // variables which control the breadth and depth of the sub org tree
    uint private DEPTH_LIMIT = 4;
    uint private BREADTH_LIMIT = 4;

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
    event OrgApproved(string _orgId, string _porgId, string _ultParent,
        uint _level, uint _status);
    event OrgPendingApproval(string _orgId, string _porgId, string _ultParent,
        uint _level, uint _status);
    event OrgSuspended(string _orgId, string _porgId, string _ultParent,
        uint _level);
    event OrgSuspensionRevoked(string _orgId, string _porgId, string _ultParent,
        uint _level);

    /** @notice confirms that the caller is the address of implementation
        contract
    */
    modifier onlyImplementation{
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /** @notice checks if the org id does not exists
      * @param _orgId - org id
      * @return true if org does not exist
      */
    modifier orgDoesNotExist(string memory _orgId) {
        require(checkOrgExists(_orgId) == false, "org exists");
        _;
    }

    /** @notice checks if the org id does exists
      * @param _orgId - org id
      * @return true if org exists
      */
    modifier orgExists(string memory _orgId) {
        require(checkOrgExists(_orgId) == true, "org does not exist");
        _;
    }

    /** @notice constructor. sets the permissions upgradable address
      */
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice called at the time of network initialization. sets the depth
        breadth for sub orgs creation. and creates the default network
        admin org as per config file
      */
    function setUpOrg(string calldata _orgId, uint256 _breadth, uint256 _depth) external
    onlyImplementation {
        _addNewOrg("", _orgId, 1, 2);
        DEPTH_LIMIT = _depth;
        BREADTH_LIMIT = _breadth;
    }
    /** @notice function for adding a new master org to the network
      * @param _orgId unique org id to be added
      * @dev org will be added if it does exist
      */
    function addOrg(string calldata _orgId) external
    onlyImplementation
    orgDoesNotExist(_orgId) {
        _addNewOrg("", _orgId, 1, 1);
    }

    /** @notice function for adding a new sub org under a parent org
      * @param _pOrgId unique org id to be added
      * @dev org will be added if it does exist
      */
    function addSubOrg(string calldata _pOrgId, string calldata _orgId) external
    onlyImplementation
    orgDoesNotExist(string(abi.encodePacked(_pOrgId, ".", _orgId))) {
        _addNewOrg(_pOrgId, _orgId, 2, 2);
    }

    /** @notice updates the status of a master org.
      * @param _orgId unique org id to be added
      * @param _action action being performed
      * @dev status cannot be updated for sub orgs.
        This function can be called for the following actions:
            1 - to suspend an org
            2 - to activate the org back
      */
    function updateOrg(string calldata _orgId, uint256 _action) external
    onlyImplementation
    orgExists(_orgId)
    returns (uint256){
        require((_action == 1 || _action == 2), "invalid action. operation not allowed");
        uint256 id = _getOrgIndex(_orgId);
        require(orgList[id].level == 1, "not a master org. operation not allowed");

        uint256 reqStatus;
        uint256 pendingOp;
        if (_action == 1) {
            reqStatus = 2;
            pendingOp = 2;
        }
        else if (_action == 2) {
            reqStatus = 4;
            pendingOp = 3;
        }
        require(checkOrgStatus(_orgId, reqStatus) == true,
            "org status does not allow the operation");
        if (_action == 1) {
            _suspendOrg(_orgId);
        }
        else {
            _revokeOrgSuspension(_orgId);
        }
        return pendingOp;
    }

    /** @notice function to approve org status change for master orgs
      * @param _orgId unique org id to be added
      * @param _action approval for action
      * @dev status cannot be updated for sub orgs.
        This function can be called for the following actions:
            1 - to suspend an org
            2 - to activate the org back
      */
    function approveOrgStatusUpdate(string calldata _orgId, uint256 _action) external
    onlyImplementation
    orgExists(_orgId) {
        if (_action == 1) {
            _approveOrgSuspension(_orgId);
        }
        else {
            _approveOrgRevokeSuspension(_orgId);
        }
    }

    /** @notice function to approve org status change for master orgs
      * @param _orgId unique org id to be added
      */
    function approveOrg(string calldata _orgId) external
    onlyImplementation {
        require(checkOrgStatus(_orgId, 1) == true, "nothing to approve");
        uint256 id = _getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgApproved(orgList[id].orgId, orgList[id].parentId,
            orgList[id].ultParent, orgList[id].level, 2);
    }

    /** @notice returns org info for a given org index
      * @param _orgIndex org index
      * @return org id
      * @return parent org id
      * @return ultimate parent id
      * @return level in the org tree
      * @return status
      */
    function getOrgInfo(uint256 _orgIndex) external view returns (string memory,
        string memory, string memory, uint256, uint256) {
        return (orgList[_orgIndex].orgId, orgList[_orgIndex].parentId,
        orgList[_orgIndex].ultParent, orgList[_orgIndex].level, orgList[_orgIndex].status);
    }

    /** @notice returns org info for a given org id
      * @param _orgId org id
      * @return org id
      * @return parent org id
      * @return ultimate parent id
      * @return level in the org tree
      * @return status
      */
    function getOrgDetails(string calldata _orgId) external view returns (string memory,
        string memory, string memory, uint256, uint256) {
        if (!checkOrgExists(_orgId)) {
            return (_orgId, "", "", 0, 0);
        }
        uint256 _orgIndex = _getOrgIndex(_orgId);
        return (orgList[_orgIndex].orgId, orgList[_orgIndex].parentId,
        orgList[_orgIndex].ultParent, orgList[_orgIndex].level, orgList[_orgIndex].status);
    }

    /** @notice returns the array of sub org indexes for the given org
      * @param _orgId org id
      * @return array of sub org indexes
      */
    function getSubOrgIndexes(string calldata _orgId) external view returns (uint[] memory) {
        require(checkOrgExists(_orgId) == true, "org does not exist");
        uint256 _orgIndex = _getOrgIndex(_orgId);
        return (orgList[_orgIndex].subOrgIndexList);
    }
    /** @notice returns the master org id for the given org or sub org
      * @param _orgId org id
      * @return master org id
      */
    function getUltimateParent(string calldata _orgId) external view
    onlyImplementation
    returns (string memory) {
        return orgList[_getOrgIndex(_orgId)].ultParent;
    }

    /** @notice returns the total number of orgs in the network
      * @return master org id
      */
    function getNumberOfOrgs() public view returns (uint256) {
        return orgList.length;
    }

    /** @notice confirms that org status is same as passed status
      * @param _orgId org id
      * @param _orgStatus org status
      * @return true or false
      */
    function checkOrgStatus(string memory _orgId, uint256 _orgStatus)
    public view returns (bool){
        if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
            return false;
        }
        uint256 id = _getOrgIndex(_orgId);
        return ((OrgIndex[keccak256(abi.encodePacked(_orgId))] != 0)
        && orgList[id].status == _orgStatus);
    }

    /** @notice confirms that org status either active or pending suspension
      * @param _orgId org id
      * @return true or false
      */
    function checkOrgActive(string memory _orgId)
    public view returns (bool){
        if (OrgIndex[keccak256(abi.encodePacked(_orgId))] != 0) {
            uint256 id = _getOrgIndex(_orgId);
            if (orgList[id].status == 2 || orgList[id].status == 3) {
                uint256 uid = _getOrgIndex(orgList[id].ultParent);
                if (orgList[uid].status == 2 || orgList[uid].status == 3) {
                    return true;
                }
            }
        }
        return false;
    }

    /** @notice confirms if the org exists in the network
      * @param _orgId org id
      * @return true or false
      */
    function checkOrgExists(string memory _orgId) public view returns (bool) {
        return (!(OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0));
    }

    /** @notice updates the org status to suspended
      * @param _orgId org id
      */
    function _suspendOrg(string memory _orgId) internal {
        require(checkOrgStatus(_orgId, 2) == true,
            "org not in approved status. operation cannot be done");
        uint256 id = _getOrgIndex(_orgId);
        orgList[id].status = 3;
        emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId,
            orgList[id].ultParent, orgList[id].level, 3);
    }

    /** @notice revokes the suspension of an org
      * @param _orgId org id
      */
    function _revokeOrgSuspension(string memory _orgId) internal {
        require(checkOrgStatus(_orgId, 4) == true, "org not in suspended state");
        uint256 id = _getOrgIndex(_orgId);
        orgList[id].status = 5;
        emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId,
            orgList[id].ultParent, orgList[id].level, 5);
    }

    /** @notice approval function for org suspension activity
      * @param _orgId org id
      */
    function _approveOrgSuspension(string memory _orgId) internal {
        require(checkOrgStatus(_orgId, 3) == true, "nothing to approve");
        uint256 id = _getOrgIndex(_orgId);
        orgList[id].status = 4;
        emit OrgSuspended(orgList[id].orgId, orgList[id].parentId,
            orgList[id].ultParent, orgList[id].level);
    }

    /** @notice approval function for revoking org suspension
      * @param _orgId org id
      */
    function _approveOrgRevokeSuspension(string memory _orgId) internal {
        require(checkOrgStatus(_orgId, 5) == true, "nothing to approve");
        uint256 id = _getOrgIndex(_orgId);
        orgList[id].status = 2;
        emit OrgSuspensionRevoked(orgList[id].orgId, orgList[id].parentId,
            orgList[id].ultParent, orgList[id].level);
    }

    /** @notice function to add a new organization
      * @param _pOrgId parent org id
      * @param _orgId org id
      * @param _level level in org hierarchy
      * @param _status status of the org
      */
    function _addNewOrg(string memory _pOrgId, string memory _orgId,
        uint256 _level, uint _status) internal {
        bytes32 pid = "";
        bytes32 oid = "";
        uint256 parentIndex = 0;

        if (_level == 1) {//root
            oid = keccak256(abi.encodePacked(_orgId));
        } else {
            pid = keccak256(abi.encodePacked(_pOrgId));
            oid = keccak256(abi.encodePacked(_pOrgId, ".", _orgId));
        }
        orgNum++;
        OrgIndex[oid] = orgNum;
        uint256 id = orgList.length++;
        if (_level == 1) {
            orgList[id].level = _level;
            orgList[id].pindex = 0;
            orgList[id].fullOrgId = _orgId;
            orgList[id].ultParent = _orgId;
        } else {
            parentIndex = OrgIndex[pid] - 1;

            require(orgList[parentIndex].subOrgIndexList.length < BREADTH_LIMIT,
                "breadth level exceeded");
            require(orgList[parentIndex].level < DEPTH_LIMIT,
                "depth level exceeded");

            orgList[id].level = orgList[parentIndex].level + 1;
            orgList[id].pindex = parentIndex;
            orgList[id].ultParent = orgList[parentIndex].ultParent;
            uint256 subOrgId = orgList[parentIndex].subOrgIndexList.length++;
            orgList[parentIndex].subOrgIndexList[subOrgId] = id;
            orgList[id].fullOrgId = string(abi.encodePacked(_pOrgId, ".", _orgId));
        }
        orgList[id].orgId = _orgId;
        orgList[id].parentId = _pOrgId;
        orgList[id].status = _status;
        if (_status == 1) {
            emit OrgPendingApproval(orgList[id].orgId, orgList[id].parentId,
                orgList[id].ultParent, orgList[id].level, 1);
        }
        else {
            emit OrgApproved(orgList[id].orgId, orgList[id].parentId,
                orgList[id].ultParent, orgList[id].level, 2);
        }
    }

    /** @notice returns the org index from the org list for the given org
      * @return org index
      */
    function _getOrgIndex(string memory _orgId) private view returns (uint){
        return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
    }

}
