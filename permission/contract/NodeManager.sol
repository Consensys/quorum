pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";
/// @title Node manager contract
/// @notice This contract holds implementation logic for all node management
/// @notice functionality. This can be called only by the implementation
/// @notice contract only. there are few view functions exposed as public and
/// @notice can be called directly. these are invoked by quorum for populating
/// @notice permissions data in cache

contract NodeManager {
    PermissionsUpgradable private permUpgradable;
    struct NodeDetails {
        string enodeId; //e.g. 127.0.0.1:20005
        string orgId;
        uint status;
    }
    // use an array to store node details
    // if we want to list all node one day, mapping is not capable
    NodeDetails[] private nodeList;
    // mapping of enodeid to array index to track node
    mapping(bytes32 => uint) private nodeIdToIndex;
    // tracking total number of nodes in network
    uint private numberOfNodes;


    // node permission events for new node propose
    event NodeProposed(string _enodeId, string _orgId);
    event NodeApproved(string _enodeId, string _orgId);

    // node permission events for node deactivation
    event NodeDeactivated(string _enodeId, string _orgId);

    // node permission events for node activation
    event NodeActivated(string _enodeId, string _orgId);

    // node permission events for node blacklist
    event NodeBlacklisted(string _enodeId, string _orgId);

    /// @notice checks if the caller is implementation contract
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl());
        _;
    }

    /// @notice  checks if the node exists in the network
    /// @param _enodeId full enode id
    modifier enodeExists(string memory _enodeId) {
        require(nodeIdToIndex[keccak256(abi.encode(_enodeId))] != 0,
            "passed enode id does not exist");
        _;
    }

    /// @notice  checks if the node does not exist in the network
    /// @param _enodeId full enode id
    modifier enodeDoesNotExists(string memory _enodeId) {
        require(nodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0,
            "passed enode id exists");
        _;
    }

    /// @notice constructor. sets the permissions upgradable address
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /// @notice fetches the node details given an enode id
    /// @param _enodeId full enode id
    /// @return org id
    /// @return enode id
    /// @return status of the node
    function getNodeDetails(string calldata enodeId) external view
    returns (string memory _orgId, string memory _enodeId, uint _nodeStatus) {
        uint nodeIndex = _getNodeIndex(enodeId);
        return (nodeList[nodeIndex].orgId, nodeList[nodeIndex].enodeId,
        nodeList[nodeIndex].status);
    }

    /// @notice fetches the node details given the index of the enode
    /// @param _nodeIndex node index
    /// @return org id
    /// @return enode id
    /// @return status of the node
    function getNodeDetailsFromIndex(uint _nodeIndex) external view
    returns (string memory _orgId, string memory _enodeId, uint _nodeStatus) {
        return (nodeList[_nodeIndex].orgId, nodeList[_nodeIndex].enodeId,
        nodeList[_nodeIndex].status);
    }

    /// @notice returns the total number of enodes in the network
    /// @return number of nodes
    function getNumberOfNodes() external view returns (uint) {
        return numberOfNodes;
    }

    /// @notice called at the time of network initialization for adding
    /// @notice admin nodes
    /// @param _enodeId enode id
    /// @param _orgId org id to which the enode belongs
    function addAdminNode(string calldata _enodeId, string calldata _orgId) external
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _orgId, 2));
        emit NodeApproved(_enodeId, _orgId);
    }

    /// @notice called at the time of new org creation to add node to org
    /// @param _enodeId enode id
    /// @param _orgId org id to which the enode belongs
    function addNode(string calldata _enodeId, string calldata _orgId) external
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _orgId, 1));
        emit NodeProposed(_enodeId, _orgId);
    }

    /// @notice called org admins to add new enodes to the org or sub orgs
    /// @param _enodeId enode id
    /// @param _orgId org or sub org id to which the enode belongs
    function addOrgNode(string calldata _enodeId, string calldata _orgId) external
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _orgId, 2));
        emit NodeApproved(_enodeId, _orgId);
    }

    /// @notice function to approve the node addition. only called at the time
    /// @notice master org creation by network admin
    /// @param _enodeId enode id
    /// @param _orgId org or sub org id to which the enode belongs
    function approveNode(string calldata _enodeId, string calldata _orgId) external
    onlyImplementation
    enodeExists(_enodeId) {
        // node should belong to the passed org
        require(_checkOrg(_enodeId, _orgId), "enode id does not belong to the passed org id");
        require(_getNodeStatus(_enodeId) == 1, "nothing pending for approval");
        uint nodeIndex = _getNodeIndex(_enodeId);
        nodeList[nodeIndex].status = 2;
        emit NodeApproved(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].orgId);
    }

    /// @notice updates the node status. can be called for deactivating/
    /// @notice blacklisting  and reactivating a deactivated node
    /// @param _enodeId enode id
    /// @param _orgId org or sub org id to which the enode belong
    /// @param _action 1- deactivate, 2- reactivate, 3- blacklist node
    function updateNodeStatus(string calldata _enodeId, string calldata _orgId, uint _action) external
    onlyImplementation
    enodeExists(_enodeId) {
        // node should belong to the org
        require(_checkOrg(_enodeId, _orgId), "enode id does not belong to the passed org");
        require((_action == 1 || _action == 2 || _action == 3),
            "invalid operation. wrong action passed");

        if (_action == 1) {
            require(_getNodeStatus(_enodeId) == 2, "operation cannot be performed");
            nodeList[_getNodeIndex(_enodeId)].status = 3;
            emit NodeDeactivated(_enodeId, _orgId);
        }
        else if (_action == 2) {
            require(_getNodeStatus(_enodeId) == 3, "operation cannot be performed");
            nodeList[_getNodeIndex(_enodeId)].status = 2;
            emit NodeActivated(_enodeId, _orgId);
        }
        else {
            nodeList[_getNodeIndex(_enodeId)].status = 5;
            emit NodeBlacklisted(_enodeId, _orgId);
        }
    }

    /* private functions */
    /// @notice returns the node index for given enode id
    /// @param _enodeId enode id
    /// @return trur or false
    function _getNodeIndex(string memory _enodeId) internal view
    returns (uint) {
        return nodeIdToIndex[keccak256(abi.encode(_enodeId))] - 1;
    }

    /// @notice checks if enode id is linked to the org id passed
    /// @param _enodeId enode id
    /// @param _orgId org or sub org id to which the enode belongs
    /// @return true or false
    function _checkOrg(string memory _enodeId, string memory _orgId) internal view
    returns (bool) {
        return (keccak256(abi.encode(nodeList[_getNodeIndex(_enodeId)].orgId)) == keccak256(abi.encode(_orgId)));
    }

    /// @notice returns the node status for a given enode id
    /// @param _enodeId enode id
    /// @return node status
    function _getNodeStatus(string memory _enodeId) internal view returns (uint) {
        if (nodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0) {
            return 0;
        }
        return nodeList[_getNodeIndex(_enodeId)].status;
    }
}
