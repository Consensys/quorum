pragma solidity ^0.5.3;
import "./PermissionsUpgradable.sol";


contract NodeManager {
    PermissionsUpgradable private permUpgradable;
    // enum and struct declaration
    // changing node status to integer (0-NotInList, 1- PendingApproval, 2-Approved,
    //      PendingDeactivation, Deactivated, PendingActivation, PendingBlacklisting, Blacklisted)
//    enum NodeStatus {NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated, PendingActivation, PendingBlacklisting, Blacklisted}
    struct NodeDetails {
        string enodeId; //e.g. 127.0.0.1:20005
        string orgId;
        uint status;
    }
    // use an array to store node details
    // if we want to list all node one day, mapping is not capable
    NodeDetails[] private nodeList;
    // use a mapping of enodeid to array index to track node
    mapping(bytes32 => uint) private nodeIdToIndex;
    // keep track of node number
    uint private numberOfNodes;


    // node permission events for new node propose
    event NodeProposed(string _enodeId);
    event NodeApproved(string _enodeId);

    // node permission events for node decativation
    event NodePendingDeactivation (string _enodeId);
    event NodeDeactivated(string _enodeId);

    // node permission events for node activation
    event NodePendingActivation(string _enodeId);
    event NodeActivated(string _enodeId);

    // node permission events for node blacklist
    event NodePendingBlacklist(string _enodeId);
    event NodeBlacklisted(string);

    modifier onlyImpl
    {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    // Checks if the given enode exists
    modifier enodeInList(string memory _enodeId)
    {
        require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0, "Enode is not in the list");
        _;
    }

    // Checks if the given enode does not exists
    modifier enodeNotInList(string memory _enodeId)
    {
        require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] == 0, "Enode is in the list");
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    // Get node details given enode Id
    function getNodeDetails(string memory enodeId) public view returns (string memory _enodeId, uint _nodeStatus)
    {
        uint nodeIndex = getNodeIndex(enodeId);
        return (nodeList[nodeIndex].enodeId, nodeList[nodeIndex].status);
    }
    // Get node details given index
    function getNodeDetailsFromIndex(uint nodeIndex) public view returns (string memory _orgId, string memory _enodeId, uint _nodeStatus)
    {
        return (nodeList[nodeIndex].orgId, nodeList[nodeIndex].enodeId, nodeList[nodeIndex].status);
    }
    // Get number of nodes
    function getNumberOfNodes() public view returns (uint)
    {
        return numberOfNodes;
    }

    // Get node status by enode id
    function getNodeStatus(string memory _enodeId) public view returns (uint)
    {
        if (nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] == 0){
            return 0;
        }
        return nodeList[getNodeIndex(_enodeId)].status;
    }

    function addNode(string calldata _enodeId, string calldata _orgId) external
    onlyImpl
    enodeNotInList(_enodeId)
    {
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _orgId,  1));
        emit NodeProposed(_enodeId);
    }

    function addOrgNode(string calldata _enodeId, string calldata _orgId) external
    onlyImpl
    enodeNotInList(_enodeId)
    {
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _orgId,  2));
        emit NodeApproved(_enodeId);
    }

    // Adds a node to the nodeList mapping and emits node added event if successfully and node exists event of node is already present
    function approveNode(string calldata _enodeId) external
    onlyImpl
    {
        require(getNodeStatus(_enodeId) == 1, "Node need to be in PendingApproval status");
        uint nodeIndex = getNodeIndex(_enodeId);
        // vote node
        nodeList[nodeIndex].status = 2;
        emit NodeApproved(nodeList[nodeIndex].enodeId);
    }

//    // Propose a node for deactivation from network
//    function proposeDeactivation(string calldata _enodeId) external enodeInList(_enodeId)
//    {
//        require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
//        uint nodeIndex = getNodeIndex(_enodeId);
//        nodeList[nodeIndex].status = NodeStatus.PendingDeactivation;
//        emit NodePendingDeactivation(_enodeId);
//
//    }
//
//    //deactivates a given Enode and emits the decativation event
//    function deactivateNode(string calldata _enodeId) external
//    {
//        require(getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation, "Node need to be in PendingDeactivation status");
//        uint nodeIndex = getNodeIndex(_enodeId);
//        nodeList[nodeIndex].status = NodeStatus.Deactivated;
//        emit NodeDeactivated(nodeList[nodeIndex].enodeId);
//
//    }
//
//    // Propose node for blacklisting
//    function proposeNodeActivation(string calldata _enodeId) external
//    {
//        require(getNodeStatus(_enodeId) == NodeStatus.Deactivated, "Node need to be in Deactivated status");
//        uint nodeIndex = getNodeIndex(_enodeId);
//        nodeList[nodeIndex].status = NodeStatus.PendingActivation;
//        // emit event
//        emit NodePendingActivation(_enodeId);
//    }

//    //deactivates a given Enode and emits the decativation event
//    function activateNode(string calldata _enodeId) external
//    {
//        require(getNodeStatus(_enodeId) == NodeStatus.PendingActivation, "Node need to be in PendingActivation status");
//        uint nodeIndex = getNodeIndex(_enodeId);
//        require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
//        // vote node
//        updateVoteStatus(nodeIndex);
//        // emit event
//        // check if node vote reach majority
//        if (checkEnoughVotes(nodeIndex)) {
//            nodeList[nodeIndex].status = NodeStatus.Approved;
//            emit NodeActivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
//        }
//    }
//
//    // Propose node for blacklisting
//    function proposeNodeBlacklisting(string calldata _enodeId, string calldata _ipAddrPort, string calldata _discPort, string calldata _raftPort) external
//    {
//        if (checkVotingAccountExist()) {
//            uint nodeIndex = getNodeIndex(_enodeId);
//            // check if node is in the nodeList
//            if (nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0) {
//                // no matter what status the node is in, vote will reset and node status change to PendingBlacklisting
//                nodeList[nodeIndex].status = NodeStatus.PendingBlacklisting;
//                nodeIndex = getNodeIndex(_enodeId);
//            } else {
//                // increment node number, add node to the list
//                numberOfNodes++;
//                nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
//                nodeList.push(NodeDetails(_enodeId, _ipAddrPort, _discPort, _raftPort, NodeStatus.PendingBlacklisting));
//                nodeIndex = numberOfNodes;
//            }
//            // add voting status, numberOfNodes is the index of current proposed node
//            initNodeVoteStatus(nodeIndex);
//            // emit event
//            emit NodePendingBlacklist(_enodeId);
//        }
//    }
//
//    //Approve node blacklisting
//    function blacklistNode(string calldata _enodeId) external
//    {
//        require(getNodeStatus(_enodeId) == NodeStatus.PendingBlacklisting, "Node need to be in PendingBlacklisting status");
//        uint nodeIndex = getNodeIndex(_enodeId);
//        require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
//        // vote node
//        voteStatus[nodeIndex][msg.sender] = true;
//        voteCount[nodeIndex]++;
//        // emit event
//        // check if node vote reach majority
//        if (checkEnoughVotes(nodeIndex)) {
//            nodeList[nodeIndex].status = NodeStatus.Blacklisted;
//            emit NodeBlacklisted(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
//        }
//    }

    /* private functions */
    function getNodeIndex(string memory _enodeId) internal view returns (uint)
    {
        return nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] - 1;
    }


}
