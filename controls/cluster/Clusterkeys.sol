pragma solidity ^0.4.23;

contract Clusterkeys {

  // Struct for managing the org details
  enum Operation {None, Add, Delete}
  struct OrgDetails {
    string orgId;
    string morgId;
    string [] privateKey;
    string pendingKey;
    Operation pendingOp;
  }
  OrgDetails [] private orgList;
  mapping(bytes32 => uint) private OrgIndex;

  // Struct for managing the voter accounst for the org 
  struct MasterOrgDetails {
    string orgId;
    address [] orgVoterAccount;
    string [] privateKey;
  }
  MasterOrgDetails [] private morgList;
  mapping(bytes32 => uint) private MasterOrgIndex;

  // mapping to monitor the voting status for each acount and
  // overall voting count
  mapping (uint => mapping (address => bool)) private voteStatus;
  mapping (uint => uint) private voteCount;

  uint private numberOfOrgs = 0;
  uint private orgVoterNum = 0;

  // events related to Org level key management
  event OrgKeyAdded(string _orgId, string _privateKey);
  event OrgKeyDeleted(string _orgId, string _privateKey);
  event KeyNotFound(string _privateKey);
  event KeyExists(string _orgId, string _privateKey);
  event OrgNotFound(string _orgId);

  // events related to org level approval process
  event PendingApproval(string _orgId);
  event ItemForApproval(string _orgId, Operation _pendingOp, string _privateKey);
  event NothingToApprove(string _orgId);

  // events related to managing voting accounts for the org
  event NoVotingAccount(string _orgId);
  event VoterAdded(string _orgId, address _address);
  event VoterNotFound(string _orgId, address _address);
  event VoterAccountDeleted(string _orgId, address _address);
  event VoterExists(string _orgId, address _address);

  // events related to helper functions to print all org keys and voter keys
  /* event PrintAll(string _orgId, string _privateKey); */
  /* event PrintVoter(string _orgId, address _voterAccount); */

  // returns the org index for the org list
  function getOrgIndex(string _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // returns the voter index for the org from voter list
  function getMasterIndex(string _orgId) internal view returns (uint)
  {
    return MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  // checks if the sender is one of the registered voter account for the org
  modifier canVote(string _orgId){
    bool flag = false;
    uint orgIndex = getOrgIndex(_orgId);
    uint morgIndex = getMasterIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < morgList[morgIndex].orgVoterAccount.length; i++){
      if ( morgList[morgIndex].orgVoterAccount[i] == msg.sender){
        flag = true;
        break;
      }
    }
    require(flag, "Account cannot vote");
    _;
  }

  // checks if the org has any voter accounts set up or not
  function checkIfVoterExists(string _orgId, address _address) internal view returns (bool, uint){
    bool keyExists = false;
    uint orgIndex = getOrgIndex(_orgId);
    uint voterIndex = getMasterIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < morgList[voterIndex].orgVoterAccount.length; i++){
      if(keccak256(abi.encodePacked(morgList[voterIndex].orgVoterAccount[i])) == keccak256(abi.encodePacked(_address))){
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
    if (MasterOrgIndex[keccak256(abi.encodePacked(orgList[orgIndex].morgId))] == 0){
      emit NoVotingAccount(_orgId);
      return false;
    }
    uint morgIndex = getMasterIndex(orgList[orgIndex].morgId);
    if (morgList[morgIndex].orgVoterAccount.length == 0) {
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
  function checkIfKeyExists(string _orgId, string _privateKey) internal view returns (bool, uint){
    bool keyExists = false;
    uint orgIndex = getOrgIndex(_orgId);
    for (uint i = 0; i < orgList[orgIndex].privateKey.length; i++){
      if(keccak256(abi.encodePacked(orgList[orgIndex].privateKey[i])) == keccak256(abi.encodePacked(_privateKey))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }

  // function for adding a voter account to a master org
  function addVoter(string _orgId, address _address) external
  {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      orgVoterNum++;
      MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] = orgVoterNum;
      morgList.push( MasterOrgDetails(_orgId, new address[](0), new string[](0)));
      morgList[orgVoterNum - 1].orgVoterAccount.push(_address);
      emit VoterAdded(_orgId, _address);
    }
    else {
      bool voterExists = false;
      uint i = 0;
      (voterExists, i) = checkIfVoterExists(_orgId, _address);
      if (voterExists) {
        emit VoterExists(_orgId, _address);
      }
      else {
        uint voterIndex = getMasterIndex(_orgId);
        morgList[voterIndex].orgVoterAccount.push(_address);
        emit VoterAdded(_orgId, _address);
      }
    }
  }

  // function for deleting a voter account to a master org
  function deleteVoter(string _orgId, address _address) external
  {
    if (MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      emit OrgNotFound(_orgId);
    }
    else {
      uint voterIndex = getMasterIndex(_orgId);
      (bool voterExists, uint i) = checkIfVoterExists(_orgId, _address);

      if (voterExists == true) {
        for (uint j = i; j <  morgList[voterIndex].orgVoterAccount.length -1; j++){
          morgList[voterIndex].orgVoterAccount[j] = morgList[voterIndex].orgVoterAccount[j+1];
        }
        delete morgList[voterIndex].orgVoterAccount[morgList[voterIndex].orgVoterAccount.length -1];
        morgList[voterIndex].orgVoterAccount.length --;
        emit VoterAccountDeleted(_orgId, _address);
      }
      else {
        emit VoterNotFound(_orgId, _address);
      }
    }
  }

  // function to create master org
  function addMasterOrg(string _orgId) external
  {
    require (MasterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0);


  }

  // function for adding a private key for the org. Thsi will be added once
  // approval process is complete
  function addOrgKey(string _orgId, string _morgId, string _privateKey) external
  {
    if (checkVotingAccountExists(_orgId)){
      if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
        numberOfOrgs++;
        OrgIndex[keccak256(abi.encodePacked(_orgId))] = numberOfOrgs;
        orgList.push( OrgDetails(_orgId, _morgId, new string[](0), _privateKey, Operation.Add));
        voterInit(_orgId);
        emit ItemForApproval(_orgId, Operation.Add, _privateKey);
      }
      else {
        if (checkingPendingOp(_orgId)){
          emit PendingApproval(_orgId);
        }
        else {
          bool keyExists = false;
          uint i = 0;
          (keyExists, i) = checkIfKeyExists(_orgId, _privateKey);
          if (keyExists) {
            emit KeyExists(_orgId, _privateKey);
          }
          else {
            uint orgIndex;
            orgIndex = getOrgIndex(_orgId);
            orgList[orgIndex].pendingKey = _privateKey;
            orgList[orgIndex].pendingOp = Operation.Add;
            voterInit(_orgId);
            emit ItemForApproval(_orgId,Operation.Add,  _privateKey);
          }
        }
      }
    }
  }

  // function for deleting a private key for the org. Thsi will be deleted once
  // approval process is complete
  function deleteOrgKey(string _orgId, string _privateKey) external
  {
    if (checkVotingAccountExists(_orgId)){
      if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
        emit OrgNotFound(_orgId);
      }
      else {
        if (checkingPendingOp(_orgId)){
          emit PendingApproval(_orgId);
        }
        else {
          uint orgIndex = getOrgIndex(_orgId);
          uint i = 0;
          bool keyExists = false;

          (keyExists, i) = checkIfKeyExists (_orgId, _privateKey);
          if (keyExists == true) {
            orgList[orgIndex].pendingKey = _privateKey;
            orgList[orgIndex].pendingOp = Operation.Delete;
            voterInit(_orgId);
            emit ItemForApproval(_orgId, Operation.Delete,  _privateKey);

          }
          else {
            emit KeyNotFound(_privateKey);
          }
        }
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
    uint morgIndex = getMasterIndex(orgList[orgIndex].morgId);
    for (uint i = 0; i < morgList[morgIndex].orgVoterAccount.length; i++){
      voteStatus[orgIndex][morgList[morgIndex].orgVoterAccount[i]] = false;
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
    if (voteCount[orgIndex] > morgList[orgIndex].orgVoterAccount.length / 2 ){
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
        for (uint j = 0; j < orgList[i].privateKey.length; j++){
          if (keccak256(abi.encodePacked(orgList[i].privateKey[j])) == keccak256(abi.encodePacked(_key))){
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
    uint morgIndex = getMasterIndex(_morgId);
    if (_op == Operation.Add) {
      // check if the key is existing. if yes ignore else add to master list
      bool keyExists = false;
      for (uint i = 0; i < morgList[morgIndex].privateKey.length; i++){
        if(keccak256(abi.encodePacked(orgList[morgIndex].privateKey[i])) == keccak256(abi.encodePacked(_key))){
          keyExists = true;
          break;
        }
      }
      if (keyExists == false ){
        morgList[morgIndex].privateKey.push(_key);
      }
    }
    else {
      // the key can be deleted from master list only when none of the suborgs have the
      // key in the private keys
      if (!(checkKeyInUse(_orgId, _morgId, _key))){
        uint index;
        for (index = 0; i < morgList[morgIndex].privateKey.length; index++) {
          if(keccak256(abi.encodePacked(morgList[morgIndex].privateKey[index])) == keccak256(abi.encodePacked(_key))) {
            break;
          }
        }
        for (uint j = index; j <  morgList[morgIndex].privateKey.length -1; j++){
          morgList[morgIndex].privateKey[j] = morgList[morgIndex].privateKey[j+1];
        }
        delete morgList[morgIndex].privateKey[morgList[morgIndex].privateKey.length -1];
        morgList[morgIndex].privateKey.length --;
      }
    }
  }

  // function to process the approavl for add or delete
  function processApproval(uint _orgIndex) internal {
    if(checkEnoughVotes(orgList[_orgIndex].orgId)){
      string storage locKey = orgList[_orgIndex].pendingKey;
      if (orgList[_orgIndex].pendingOp == Operation.Add){
        orgList[_orgIndex].privateKey.push(orgList[_orgIndex].pendingKey);
        masterKeyUpdate(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId, orgList[_orgIndex].pendingKey, Operation.Add);
        emit OrgKeyAdded(orgList[_orgIndex].orgId, locKey);
      }
      else {
        bool keyExists = false;
        uint i = 0;
        (keyExists, i) = checkIfKeyExists (orgList[_orgIndex].orgId, locKey);
        for (uint j = i; j <  orgList[_orgIndex].privateKey.length -1; j++){
          orgList[_orgIndex].privateKey[j] = orgList[_orgIndex].privateKey[j+1];
        }
        delete orgList[_orgIndex].privateKey[orgList[_orgIndex].privateKey.length -1];
        orgList[_orgIndex].privateKey.length --;
        masterKeyUpdate(orgList[_orgIndex].orgId, orgList[_orgIndex].morgId, orgList[_orgIndex].pendingKey, Operation.Delete);
        emit OrgKeyDeleted(orgList[_orgIndex].orgId, locKey);
      }
      orgList[_orgIndex].pendingOp = Operation.None;
      orgList[_orgIndex].pendingKey = "";
    }
  }

  /* // helper function to print all privates keys for an org */
  /* function printAllOrg () public { */
  /*   for (uint i = 0; i < orgList.length; i++){ */
  /*     for (uint j = 0; j < orgList[i].privateKey.length ; j++){ */
  /*       emit PrintAll(orgList[i].orgId, orgList[i].privateKey[j]); */
  /*     } */
  /*   } */
  /* } */

  /* // helper function to print all voters accounts for an org */
  /* function printAllVoter () public { */
  /*   for (uint i = 0; i < morgList.length; i++){ */
  /*     for (uint j = 0; j < morgList[i].orgVoterAccount.length ; j++){ */
  /*       emit PrintVoter(morgList[i].orgId, morgList[i].orgVoterAccount[j]); */
  /*     } */
  /*   } */
  /* } */

}
