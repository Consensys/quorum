pragma solidity ^0.5.3;
import "./PermissionsUpgradable.sol";


contract NodeManager {
    PermissionsUpgradable private permUpgradable;
    // enum and struct declaration
    // changing node status to integer (0-NotInList, 1- PendingApproval, 2-Approved, 3-Deactivated, 4-Blacklisted)
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
    event NodeDeactivated(string _enodeId);

    // node permission events for node activation
    event NodeActivated(string _enodeId);

    // node permission events for node blacklist
    event NodeBlacklisted(string _enodeId);

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
    function getNodeDetails(string memory enodeId) public view returns (string memory _orgId, string memory _enodeId, uint _nodeStatus)
    {
        uint nodeIndex = getNodeIndex(enodeId);
        return (nodeList[nodeIndex].orgId, nodeList[nodeIndex].enodeId, nodeList[nodeIndex].status);
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

    function addAdminNode(string calldata _enodeId, string calldata _orgId) external
    onlyImpl
    enodeNotInList(_enodeId)
    {
        addNode(_enodeId, _orgId);
        approveNode(_enodeId, _orgId);
    }
    function addNode(string memory _enodeId, string memory _orgId) public
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
    function approveNode(string memory _enodeId, string memory _orgId) public
    onlyImpl
    enodeInList(_enodeId)
    {
        // node should belong to the passed org
        require(checkOrg(_enodeId, _orgId), "Node does not belong to the org");
        require(getNodeStatus(_enodeId) == 1, "Node need to be in PendingApproval status");
        uint nodeIndex = getNodeIndex(_enodeId);
        // vote node
        nodeList[nodeIndex].status = 2;
        emit NodeApproved(nodeList[nodeIndex].enodeId);
    }

    function updateNodeStatus(string calldata _enodeId, string calldata _orgId, uint _status) external
    onlyImpl
    enodeInList(_enodeId)
    {
        // node should belong to the org
        require(checkOrg(_enodeId, _orgId), "Node does not belong to the org");
        // changing node status to integer (0-NotInList, 1- PendingApproval, 2-Approved, 3-Deactivated, 4-Blacklisted)
        // operations that can be done 3-Deactivate Node, 4-ActivateNode, 5-Blacklist nodeList
        require((_status == 3 || _status == 4 || _status == 5), "invalid operation");

        if (_status == 3){
            require(getNodeStatus(_enodeId) == 2, "Op cannot be performed");
            nodeList[getNodeIndex(_enodeId)].status = 3;
            emit NodeDeactivated(_enodeId);
        }
        else if (_status == 4){
            require(getNodeStatus(_enodeId) == 3, "Op cannot be performed");
            nodeList[getNodeIndex(_enodeId)].status = 2;
            emit NodeActivated(_enodeId);
        }
        else {
            nodeList[getNodeIndex(_enodeId)].status = 5;
            emit NodeBlacklisted(_enodeId);
        }
    }

    /* private functions */
    function getNodeIndex(string memory _enodeId) internal view
    returns (uint)
    {
        return nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] - 1;
    }

    function checkOrg(string memory _enodeId, string memory _orgId) internal view
    returns(bool)
    {
        return (keccak256(abi.encodePacked(nodeList[getNodeIndex(_enodeId)].orgId)) == keccak256(abi.encodePacked(_orgId)));
    }

}
