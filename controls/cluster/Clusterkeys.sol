pragma solidity ^0.5.3;

contract Clusterkeys {

  // Struct for managing the org details
  enum Operation {None, Add, Delete}
  struct OrgDetails {
    string orgId;
    string morgId;
    string [] tmKey;
    string pendingKey;
    Operation pendingOp;
  }
  OrgDetails [] private orgList;
  mapping(bytes32 => uint) private OrgIndex;

  // Struct for managing the voter accounst for the org
  struct MasterOrgDetails {
    string orgId;
    address [] voterAccount;
    string [] tmKey;
  }
  MasterOrgDetails [] private masterOrgList;
  mapping(bytes32 => uint) private MasterOrgIndex;

  // mapping to monitor the voting status for each acount and
  // overall voting count
  mapping (uint => mapping (address => bool)) private voteStatus;
  mapping (uint => uint) private voteCount;

  uint private orgNum = 0;
  uint private morgNum = 0;

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

  // initialize the voter account votes to false. This will be called when a
  // new item is initiated for approval
  function voterInit(string memory _orgId) internal {
    uint orgIndex = getOrgIndex(_orgId);
    uint morgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < masterOrgList[morgIndex].voterAccount.length; i++){
      voteStatus[orgIndex][masterOrgList[morgIndex].voterAccount[i]] = false;
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
    if (voteCount[orgIndex] > masterOrgList[morgIndex].voterAccount.length / 2 ){
      return true;
    }
    return false;
  }

  function checkKeyInUse(string memory _orgId, string memory _morgId, string memory _key) internal view returns (bool) {
    bool keyInUse = false;
    for (uint i = 0; i < orgList.length; i++){
      if (keccak256(abi.encodePacked(orgList[i].orgId)) == keccak256(abi.encodePacked(_orgId))){
        continue;
      }
      if (keccak256(abi.encodePacked(orgList[i].morgId)) == keccak256(abi.encodePacked(_morgId))){
        for (uint j = 0; j < orgList[i].tmKey.length; j++){
          if (keccak256(abi.encodePacked(orgList[i].tmKey[j])) == keccak256(abi.encodePacked(_key))){
            keyInUse = true;
            break;
          }
        }
      }
      if (keyInUse == true) {
        return true;
      }
    }
    return false;
  }
  // updates the master keys list with the key being added or deleted
  function masterKeyUpdate(string memory _orgId, string memory _morgId, string memory _key, Operation _op) internal {
    uint morgIndex = getMasterOrgIndex(_morgId);
    if (_op == Operation.Add) {
      // check if the key is existing. if yes ignore else add to master list
      bool keyExists = false;
      for (uint i = 0; i < masterOrgList[morgIndex].tmKey.length; i++){
        if(keccak256(abi.encodePacked(masterOrgList[morgIndex].tmKey[i])) == keccak256(abi.encodePacked(_key))){
          keyExists = true;
          break;
        }
      }
      if (keyExists == false ){
        masterOrgList[morgIndex].tmKey.push(_key);
      }
    }
    else {
      // the key can be deleted from master list only when none of the suborgs have the
      // key in the private keys
      if (!(checkKeyInUse(_orgId, _morgId, _key))){
        uint index;
        for (index = 0; index < masterOrgList[morgIndex].tmKey.length; index++) {
          if(keccak256(abi.encodePacked(masterOrgList[morgIndex].tmKey[index])) == keccak256(abi.encodePacked(_key))) {
            break;
          }
        }
        for (uint j = index; j <  masterOrgList[morgIndex].tmKey.length -1; j++){
          masterOrgList[morgIndex].tmKey[j] = masterOrgList[morgIndex].tmKey[j+1];
        }
        delete masterOrgList[morgIndex].tmKey[masterOrgList[morgIndex].tmKey.length -1];
        masterOrgList[morgIndex].tmKey.length --;
      }
    }
  }

  // function to process the approavl for add or delete
  function processApproval(uint _orgIndex) internal {
    if(checkEnoughVotes(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId)){
      string storage locKey = orgList[_orgIndex].pendingKey;
      if (orgList[_orgIndex].pendingOp == Operation.Add){
        orgList[_orgIndex].tmKey.push(orgList[_orgIndex].pendingKey);
        masterKeyUpdate(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId, orgList[_orgIndex].pendingKey, Operation.Add);
        emit OrgKeyAdded(orgList[_orgIndex].orgId, locKey);
      }
      else {
        bool keyExists = false;
        uint i = 0;
        (keyExists, i) = checkIfKeyExists (orgList[_orgIndex].orgId, locKey);
        for (uint j = i; j <  orgList[_orgIndex].tmKey.length -1; j++){
          orgList[_orgIndex].tmKey[j] = orgList[_orgIndex].tmKey[j+1];
        }
        delete orgList[_orgIndex].tmKey[orgList[_orgIndex].tmKey.length -1];
        orgList[_orgIndex].tmKey.length --;
        masterKeyUpdate(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId, orgList[_orgIndex].pendingKey, Operation.Delete);
        emit OrgKeyDeleted(orgList[_orgIndex].orgId, locKey);
      }
      orgList[_orgIndex].pendingOp = Operation.None;
      orgList[_orgIndex].pendingKey = "";
    }
  }
  // All public functions

  // checks if the org has any voter accounts set up or not
  function checkIfVoterExists(string memory _morgId, address _address) public view returns (bool, uint){
    bool keyExists = false;
    uint voterIndex = getMasterOrgIndex(_morgId);
    uint i;
    for (i = 0; i < masterOrgList[voterIndex].voterAccount.length; i++){
      if(keccak256(abi.encodePacked(masterOrgList[voterIndex].voterAccount[i])) == keccak256(abi.encodePacked(_address))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }

  // checks if there the key is already in the list of private keys for the org
  function checkIfKeyExists(string memory _orgId, string memory _tmKey) public view returns (bool, uint){
    bool keyExists = false;
    uint orgIndex = getOrgIndex(_orgId);
    uint i;
    for (i = 0; i < orgList[orgIndex].tmKey.length; i++){
      if(keccak256(abi.encodePacked(orgList[orgIndex].tmKey[i])) == keccak256(abi.encodePacked(_tmKey))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }
  // All extenal view functions

  // Get number of voters
  function getNumberOfVoters(string memory _morgId) public view returns (uint)
  {
    return masterOrgList[getMasterOrgIndex(_morgId)].voterAccount.length;
  }

  // Get voter
  function getVoter(string memory _morgId, uint i) public view returns (address _addr)
  {
  	return masterOrgList[getMasterOrgIndex(_morgId)].voterAccount[i];
  }
  // returns the number of orgs
  function getNumberOfOrgs() external view returns (uint){
    return orgNum;
  }

  function getOrgKeyCount(string calldata _orgId) external view returns (uint){
    return orgList[getOrgIndex(_orgId)].tmKey.length;
  }

  function getOrgKey(string calldata _orgId, uint _keyIndex) external view returns (string memory ){
    return orgList[getOrgIndex(_orgId)].tmKey[_keyIndex];
  }

  function getOrgInfo(uint _orgIndex) external view returns (string memory, string memory){
    return (orgList[_orgIndex].orgId, orgList[_orgIndex].morgId);
  }

  // checks if the sender is one of the registered voter account for the org
  function isVoter (string calldata _orgId, address account) external view returns (bool){
    bool flag = false;
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < masterOrgList[vorgIndex].voterAccount.length; i++){
      if ( masterOrgList[vorgIndex].voterAccount[i] == account){
        flag = true;
        break;
      }
    }
    return flag;
  }

  // checks if the voter account is already in the voter accounts list for the org
  function checkVotingAccountExists(string calldata _orgId) external view returns (bool)
  {
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    if (masterOrgList[vorgIndex].voterAccount.length == 0) {
      return false;
    }
    return true;
  }

  // function to check if morg exists
  function checkMasterOrgExists (string calldata _morgId) external view returns (bool) {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0) {
      return false;
    }
    else {
      return true;
    }
  }

  // function to check if morg exists
  function checkOrgExists (string calldata _orgId) external view returns (bool) {
    if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      return false;
    }
    else {
      return true;
    }
  }

  // function for checking if org exists and if there are any pending ops
  function checkOrgPendingOp (string calldata _orgId) external view returns (bool) {
    uint orgIndex = getOrgIndex(_orgId);
    if (orgList[orgIndex].pendingOp != Operation.None) {
      return true;
    }
    return false;
  }

  // function for checking if org exists and if there are any pending ops
  function getOrgPendingOp (string calldata _orgId) external view returns (string memory, Operation) {
    uint orgIndex = getOrgIndex(_orgId);
    return (orgList[orgIndex].pendingKey, orgList[orgIndex].pendingOp);
  }

  // this function checks of the key proposed is in use in another master org
  function checkKeyClash (string calldata _orgId, string calldata _key) external view returns (bool) {
    bool ret = false;
    uint orgIndex = getOrgIndex(_orgId);
    // check if the key is already in use with other orgs
    for (uint i = 0; i < masterOrgList.length; i++){
      if (keccak256( abi.encodePacked (masterOrgList[i].orgId)) != keccak256( abi.encodePacked(orgList[orgIndex].morgId))) {
        // check if the key is already present in the key list for the org
        for (uint j = 0; j < masterOrgList[i].tmKey.length; j++){
          if (keccak256(abi.encodePacked(masterOrgList[i].tmKey[j])) == keccak256(abi.encodePacked(_key))) {
            ret = true;
            break;
          }
        }
      }
      if (ret) {
        break;
      }
    }
    if (ret){
      return ret;
    }
    // check if the key is pending approval for any of the orgs
    for (uint k = 0; k < orgList.length; k++){
      if ((keccak256(abi.encodePacked(orgList[k].orgId)) != keccak256(abi.encodePacked(_orgId))) &&
        (keccak256(abi.encodePacked(orgList[k].morgId)) != keccak256(abi.encodePacked(orgList[orgIndex].morgId))))
        {
        if ((orgList[k].pendingOp == Operation.Add) &&
            (keccak256(abi.encodePacked(orgList[k].pendingKey)) == keccak256(abi.encodePacked(_key)))){
          ret = true;
          break;
        }
      }
    }
    return ret;
  }

  // All extenal update functions

  // function for adding a new master org 
  function addMasterOrg(string calldata _morgId) external
  {
    morgNum++;
    MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] = morgNum;
    masterOrgList.push( MasterOrgDetails(_morgId, new address[](0), new string[](0)));
    emit MasterOrgAdded(_morgId);
  }

  // function for adding a new master org 
  function addSubOrg(string calldata _orgId, string calldata _morgId) external
  {
    orgNum++;
    OrgIndex[keccak256(abi.encodePacked(_orgId))] = orgNum;
    orgList.push( OrgDetails(_orgId, _morgId, new string[](0), new string(0), Operation.None ));
    emit SubOrgAdded(_morgId);
  }

  // function for adding a voter account to a master org
  function addVoter(string calldata _morgId, address _address) external
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
    masterOrgList[morgIndex].voterAccount.push(_address);
    emit VoterAdded(_morgId, _address);
  }

  // function for deleting a voter account to a master org
  function deleteVoter(string calldata _morgId, address _address) external
  {
    uint morgIndex = getMasterOrgIndex(_morgId);
    (bool voterExists, uint i) = checkIfVoterExists(_morgId, _address);

    if (voterExists == true) {
      for (uint j = i; j <  masterOrgList[morgIndex].voterAccount.length -1; j++){
        masterOrgList[morgIndex].voterAccount[j] = masterOrgList[morgIndex].voterAccount[j+1];
      }
      delete masterOrgList[morgIndex].voterAccount[masterOrgList[morgIndex].voterAccount.length -1];
      masterOrgList[morgIndex].voterAccount.length --;
      emit VoterDeleted(_morgId, _address);
    }
  }

  // function for adding a private key for the org. Thsi will be added once
  // approval process is complete
  function addOrgKey(string calldata _orgId, string calldata _tmKey) external
  {
    uint orgIndex = getOrgIndex(_orgId);
    orgList[orgIndex].pendingKey = _tmKey;
    orgList[orgIndex].pendingOp = Operation.Add;
    voterInit(_orgId);
    emit ItemForApproval(_orgId,Operation.Add,  _tmKey);
  }

  // function for deleting a private key for the org. Thsi will be deleted once
  // approval process is complete
  function deleteOrgKey(string calldata _orgId, string calldata _tmKey) external
  {
    uint orgIndex = getOrgIndex(_orgId);
    uint i = 0;
    bool keyExists = false;
    (keyExists, i) = checkIfKeyExists (_orgId, _tmKey);
    if (keyExists == true) {
      orgList[orgIndex].pendingKey = _tmKey;
      orgList[orgIndex].pendingOp = Operation.Delete;
      voterInit(_orgId);
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
