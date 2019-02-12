pragma solidity ^0.5.3;

contract Clusterkeys {

  // Struct for managing the org details
  enum Operation {None, Add, Delete}
  struct OrgKeyDetails {
    string tmKey;
    bool active;
  }
  struct OrgDetails {
    string orgId;
    string morgId;
    string pendingKey;
    Operation pendingOp;
    uint keyCount;
    OrgKeyDetails []orgKeys;
    mapping (bytes32 => uint) orgKeyIndex;

  }
  OrgDetails [] private orgList;
  mapping(bytes32 => uint) private OrgIndex;

  // Struct for managing the voter accounst for the org
  // voter struct which will be part of Master org
  struct VoterDetails {
    address vAccount;
    bool active;
  }
  struct MasterOrgDetails {
    string orgId;
    uint voterCount;
    uint validVoterCount;
    VoterDetails []voterList;
    mapping (address => uint) voterIndex;
  }
  MasterOrgDetails [] private masterOrgList;
  mapping(bytes32 => uint) private MasterOrgIndex;

  // Struct to monitor the key usage
  struct KeyUsageDetails {
    string tmKey;
    string morgId;
    uint count;
    bool pending;
  }
  KeyUsageDetails [] private keyUsage;
  mapping (bytes32 => uint) KeyIndex;

  // mapping to monitor the voting status for each acount and
  // overall voting count
  mapping (uint => mapping (address => bool)) private voteStatus;
  mapping (uint => uint) private voteCount;

  uint private orgNum = 0;
  uint private morgNum = 0;
  uint private keyNum = 0;

  // events related to Master Org add
  event MasterOrgAdded(string _orgId);
  event SubOrgAdded(string _orgId);

  // events related to Org level key management
  event OrgKeyAdded(string _orgId, string _tmKey);
  event OrgKeyDeleted(string _orgId, string _tmKey);

  // events related to org level approval process
  event ItemForApproval(string _orgId, Operation _pendingOp, string _tmKey);

  // events related to managing voting accounts for the org
  event VoterAdded(string _orgId, address _address);
  event VoterDeleted(string _orgId, address _address);


  // functions to test
  function checkOrgContractExists() external pure returns (bool){
    return true;
  }
  function getOrgVoteCount(string calldata _orgId) external view returns (uint) {
    return voteCount[getOrgIndex(_orgId)];
  }

  function getPendingOp(string calldata _orgId) external view returns (string memory, Operation) {
    uint i = getOrgIndex(_orgId);
    return (orgList[i].pendingKey, orgList[i].pendingOp);
  }

  function getVoteStatus(string calldata _orgId) external view returns (bool){
    return voteStatus[getOrgIndex(_orgId)][msg.sender];
  }
  // All internal functions
  function getVoterIndex(string memory _morgId, address _vAccount) internal view returns (uint)
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
    return masterOrgList[morgIndex].voterIndex[_vAccount] - 1;

  }
  // returns the org index for the org list
  function getOrgIndex(string memory _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // returns the voter index for the org from voter list
  function getMasterOrgIndex(string memory _orgId) internal view returns (uint)
  {
    return MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // returns the key index for the key usage list
  function getOrgKeyIndex(uint _orgIndex, string memory _tmKey) internal view returns (uint)
  {
    return orgList[_orgIndex].orgKeyIndex[keccak256(abi.encodePacked(_tmKey))] - 1;
  }

  // returns the key index for the key usage list
  function getKeyIndex(string memory _tmKey) internal view returns (uint)
  {
    return KeyIndex[keccak256(abi.encodePacked(_tmKey))] - 1;
  }

  // initialize the voter account votes to false. This will be called when a
  // new item is initiated for approval
  function voterInit(string memory _orgId) internal {
    uint orgIndex = getOrgIndex(_orgId);
    uint morgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < masterOrgList[morgIndex].voterList.length; i++){
      if (masterOrgList[morgIndex].voterList[i].active){
        voteStatus[orgIndex][masterOrgList[morgIndex].voterList[i].vAccount] = false;
      }
    }
    voteCount[orgIndex] = 0;
  }

  // processes the vote from the voter account.
  function processVote (string memory _orgId) internal {
    uint orgIndex = getOrgIndex(_orgId);
    if (voteStatus[orgIndex][msg.sender] == false ){
      voteStatus[orgIndex][msg.sender] = true;
      voteCount[orgIndex]++;
    }
  }

  // checks if enough votes have been cast for the pending operation. If yes
  // returns true
  function checkEnoughVotes (string memory _orgId, string memory _morgId) internal view returns (bool) {
    uint orgIndex = getOrgIndex(_orgId);
    uint morgIndex = getMasterOrgIndex(_morgId);

    return (voteCount[orgIndex] > masterOrgList[morgIndex].validVoterCount / 2 );
  }

  function updateKeyUsage(string memory _tmKey, string memory _morgId, Operation op) internal {
    uint keyIndex = getKeyIndex(_tmKey);
    keyUsage[keyIndex].pending = false;
    if (op == Operation.Add){
      keyUsage[keyIndex].count++;
      keyUsage[keyIndex].morgId = _morgId;
    }
    else {
      keyUsage[keyIndex].count--;
    }
  }

  // function to process the approavl for add or delete
  function processApproval(uint _orgIndex) internal {
    if(checkEnoughVotes(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId)){
      string storage locKey = orgList[_orgIndex].pendingKey;
      if (orgList[_orgIndex].pendingOp == Operation.Add){
        if (checkIfKeyExists(orgList[_orgIndex].orgId, locKey)){
          uint keyIndex = getOrgKeyIndex(_orgIndex, locKey);
          orgList[_orgIndex].orgKeys[keyIndex].active = true;
        }
        else {
          orgList[_orgIndex].keyCount++;
          orgList[_orgIndex].orgKeyIndex[keccak256(abi.encodePacked(locKey))] = orgList[_orgIndex].keyCount;
          orgList[_orgIndex].orgKeys.push(OrgKeyDetails(locKey, true));
          updateKeyUsage(orgList[_orgIndex].pendingKey, orgList[_orgIndex].morgId, orgList[_orgIndex].pendingOp);
        }
        emit OrgKeyAdded(orgList[_orgIndex].orgId, locKey);
      }
      else {
        if (checkIfKeyExists (orgList[_orgIndex].orgId, locKey)){
          uint keyIndex = getOrgKeyIndex(_orgIndex, locKey);
          orgList[_orgIndex].orgKeys[keyIndex].active = false;
          emit OrgKeyDeleted(orgList[_orgIndex].orgId, locKey);
        }
      }
      orgList[_orgIndex].pendingOp = Operation.None;
      orgList[_orgIndex].pendingKey = "";
    }
  }
  // All public functions

  // checks if the org has any voter accounts set up or not
  function checkIfVoterExists(string memory _morgId, address _address) public view returns (bool){
    uint morgIndex = getMasterOrgIndex(_morgId);
    if (masterOrgList[morgIndex].voterIndex[_address] == 0){
      return false;
    }
    uint voterIndex = getVoterIndex(_morgId, _address);
    return masterOrgList[morgIndex].voterList[voterIndex].active;
  }

  // checks if there the key is already in the list of private keys for the org
  function checkIfKeyExists(string memory _orgId, string memory _tmKey) public view returns (bool){
    uint orgIndex = getOrgIndex(_orgId);
    if (orgList[orgIndex].orgKeyIndex[keccak256(abi.encodePacked(_tmKey))] == 0){
      return false;
    }
    uint keyIndex = getOrgKeyIndex(orgIndex, _tmKey);
    return orgList[orgIndex].orgKeys[keyIndex].active;
  }
  // All extenal view functions

  // Get number of voters
  function getNumberOfVoters(string memory _morgId) public view returns (uint)
  {
    return masterOrgList[getMasterOrgIndex(_morgId)].voterCount;
  }

  // Get voter
  function getVoter(string memory _morgId, uint i) public view returns (address _addr, bool _active)
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
  	return (masterOrgList[morgIndex].voterList[i].vAccount, masterOrgList[morgIndex].voterList[i].active);
  }
  // returns the number of orgs
  function getNumberOfOrgs() external view returns (uint){
    return orgNum;
  }

  function getOrgKeyCount(string calldata _orgId) external view returns (uint){
    return orgList[getOrgIndex(_orgId)].orgKeys.length;
  }

  function getOrgKey(string calldata _orgId, uint _keyIndex) external view returns (string memory, bool){
    uint orgIndex = getOrgIndex(_orgId);
    return (orgList[orgIndex].orgKeys[_keyIndex].tmKey,orgList[orgIndex].orgKeys[_keyIndex].active);
  }

  function getOrgInfo(uint _orgIndex) external view returns (string memory, string memory){
    return (orgList[_orgIndex].orgId, orgList[_orgIndex].morgId);
  }

  // checks if the sender is one of the registered voter account for the org
  function isVoter (string calldata _orgId, address account) external view returns (bool){
    uint orgIndex = getOrgIndex(_orgId);
    uint morgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    return (masterOrgList[morgIndex].voterIndex[account] == 0);
  }

  // checks if the voter account is already in the voter accounts list for the org
  function checkVotingAccountExists(string calldata _orgId) external view returns (bool)
  {
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    return (masterOrgList[vorgIndex].validVoterCount > 0); 
  }

  // function to check if morg exists
  function checkMasterOrgExists (string calldata _morgId) external view returns (bool) {
    return (!(MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0));
  }

  // function to check if morg exists
  function checkOrgExists (string calldata _orgId) external view returns (bool) {
    return(!(OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0));
  }

  // function for checking if org exists and if there are any pending ops
  function checkOrgPendingOp (string calldata _orgId) external view returns (bool) {
    uint orgIndex = getOrgIndex(_orgId);
    return (orgList[orgIndex].pendingOp != Operation.None); 
  }

  // function for checking if org exists and if there are any pending ops
  function getOrgPendingOp (string calldata _orgId) external view returns (string memory, Operation) {
    uint orgIndex = getOrgIndex(_orgId);
    return (orgList[orgIndex].pendingKey, orgList[orgIndex].pendingOp);
  }

  // this function checks of the key proposed is in use in another master org
  function checkKeyClash (string calldata _orgId, string calldata _key) external view returns (bool) {
    uint orgIndex = getOrgIndex(_orgId);

    uint keyIndex = getKeyIndex(_key);
    if ((KeyIndex[keccak256(abi.encodePacked(_key))] != 0) &&
        (keccak256(abi.encodePacked (keyUsage[keyIndex].morgId)) != keccak256(abi.encodePacked(orgList[orgIndex].morgId)))){
      // check the count if count is greather than zero, key already in use
      if ((keyUsage[keyIndex].count > 0) || (keyUsage[keyIndex].pending)){
        return true;
      }
    }
    return false;
  }

  // All extenal update functions

  // function for adding a new master org 
  function addMasterOrg(string calldata _morgId) external
  {
    morgNum++;
    MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] = morgNum;

    uint id = masterOrgList.length++;
    masterOrgList[id].orgId = _morgId;
    masterOrgList[id].voterCount = 0;
    masterOrgList[id].validVoterCount = 0;
    emit MasterOrgAdded(_morgId);
  }

  // function for adding a new master org 
  function addSubOrg(string calldata _orgId, string calldata _morgId) external
  {
    orgNum++;
    OrgIndex[keccak256(abi.encodePacked(_orgId))] = orgNum;
    uint id = orgList.length++;
    orgList[id].orgId = _orgId;
    orgList[id].morgId = _morgId;
    orgList[id].keyCount = 0;
    orgList[id].pendingKey = "";
    orgList[id].pendingOp = Operation.None;
    emit SubOrgAdded(_morgId);
  }

  // function for adding a voter account to a master org
  function addVoter(string calldata _morgId, address _address) external
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
    masterOrgList[morgIndex].voterCount++;
    masterOrgList[morgIndex].validVoterCount++;
    masterOrgList[morgIndex].voterIndex[_address] = masterOrgList[morgIndex].voterCount;
    masterOrgList[morgIndex].voterList.push(VoterDetails(_address, true));
    emit VoterAdded(_morgId, _address);
  }

  // function for deleting a voter account to a master org
  function deleteVoter(string calldata _morgId, address _address) external
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
    if(checkIfVoterExists(_morgId, _address)){
      uint vIndex = getVoterIndex(_morgId, _address);
      masterOrgList[morgIndex].validVoterCount --;
      masterOrgList[morgIndex].voterList[vIndex].active = false;
      emit VoterDeleted(_morgId, _address);
    }

  }

  // function for adding a private key for the org. Thsi will be added once
  // approval process is complete
  function addOrgKey(string calldata _orgId, string calldata _tmKey) external
  {
    uint orgIndex = getOrgIndex(_orgId);
    if (!checkIfKeyExists (_orgId, _tmKey)){
      orgList[orgIndex].pendingKey = _tmKey;
      orgList[orgIndex].pendingOp = Operation.Add;
      voterInit(_orgId);
      // add key to key usage list for tracking
      uint keyIndex = getKeyIndex(_tmKey);
      if (KeyIndex[keccak256(abi.encodePacked(_tmKey))] == 0){
        keyNum ++;
        KeyIndex[keccak256(abi.encodePacked(_tmKey))] = keyNum;
        keyUsage.push(KeyUsageDetails(_tmKey, orgList[orgIndex].morgId, 0, true));
      }
      else {
        keyUsage[keyIndex].pending = true;
      }

      emit ItemForApproval(_orgId,Operation.Add,  _tmKey);
    }
  }

  // function for deleting a private key for the org. Thsi will be deleted once
  // approval process is complete
  function deleteOrgKey(string calldata _orgId, string calldata _tmKey) external
  {
    uint orgIndex = getOrgIndex(_orgId);
    if (checkIfKeyExists(_orgId, _tmKey)) {
      orgList[orgIndex].pendingKey = _tmKey;
      orgList[orgIndex].pendingOp = Operation.Delete;
      voterInit(_orgId);
      uint keyIndex = getKeyIndex(_tmKey);
      keyUsage[keyIndex].pending = true;
      emit ItemForApproval(_orgId, Operation.Delete,  _tmKey);
    }
  }

  // function for approving key add or delete operations
  function approvePendingOp(string calldata _orgId) external 
  {
    uint orgIndex = getOrgIndex(_orgId);
    processVote(_orgId);
    processApproval(orgIndex);
  }

}
