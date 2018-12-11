pragma solidity ^0.4.23;

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
  event MasterOrgExists(string _orgId);
  event MasterOrgNotFound(string _orgId);

  // events related to Sub Org add
  event SubOrgAdded(string _orgId);
  event SubOrgExists(string _orgId);
  event SubOrgNotFound(string _orgId);

  // events related to Org level key management
  event OrgKeyAdded(string _orgId, string _tmKey);
  event OrgKeyDeleted(string _orgId, string _tmKey);
  event KeyNotFound(string _tmKey);
  event KeyExists(string _orgId, string _tmKey);
  event OrgNotFound(string _orgId);

  // events related to org level approval process
  event PendingApproval(string _orgId);
  event ItemForApproval(string _orgId, Operation _pendingOp, string _tmKey);
  event NothingToApprove(string _orgId);

  // events related to managing voting accounts for the org
  event NoVotingAccount(string _orgId);
  event VoterAdded(string _orgId, address _address);
  event VoterNotFound(string _orgId, address _address);
  event VoterDeleted(string _orgId, address _address);
  event VoterExists(string _orgId, address _address);

  // events related to helper functions to print all org keys and voter keys
  event PrintAll(string _orgId, string _tmKey);
  event PrintVoter(string _orgId, address _voterAccount);

  // returns the org index for the org list
  function getOrgIndex(string _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // returns the voter index for the org from voter list
  function getMasterOrgIndex(string _orgId) internal view returns (uint)
  {
    return MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // checks if the sender is one of the registered voter account for the org
  modifier canVote(string _orgId){
    bool flag = false;
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < masterOrgList[vorgIndex].voterAccount.length; i++){
      if ( masterOrgList[vorgIndex].voterAccount[i] == msg.sender){
        flag = true;
        break;
      }
    }
    require(flag, "Account cannot vote");
    _;
  }

  // checks if the org has any voter accounts set up or not
  function checkIfVoterExists(string _morgId, address _address) internal view returns (bool, uint){
    bool keyExists = false;
    uint voterIndex = getMasterOrgIndex(_morgId);
    for (uint i = 0; i < masterOrgList[voterIndex].voterAccount.length; i++){
      if(keccak256(abi.encodePacked(masterOrgList[voterIndex].voterAccount[i])) == keccak256(abi.encodePacked(_address))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }

  // checks if the voter account is already in the voter accounts list for the org
  function checkVotingAccountExists(string _orgId) internal returns (bool)
  {
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    if (masterOrgList[vorgIndex].voterAccount.length == 0) {
      emit NoVotingAccount(_orgId);
      return false;
    }
    return true;
  }

  // checks if there are any pending unapproved actions for the org
  function checkingPendingOp(string _orgId) internal view returns (bool)
  {
    if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0){
      return false;
    }
    uint orgIndex = getOrgIndex(_orgId);
    if (orgList[orgIndex].pendingOp != Operation.None) {
      return true;
    }
    return false;
  }

  // checks if there the key is already in the list of private keys for the org
  function checkIfKeyExists(string _orgId, string _tmKey) internal view returns (bool, uint){
    bool keyExists = false;
    uint orgIndex = getOrgIndex(_orgId);
    for (uint i = 0; i < orgList[orgIndex].tmKey.length; i++){
      if(keccak256(abi.encodePacked(orgList[orgIndex].tmKey[i])) == keccak256(abi.encodePacked(_tmKey))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }
  // function to check if morg exists
  function checkMasterOrgExists (string _morgId) external view returns (bool) {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0) {
      return false;
    }
    else {
      return true;
    }
  }
  // function for adding a new master org 
  function addMasterOrg(string _morgId) external
  {
    morgNum++;
    MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] = morgNum;
    masterOrgList.push( MasterOrgDetails(_morgId, new address[](0), new string[](0)));
    emit MasterOrgAdded(_morgId);
  }

  // function for adding a voter account to a master org
  function addVoter(string _morgId, address _address) external
  {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0) {
      emit MasterOrgNotFound(_morgId);
    }
    else {
      bool voterExists = false;
      uint i = 0;
      (voterExists, i) = checkIfVoterExists(_morgId, _address);
      if (voterExists) {
        emit VoterExists(_morgId, _address);
      }
      else {
        uint morgIndex = getMasterOrgIndex(_morgId);
        masterOrgList[morgIndex].voterAccount.push(_address);
        emit VoterAdded(_morgId, _address);
      }
    }
  }

  // function for deleting a voter account to a master org
  function deleteVoter(string _morgId, address _address) external
  {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0) {
      emit MasterOrgNotFound(_morgId);
    }
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
    else {
      emit VoterNotFound(_morgId, _address);
    }
  }

  // function for adding a new master org 
  function addSubOrg(string _orgId, string _morgId) external
  {
    // check if master org exists
    if (MasterOrgIndex[keccak256(abi.encodePacked(_morgId))] == 0){
      emit MasterOrgNotFound(_morgId);
    }
    else {
      if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
        orgNum++;
        OrgIndex[keccak256(abi.encodePacked(_orgId))] = orgNum;
        orgList.push( OrgDetails(_orgId, _morgId, new string[](0), new string(0), Operation.None ));
        emit SubOrgAdded(_morgId);
      }
      else {
        emit SubOrgExists(_morgId);
      }
    }
  }

  // function for checking if org exists and if there are any pending ops
  function checkOrgPendingOp (string _orgId) internal returns (bool) {
    if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      emit OrgNotFound(_orgId);
      return false;
    }
    if (checkVotingAccountExists(_orgId)){
      if (checkingPendingOp(_orgId)){
        emit PendingApproval(_orgId);
        return false;
      }
    }
    else {
      emit NoVotingAccount(_orgId);
      return false;
    }
    return true;
  }

  // function for adding a private key for the org. Thsi will be added once
  // approval process is complete
  function addOrgKey(string _orgId, string _tmKey) external
  {
    bool ret = checkOrgPendingOp(_orgId);
    if (ret){
      bool keyExists = false;
      uint i = 0;
      (keyExists, i) = checkIfKeyExists(_orgId, _tmKey);
      if (keyExists) {
        emit KeyExists(_orgId, _tmKey);
      }
      else {
        uint orgIndex;
        orgIndex = getOrgIndex(_orgId);
        orgList[orgIndex].pendingKey = _tmKey;
        orgList[orgIndex].pendingOp = Operation.Add;
        voterInit(_orgId);
        emit ItemForApproval(_orgId,Operation.Add,  _tmKey);
      }
    }
  }

  // function for deleting a private key for the org. Thsi will be deleted once
  // approval process is complete
  function deleteOrgKey(string _orgId, string _tmKey) external
  {
    bool ret = checkOrgPendingOp(_orgId);
    if(ret) {
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
      else {
        emit KeyNotFound(_tmKey);
      }
    }
  }

  // function for approving key add or delete operations
  function approvePendingOp(string _orgId) external canVote(_orgId)
  {
    if (checkingPendingOp(_orgId)){
      uint orgIndex = getOrgIndex(_orgId);
      processVote(_orgId);
      processApproval(orgIndex);
    }
    else {
      emit NothingToApprove(_orgId);
    }
  }

  // initialize the voter account votes to false. This will be called when a
  // new item is initiated for approval
  function voterInit(string _orgId) internal {
    uint orgIndex = getOrgIndex(_orgId);
    uint vorgIndex = getMasterOrgIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < masterOrgList[vorgIndex].voterAccount.length; i++){
      voteStatus[orgIndex][masterOrgList[vorgIndex].voterAccount[i]] = false;
    }
    voteCount[orgIndex] = 0;
  }

  // processes the vote from the voter account.
  function processVote (string _orgId) internal {
    uint orgIndex = getOrgIndex(_orgId);
    if (voteStatus[orgIndex][msg.sender] == false ){
      voteStatus[orgIndex][msg.sender] = true;
      voteCount[orgIndex]++;
    }
  }

  // checks if enough votes have been cast for the pending operation. If yes
  // returns true
  function checkEnoughVotes (string _orgId) internal view returns (bool) {
    uint orgIndex = getOrgIndex(_orgId);
    if (voteCount[orgIndex] > masterOrgList[orgIndex].voterAccount.length / 2 ){
      return true;
    }
    return false;
  }

  function checkKeyInUse(string _orgId, string _morgId, string _key) internal view returns (bool) {
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
  function masterKeyUpdate(string _orgId, string _morgId, string _key, Operation _op) internal {
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
        for (index = 0; i < masterOrgList[morgIndex].tmKey.length; index++) {
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
    if(checkEnoughVotes(orgList[_orgIndex].orgId)){
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
}
