pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";

/** @title Voter manager contract
  * @notice This contract holds implementation logic for all account voter and
    voting functionality. This can be called only by the implementation
    contract only. there are few view functions exposed as public and
    can be called directly. these are invoked by quorum for populating
    permissions data in cache
  * @dev each voting record has an attribute operation type (opType)
    which denotes the activity type which is pending approval. This can
    have the following values:
        0 - None - indicates no pending records for the org
        1 - New org add activity  
        2 - Org suspension activity
        3 - Revoke of org suspension
        4 - Assigning admin role for a new account
        5 - Blacklisted node recovery
        6 - Blacklisted account recovery
  */
contract VoterManager {
    PermissionsUpgradable private permUpgradable;
    struct PendingOpDetails {
        string orgId;
        string enodeId;
        address account;
        uint256 opType;
    }

    struct Voter {
        address vAccount;
        bool active;
    }

    struct OrgVoterDetails {
        string orgId;
        uint256 voterCount;
        uint256 validVoterCount;
        uint256 voteCount;
        PendingOpDetails pendingOp;
        Voter [] voterList;
        mapping(address => uint256) voterIndex;
        mapping(uint256 => mapping(address => bool)) votingStatus;
    }

    OrgVoterDetails [] private orgVoterList;
    mapping(bytes32 => uint256) private VoterOrgIndex;
    uint256 private orgNum = 0;

    // events related to managing voting accounts for the org
    event VoterAdded(string _orgId, address _vAccount);
    event VoterDeleted(string _orgId, address _vAccount);

    event VotingItemAdded(string _orgId);
    event VoteProcessed(string _orgId);

    /** @notice confirms that the caller is the address of implementation
        contract
    */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /** @notice checks if account is a valid voter record and belongs to the org
        passed
      * @param _orgId - org id
      * @param _vAccount - voter account passed
      */
    modifier voterExists(string memory _orgId, address _vAccount) {
        require(_checkVoterExists(_orgId, _vAccount) == true, "must be a voter");
        _;
    }

    /** @notice constructor. sets the permissions upgradable address
      */
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice function to add a new voter account to the organization
      * @param _orgId org id
      * @param _vAccount - voter account
      * @dev voter capability is currently enabled for network level activities
        only. voting is not available for org related activities
      */
    function addVoter(string calldata _orgId, address _vAccount) external
    onlyImplementation {
        // check if the org exists
        if (VoterOrgIndex[keccak256(abi.encode(_orgId))] == 0) {
            orgNum++;
            VoterOrgIndex[keccak256(abi.encode(_orgId))] = orgNum;
            uint256 id = orgVoterList.length++;
            orgVoterList[id].orgId = _orgId;
            orgVoterList[id].voterCount = 1;
            orgVoterList[id].validVoterCount = 1;
            orgVoterList[id].voteCount = 0;
            orgVoterList[id].pendingOp.orgId = "";
            orgVoterList[id].pendingOp.enodeId = "";
            orgVoterList[id].pendingOp.account = address(0);
            orgVoterList[id].pendingOp.opType = 0;
            orgVoterList[id].voterIndex[_vAccount] = orgVoterList[id].voterCount;
            orgVoterList[id].voterList.push(Voter(_vAccount, true));
        }
        else {
            uint256 id = _getVoterOrgIndex(_orgId);
            // check if the voter is already present in the list
            if (orgVoterList[id].voterIndex[_vAccount] == 0) {
                orgVoterList[id].voterCount++;
                orgVoterList[id].voterIndex[_vAccount] = orgVoterList[id].voterCount;
                orgVoterList[id].voterList.push(Voter(_vAccount, true));
                orgVoterList[id].validVoterCount++;
            }
            else {
                uint256 vid = _getVoterIndex(_orgId, _vAccount);
                require(orgVoterList[id].voterList[vid].active != true, "already a voter");
                orgVoterList[id].voterList[vid].active = true;
                orgVoterList[id].validVoterCount++;
            }

        }
        emit VoterAdded(_orgId, _vAccount);
    }

    /** @notice function to delete a voter account from the organization
      * @param _orgId org id
      * @param _vAccount - voter account
      * @dev voter capability is currently enabled for network level activities
        only. voting is not available for org related activities
      */
    function deleteVoter(string calldata _orgId, address _vAccount) external
    onlyImplementation
    voterExists(_orgId, _vAccount) {
        uint256 id = _getVoterOrgIndex(_orgId);
        uint256 vId = _getVoterIndex(_orgId, _vAccount);
        orgVoterList[id].validVoterCount --;
        orgVoterList[id].voterList[vId].active = false;
        emit VoterDeleted(_orgId, _vAccount);
    }

    /** @notice function to a voting item for network admin accounts to vote
      * @param _authOrg org id of the authorizing org. it will be network admin org
      * @param _orgId - org id for which the voting record is being created
      * @param _enodeId - enode id for which the voting record is being created
      * @param _account - account id for which the voting record is being created
      * @param _pendingOp - operation for which voting is being done
      */
    function addVotingItem(string calldata _authOrg, string calldata _orgId,
        string calldata _enodeId, address _account, uint256 _pendingOp)
    external onlyImplementation {
        // check if anything is pending approval for the org.
        // If yes another item cannot be added
        require((_checkPendingOp(_authOrg, 0)),
            "items pending for approval. new item cannot be added");
        uint256 id = _getVoterOrgIndex(_authOrg);
        orgVoterList[id].pendingOp.orgId = _orgId;
        orgVoterList[id].pendingOp.enodeId = _enodeId;
        orgVoterList[id].pendingOp.account = _account;
        orgVoterList[id].pendingOp.opType = _pendingOp;
        // initialize vote status for voter accounts
        for (uint256 i = 0; i < orgVoterList[id].voterList.length; i++) {
            if (orgVoterList[id].voterList[i].active) {
                orgVoterList[id].votingStatus[id][orgVoterList[id].voterList[i].vAccount] = false;
            }
        }
        // set vote count to zero
        orgVoterList[id].voteCount = 0;
        emit VotingItemAdded(_authOrg);

    }

    /** @notice function processing vote of a voter account
      * @param _authOrg org id of the authorizing org. it will be network admin org
      * @param _vAccount - account id of the voter
      * @param _pendingOp - operation which is being approved
      * @return success of the voter process. either true or false
      */
    function processVote(string calldata _authOrg, address _vAccount, uint256 _pendingOp)
    external onlyImplementation voterExists(_authOrg, _vAccount) returns (bool) {
        // check something if anything is pending approval
        require(_checkPendingOp(_authOrg, _pendingOp) == true, "nothing to approve");
        uint256 id = _getVoterOrgIndex(_authOrg);
        // check if vote is already processed
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

    /** @notice returns the details of any pending operation to be approved
      * @param _orgId org id. this will be the org id of network admin org
      */
    function getPendingOpDetails(string calldata _orgId) external view
    onlyImplementation returns (string memory, string memory, address, uint256){
        uint256 orgIndex = _getVoterOrgIndex(_orgId);
        return (orgVoterList[orgIndex].pendingOp.orgId, orgVoterList[orgIndex].pendingOp.enodeId,
        orgVoterList[orgIndex].pendingOp.account, orgVoterList[orgIndex].pendingOp.opType);
    }

    /** @notice checks if the voter account exists and is linked to the org
      * @param _orgId org id
      * @param _vAccount voter account id
      * @return true or false
      */
    function _checkVoterExists(string memory _orgId, address _vAccount)
    internal view returns (bool){
        uint256 orgIndex = _getVoterOrgIndex(_orgId);
        if (orgVoterList[orgIndex].voterIndex[_vAccount] == 0) {
            return false;
        }
        uint256 voterIndex = _getVoterIndex(_orgId, _vAccount);
        return orgVoterList[orgIndex].voterList[voterIndex].active;
    }

    /** @notice checks if the pending operation exists or not
      * @param _orgId org id
      * @param _pendingOp type of operation
      * @return true or false
      */
    function _checkPendingOp(string memory _orgId, uint256 _pendingOp)
    internal view returns (bool){
        return (orgVoterList[_getVoterOrgIndex(_orgId)].pendingOp.opType == _pendingOp);
    }

    /** @notice returns the voter account index
      */
    function _getVoterIndex(string memory _orgId, address _vAccount)
    internal view returns (uint256) {
        uint256 orgIndex = _getVoterOrgIndex(_orgId);
        return orgVoterList[orgIndex].voterIndex[_vAccount] - 1;
    }

    /** @notice returns the org index for the org from voter list
      */
    function _getVoterOrgIndex(string memory _orgId)
    internal view returns (uint256) {
        return VoterOrgIndex[keccak256(abi.encode(_orgId))] - 1;
    }

}
