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
    string [] orgVoterAccount;
  }
  mapping(bytes32 => uint) private OrgVoterIndex;

  uint private numberOfOrgs = 0;

  uint private orgNumber = 0;

  event OrgKeyAdded(string _orgId, string _privateKey);
  event OrgKeyDeleted(string _orgId, string _privateKey);
  event orgVoterAdded(string _orgId, string _voterAccount);
  event KeyNotFound(string _privateKey);
  event OrgNotFound(string _orgId);
  event PrintAll(string _orgId, string _privateKey);
  event KeyExists(string _orgId, string _privateKey);
  event Dummy(uint _orgId, bool _keyExists, uint loopCnt );

  function checkIfKeyExists(string _orgId, string _privateKey) internal view returns (bool){
    bool keyExists = false;
    uint locOrgId = getOrgIndex(_orgId);
    for (uint i = 0; i < orgList[locOrgId].privateKey.length; i++){
      if(keccak256(abi.encodePacked(orgList[locOrgId].privateKey[i])) == keccak256(abi.encodePacked(_privateKey))){
        keyExists = true;
        break;
      }
    }
    return keyExists;
  }

  function getOrgIndex(string _orgId) internal view returns (uint)
  {
    return OrgIndex[keccak256(abi.encodePacked(_orgId))] - 1;
  }

  function addOrgKey(string _orgId, string _privateKey) external
  {
    if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      numberOfOrgs++;
      OrgIndex[keccak256(abi.encodePacked(_orgId))] = numberOfOrgs;
      orgList.push( OrgDetails(_orgId, new string[](0)));
      orgList[numberOfOrgs-1].privateKey.push(_privateKey);
      emit OrgKeyAdded(_orgId, _privateKey);
    }
    else {
      if (checkIfKeyExists (_orgId, _privateKey)) {
        emit KeyExists(_orgId, _privateKey);
      }
      else {
        uint locOrgId;
        locOrgId = getOrgIndex(_orgId);
        orgList[locOrgId].privateKey.push(_privateKey);
        emit OrgKeyAdded(_orgId, _privateKey);
      }
    }
  }

  function deleteOrgKey(string _orgId, string _privateKey) external
  {
    if (OrgIndex[keccak256(abi.encodePacked(_orgId))] == 0) {
      emit OrgNotFound(_orgId);
    }
    else {
      uint locOrgId = getOrgIndex(_orgId);
      uint i = 0;
      bool keyExists = false;

      for (i = 0; i <= orgList[locOrgId].privateKey.length -1; i++){
        if(keccak256(abi.encodePacked(orgList[locOrgId].privateKey[i])) == keccak256(abi.encodePacked(_privateKey))){
          keyExists = true;
          break;
        }
      }
      if (keyExists == true) {
        for (uint j = i; j <  orgList[locOrgId].privateKey.length -1; j++){
          orgList[locOrgId].privateKey[j] = orgList[locOrgId].privateKey[j+1];
        }
        delete orgList[locOrgId].privateKey[orgList[locOrgId].privateKey.length -1];
        orgList[locOrgId].privateKey.length --;
        emit OrgKeyDeleted(_orgId, _privateKey);
      }
      else {
        emit KeyNotFound(_privateKey);
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
