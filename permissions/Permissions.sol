pragma solidity ^0.4.23;

contract Permissions {

  enum NodeStatus { NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated }

  enum AccountAccess {FullAccess, ReadOnly, Transact, ContractDeploy}

  struct nodeDetails {
    string enodeId;
    //e.g. 127.0.0.1:20005
    string ipAddrPort;
    string discPort;
    string raftPort;
    bool canWrite;
    bool canLead;
    NodeStatus status;
  }
  mapping (bytes32 => nodeDetails) nodeList;

  struct acctAccess {
    address acctId;
    AccountAccess access;
  }
  mapping (address => acctAccess) acctAccessList;

  event NewNodeProposed (string _enodeId);
  event NodeApproved(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event NodePendingDeactivation (string _enodeId);
  event NodeDeactivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event AcctAccessModified (address acctId, AccountAccess access);

  // Checks if the Node is already added. If yes then returns true
  function updateAcctAccess (address _acctId, AccountAccess access) public {
    acctAccessList[_acctId] = acctAccess(_acctId, access);
    emit AcctAccessModified(_acctId, access);
  }
  // Checks if the Node is already added. If yes then returns true
  function getNodeStatus (string _enodeId) public view returns (NodeStatus _status) {
    return nodeList[keccak256(_enodeId)].status;
  }

  // Adds a node to the nodeList mapping and emits node added event if successfully and node exists event of node is already present
  function ApproveNode(string _enodeId) public {
    require(getNodeStatus(_enodeId) != NodeStatus.NotInList, "Node is already in the list");

    bytes32 i;
    i = keccak256(_enodeId);
    nodeList[i].status = NodeStatus.Approved;
    emit NodeApproved(nodeList[i].enodeId, nodeList[i].ipAddrPort, nodeList[i].discPort, nodeList[i].raftPort);
  }

  function ProposeNode(string _enodeId, bool _canWrite, bool _canLead, string _ipAddrPort, string _discPort, string _raftPort) public {
    require(getNodeStatus(_enodeId) == NodeStatus.NotInList, "New node cannot be in the list");
    nodeList[keccak256(_enodeId)] = nodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort,  _canWrite, _canLead, NodeStatus.PendingApproval);
    emit NewNodeProposed (_enodeId);
  }

  function ProposeDeactivation(string _enodeId) public {
    require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
    nodeList[keccak256(_enodeId)].status = NodeStatus.PendingDeactivation;
    emit NodePendingDeactivation(_enodeId);
  }

  //deactivates a given Enode and emits the decativation event
  function DeactivateNode (string _enodeId) public {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation, "Node need to be in PendingDeactivation status");

    bytes32 i;
    i = keccak256(_enodeId);
    nodeList[i].status = NodeStatus.Deactivated;
    emit NodeDeactivated(nodeList[i].enodeId, nodeList[i].ipAddrPort, nodeList[i].discPort, nodeList[i].raftPort);
  }

}
