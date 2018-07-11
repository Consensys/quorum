pragma solidity ^0.4.23;

contract Permissions {

  enum NodeStatus { NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated }

  struct nodeDetails {
    string enodeId;
    bool canWrite;
    bool canLead;
    NodeStatus status;
  }

  mapping (bytes32 => nodeDetails) nodeList;

  event NewNodeProposed (string _enodeId, bool _canWrite, bool _canLead);
  event NodeApproved(string _enodeId);
  event NodePendingDeactivation (string _enodeId);
  event NodeDeactivated (string _enodeId);

  // Checks if the Node is already added. If yes then returns true
  function getNodeStatus (string _enodeId) public view returns (NodeStatus _status) {
    return nodeList[keccak256(_enodeId)].status;
  }

  // Adds a node to the nodeList mapping and emits node added event if successfully and node exists event of node is already present
  function ApproveNode(string _enodeId) public {
    require(getNodeStatus(_enodeId) != NodeStatus.NotInList, "Node is already in the list");
    nodeList[keccak256(_enodeId)].status = NodeStatus.Approved;
    emit NodeApproved(_enodeId);
  }

  function ProposeNode(string _enodeId, bool _canWrite, bool _canLead) public {
    require(getNodeStatus(_enodeId) == NodeStatus.NotInList, "New node cannot be in the list");
    nodeList[keccak256(_enodeId)] = nodeDetails(_enodeId, _canWrite, _canLead, NodeStatus.PendingApproval);
    emit NewNodeProposed (_enodeId, _canWrite, _canLead);
  }

  function ProposeDeactivation(string _enodeId) public {
    require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
    nodeList[keccak256(_enodeId)].status = NodeStatus.PendingDeactivation;
    emit NodePendingDeactivation(_enodeId);
  }

  //deactivates a given Enode and emits the decativation event
  function DeactivateNode (string _enodeId) public {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation, "Node need to be in PendingDeactivation status");
    nodeList[keccak256(_enodeId)].status = NodeStatus.Deactivated;
    emit NodeDeactivated(_enodeId);
  }

}
