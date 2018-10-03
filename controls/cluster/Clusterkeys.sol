pragma solidity ^0.4.23;

contract Clusterkeys {

  enum Operation {None, Add, Delete}
  struct OrgDetails {
    string orgId;
    string [] privateKey;
    string pendingKey;
    Operation pendingOp;
  }
  OrgDetails [] private orgList;

  mapping(bytes32 => uint) private OrgIndex;

  struct OrgVoterDetails {
    string orgId;
    address [] orgVoterAccount;
  }
  OrgVoterDetails [] private voterList;
  mapping(bytes32 => uint) private VoterOrgIndex;

  mapping (uint => mapping (address => bool)) private voteStatus;
  mapping (uint => uint) private voteCount;

  uint private numberOfOrgs = 0;

  uint private orgVoterNum = 0;

  event OrgKeyAdded(string _orgId, string _privateKey);
  event OrgKeyDeleted(string _orgId, string _privateKey);
  event orgVoterAdded(string _orgId, string _voterAccount);
  event KeyNotFound(string _privateKey);
  event OrgNotFound(string _orgId);
  event KeyExists(string _orgId, string _privateKey);
  event VoterAdded(string _orgId, address _address);
  event VoterExists(string _orgId, address _address);
  event VoterNotFound(string _orgId, address _address);
  event VoterAccountDeleted(string _orgId, address _address);
  event NoVotingAccount(string _orgId);
  event PendingApproval(string _orgId);
  event ItemForApproval(string _orgId, Operation _pendingOp, string _privateKey);
  event NothingToApprove(string _orgId);

  event PrintAll(string _orgId, string _privateKey);
  event PrintVoter(string _orgId, address _voterAccount);
  event PrintKey(string _orgId, Operation _pendingOp, string _pendingKey);

  modifier canVote(string _orgId){
    bool flag = false;
    uint orgIndex = getOrgIndexVoter(_orgId);
    for (uint i = 0; i < voterList[orgIndex].orgVoterAccount.length; i++){
      if ( voterList[orgIndex].orgVoterAccount[i] == msg.sender){
        flag = true;
        break;
      }
    }
    require(flag, "Account cannot vote");
    _;
  }

  function getOrgIndex(string _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  function getOrgIndexVoter(string _orgId) internal view returns (uint)
  {
    return VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  function checkIfVoterExists(string _orgId, address _address) internal view returns (bool, uint){
    bool keyExists = false;
    uint voterIndex = getOrgIndexVoter(_orgId);
    for (uint i = 0; i < voterList[voterIndex].orgVoterAccount.length; i++){
      if(keccak256(abi.encodePacked(voterList[voterIndex].orgVoterAccount[i])) == keccak256(abi.encodePacked(_address))){
        keyExists = true;
        break;
      }
    }
    return (keyExists, i);
  }

  function checkVotingAccountExists(string _orgId) internal returns (bool)
  {
    if (VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0){
      emit NoVotingAccount(_orgId);
      return false;
    }
    uint orgIndex = getOrgIndexVoter(_orgId);
    if (voterList[orgIndex].orgVoterAccount.length == 0) {
      emit NoVotingAccount(_orgId);
      return false;
    }
    return true;
  }

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

  function addVoter(string _orgId, address _address) external
  {
    if (VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      orgVoterNum++;
      VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] = orgVoterNum;
      voterList.push( OrgVoterDetails(_orgId, new address[](0)));
      voterList[orgVoterNum - 1].orgVoterAccount.push(_address);
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
        uint voterIndex = getOrgIndexVoter(_orgId);
        voterList[voterIndex].orgVoterAccount.push(_address);
        emit VoterAdded(_orgId, _address);
      }
    }
  }

  function deleteVoter(string _orgId, address _address) external
  {
    if (VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      emit OrgNotFound(_orgId);
    }
    else {
      uint voterIndex = getOrgIndexVoter(_orgId);
      //      uint i = 0;
      //bool keyExists = false;

      (bool voterExists, uint i) = checkIfVoterExists(_orgId, _address);

      if (voterExists == true) {
        for (uint j = i; j <  voterList[voterIndex].orgVoterAccount.length -1; j++){
          voterList[voterIndex].orgVoterAccount[j] = voterList[voterIndex].orgVoterAccount[j+1];
        }
        delete voterList[voterIndex].orgVoterAccount[voterList[voterIndex].orgVoterAccount.length -1];
        voterList[voterIndex].orgVoterAccount.length --;
        emit VoterAccountDeleted(_orgId, _address);
      }
      else {
        emit VoterNotFound(_orgId, _address);
      }
    }
  }

  function addOrgKey(string _orgId, string _privateKey) external
  {
    if (checkVotingAccountExists(_orgId)){
      if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
        numberOfOrgs++;
        OrgIndex[keccak256(abi.encodePacked(_orgId))] = numberOfOrgs;
        orgList.push( OrgDetails(_orgId, new string[](0), _privateKey, Operation.Add));
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
            //          orgList[orgIndex].privateKey.push(_privateKey);
            orgList[orgIndex].pendingKey = _privateKey;
            orgList[orgIndex].pendingOp = Operation.Add;
            voterInit(_orgId);
            emit ItemForApproval(_orgId,Operation.Add,  _privateKey);
          }
        }
      }
    }
  }

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

  function voterInit(string _orgId) internal {
    uint orgIndex = getOrgIndexVoter(_orgId);
    for (uint i = 0; i < voterList[orgIndex].orgVoterAccount.length; i++){
      voteStatus[orgIndex][voterList[orgIndex].orgVoterAccount[i]] = false;
    }
    voteCount[orgIndex] = 0;
  }

  function processVote (string _orgId) internal {
    uint orgIndex = getOrgIndexVoter(_orgId);
    if (voteStatus[orgIndex][msg.sender] == false ){
      voteStatus[orgIndex][msg.sender] = true;
      voteCount[orgIndex]++;
    }
  }

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

  function checkEnoughVotes (string _orgId) internal view returns (bool) {
    uint orgIndex = getOrgIndexVoter(_orgId);
    if (voteCount[orgIndex] > voterList[orgIndex].orgVoterAccount.length / 2 ){
      return true;
    }
    return false;
  }

  function processApproval(uint _orgIndex) internal {
    if(checkEnoughVotes(orgList[_orgIndex].orgId)){
      string storage locKey = orgList[_orgIndex].pendingKey;
      if (orgList[_orgIndex].pendingOp == Operation.Add){
        orgList[_orgIndex].privateKey.push(orgList[_orgIndex].pendingKey);
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
        emit OrgKeyDeleted(orgList[_orgIndex].orgId, locKey);
      }
      orgList[_orgIndex].pendingOp = Operation.None;
      orgList[_orgIndex].pendingKey = "";
    }
  }

  function printAllOrg () public {
    for (uint i = 0; i < orgList.length; i++){
      for (uint j = 0; j < orgList[i].privateKey.length ; j++){
        emit PrintAll(orgList[i].orgId, orgList[i].privateKey[j]);
      }
    }
  }

  function printAllVoter () public {
    for (uint i = 0; i < voterList.length; i++){
      for (uint j = 0; j < voterList[i].orgVoterAccount.length ; j++){
        emit PrintVoter(voterList[i].orgId, voterList[i].orgVoterAccount[j]);
      }
    }
  }

}
