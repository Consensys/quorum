pragma solidity ^0.4.23;

contract Clusterkeys {

  struct OrgDetails {
    string orgId;
    string privateKeys;
  }
  OrgDetails [] private orgList;

  mapping(bytes32 => uint) private OrgIdIndex;

  uint private orgNumber;

  event OrgKeyUpdated(string _orgId, string _privateKeys);

  function updatedOrgKeys (string _orgId, string _privateKeys) external
  {
    emit OrgKeyUpdated(_orgId, _privateKeys);
  }
}
