pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";


contract VoterManager {
    PermissionsUpgradable private permUpgradable;
    //    enum PendingOpType {0-None, 1-OrgAdd, 2-OrgSuspension, 3-OrgRevokeSuspension, 4-AssignAdminRole}
    struct PendingOpDetails {
        string orgId;
        string enodeId;
        address account;
        uint opType;
    }

    struct VoterDetails {
        address vAccount;
        bool active;
    }

    struct OrgVoterDetails {
        string orgId;
        uint voterCount;
        uint validVoterCount;
        uint voteCount;
        PendingOpDetails pendingOp;
        VoterDetails [] voterList;
        mapping(address => uint) voterIndex;
        mapping(uint => mapping(address => bool)) votingStatus;
    }

    OrgVoterDetails [] private orgVoterList;
    mapping(bytes32 => uint) private VoterOrgIndex;
    uint private orgNum = 0;

    // events related to managing voting accounts for the org
    event VoterAdded(string _orgId, address _address);
    event VoterDeleted(string _orgId, address _address);

    event VotingItemAdded(string _orgId);
    event VoteProcessed(string _orgId);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    modifier voterExists(string memory _orgId, address _address) {
        require(checkIfVoterExists(_orgId, _address) == true, "must be a voter");
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    // returns the voter index
    function getVoterIndex(string memory _orgId, address _vAccount) internal view returns (uint)
    {
        uint orgIndex = getVoterOrgIndex(_orgId);
        return orgVoterList[orgIndex].voterIndex[_vAccount] - 1;
    }

    // returns the master org index for the org from voter list
    function getVoterOrgIndex(string memory _orgId) internal view returns (uint)
    {
        return VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
    }

    // checks if the org has any voter accounts set up or not
    function checkIfVoterExists(string memory _orgId, address _address) public view returns (bool){
        uint orgIndex = getVoterOrgIndex(_orgId);
        if (orgVoterList[orgIndex].voterIndex[_address] == 0) {
            return false;
        }
        uint voterIndex = getVoterIndex(_orgId, _address);
        return orgVoterList[orgIndex].voterList[voterIndex].active;
    }

    // Get number of total voters
    function getNumberOfVoters(string calldata _orgId) external view returns (uint)
    {
        return orgVoterList[getVoterOrgIndex(_orgId)].voterCount;
    }

    // Get number of valid voters
    function getNumberOfValidVoters(string calldata _orgId) external view returns (uint)
    {
        return orgVoterList[getVoterOrgIndex(_orgId)].validVoterCount;
    }

    // checks if the voting accounts exists for the org
    function checkVotingAccountExists(string calldata _orgId) external view returns (bool)
    {
        uint orgIndex = getVoterOrgIndex(_orgId);
        return (orgVoterList[orgIndex].validVoterCount > 0);
    }

    // function for adding a voter account to a master org
    function addVoter(string calldata _orgId, address _address) external
    {
        // check if the org exists
        if (VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
            orgNum++;
            VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] = orgNum;
            uint id = orgVoterList.length++;
            orgVoterList[id].orgId = _orgId;
            orgVoterList[id].voterCount = 1;
            orgVoterList[id].validVoterCount = 1;
            orgVoterList[id].voteCount = 0;
            orgVoterList[id].pendingOp.orgId = "";
            orgVoterList[id].pendingOp.enodeId = "";
            orgVoterList[id].pendingOp.account = address(0);
            orgVoterList[id].pendingOp.opType = 0;
            orgVoterList[id].voterIndex[_address] = orgVoterList[id].voterCount;
            orgVoterList[id].voterList.push(VoterDetails(_address, true));
        }
        else {
            uint id = getVoterOrgIndex(_orgId);
            // check of the voter already present in the list
            if (orgVoterList[id].voterIndex[_address] == 0) {
                orgVoterList[id].voterCount++;
                orgVoterList[id].voterIndex[_address] = orgVoterList[id].voterCount;
                orgVoterList[id].voterList.push(VoterDetails(_address, true));
                orgVoterList[id].validVoterCount++;
            }
            else {
                uint vid = getVoterIndex(_orgId, _address);
                require(orgVoterList[id].voterList[vid].active != true, "already a voter");
                orgVoterList[id].voterList[vid].active = true;
                orgVoterList[id].validVoterCount++;
            }

        }
        emit VoterAdded(_orgId, _address);
    }

    // function for deleting a voter account to a master org
    function deleteVoter(string calldata _orgId, address _address) external voterExists(_orgId, _address)
    {
        uint id = getVoterOrgIndex(_orgId);
        uint vId = getVoterIndex(_orgId, _address);
        orgVoterList[id].validVoterCount --;
        orgVoterList[id].voterList[vId].active = false;
        emit VoterDeleted(_orgId, _address);
    }

    // function for adding an item into voting queue of the org
    function addVotingItem(string calldata _authOrg, string calldata _orgId, string calldata _enodeId, address _account, uint _pendingOp) external
    {
        // check if anything is pending approval for the org. If yes another item cannot be added
        require((checkPendingOp(_authOrg, 0)), "Items pending approval. New item cannot be added");
        uint id = getVoterOrgIndex(_authOrg);
        orgVoterList[id].pendingOp.orgId = _orgId;
        orgVoterList[id].pendingOp.enodeId = _enodeId;
        orgVoterList[id].pendingOp.account = _account;
        orgVoterList[id].pendingOp.opType = _pendingOp;
        //        init vote status
        for (uint i = 0; i < orgVoterList[id].voterList.length; i++) {
            if (orgVoterList[id].voterList[i].active) {
                orgVoterList[id].votingStatus[id][orgVoterList[id].voterList[i].vAccount] = false;
            }
        }
        // set vote count to zero
        orgVoterList[id].voteCount = 0;
        emit VotingItemAdded(_authOrg);

    }

    // process vote and update status
    function processVote(string calldata _authOrg, address _vAccount, uint _pendingOp) external voterExists(_authOrg, _vAccount) returns (bool) {

        // check something is pending approval
        require(checkPendingOp(_authOrg, _pendingOp) == true, "nothing to approve");
        uint id = getVoterOrgIndex(_authOrg);
        // check if vote already processed
        require(orgVoterList[id].votingStatus[id][_vAccount] != true, "cannot double vote");
        orgVoterList[id].voteCount++;
        orgVoterList[id].votingStatus[id][_vAccount] = true;
        emit VoteProcessed(_authOrg);
        if (orgVoterList[id].voteCount > orgVoterList[id].validVoterCount / 2) {
            // majority achieved, clean up pending op
            orgVoterList[id].pendingOp.orgId = "";
            orgVoterList[id].pendingOp.enodeId = "";
            orgVoterList[id].pendingOp.account = address(0);
            orgVoterList[id].pendingOp.opType = 0;
            return true;
        }
        return false;
    }

    function checkPendingOp(string memory _orgId, uint _pendingOp) internal view returns (bool){
        return (orgVoterList[getVoterOrgIndex(_orgId)].pendingOp.opType == _pendingOp);
    }

    function getVoteCount(string calldata _orgId) external view returns (uint, uint) {
        uint orgIndex = getVoterOrgIndex(_orgId);
        return (orgVoterList[orgIndex].voteCount, orgVoterList[orgIndex].validVoterCount);
    }

    function getPendingOpDetails(string calldata _orgId) external view returns (string memory, string memory, address, uint){
        uint orgIndex = getVoterOrgIndex(_orgId);
        return (orgVoterList[orgIndex].pendingOp.orgId, orgVoterList[orgIndex].pendingOp.enodeId, orgVoterList[orgIndex].pendingOp.account, orgVoterList[orgIndex].pendingOp.opType);
    }

}
