pragma solidity ^0.4.23;

contract Clusterkeys {

  struct OrgDetails {
    string orgId;
    string [] privateKey;
  }
  OrgDetails [] private orgList;

  mapping(bytes32 => uint) private OrgIndex;

  struct OrgVoterDetails {
    string orgId;
    address [] orgVoterAccount;
  }
  OrgVoterDetails [] private voterList;
  mapping(bytes32 => uint) private VoterOrgIndex;

  uint private numberOfOrgs = 0;

  uint private orgVoterNum = 0;

  event OrgKeyAdded(string _orgId, string _privateKey);
  event OrgKeyDeleted(string _orgId, string _privateKey);
  event orgVoterAdded(string _orgId, string _voterAccount);
  event KeyNotFound(string _privateKey);
  event OrgNotFound(string _orgId);
  event PrintAll(string _orgId, string _privateKey);
  event KeyExists(string _orgId, string _privateKey);
  event Dummy(uint _orgId, bool _keyExists, uint loopCnt );
  event VoterAdded(string _orgId, address _address);
  event VoterExists(string _orgId, address _address);
  event VoterNotFound(string _orgId, address _address);
  event VoterAccountDeleted(string _orgId, address _address);
  event NoVotingAccount(string _orgId);

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

  function getOrgIndex(string _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  function getOrgIndexVoter(string _orgId) internal view returns (uint)
  {
    return VoterOrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
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
        orgList.push( OrgDetails(_orgId, new string[](0)));
        orgList[numberOfOrgs-1].privateKey.push(_privateKey);
        emit OrgKeyAdded(_orgId, _privateKey);
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
          orgList[orgIndex].privateKey.push(_privateKey);
          emit OrgKeyAdded(_orgId, _privateKey);
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
        uint orgIndex = getOrgIndex(_orgId);
        uint i = 0;
        bool keyExists = false;

        (keyExists, i) = checkIfKeyExists (_orgId, _privateKey);

        if (keyExists == true) {
          for (uint j = i; j <  orgList[orgIndex].privateKey.length -1; j++){
            orgList[orgIndex].privateKey[j] = orgList[orgIndex].privateKey[j+1];
          }
          delete orgList[orgIndex].privateKey[orgList[orgIndex].privateKey.length -1];
          orgList[orgIndex].privateKey.length --;
          emit OrgKeyDeleted(_orgId, _privateKey);
        }
        else {
          emit KeyNotFound(_privateKey);
        }
      }
    }
  }

  function printAll () public {
    for (uint i = 0; i < orgList.length; i++){
      for (uint j = 0; j < orgList[i].privateKey.length ; j++){
        emit PrintAll(orgList[i].orgId, orgList[i].privateKey[j]);
      }
    }
  }

}
