pragma solidity ^0.5.3;

import "./PermissionsUpgradable.sol";
/** @title Node manager contract
  * @notice This contract holds implementation logic for all node management
    functionality. This can be called only by the implementation contract.
    There are few view functions exposed as public and can be called directly.
    These are invoked by quorum for populating permissions data in cache
  * @dev node status is denoted by a fixed integer value. The values are
    as below:
        0 - Not in list
        1 - Node pending approval
        2 - Active
        3 - Deactivated
        4 - Blacklisted
        5 - Blacklisted node recovery initiated. Once approved the node
            status will be updated to Active (2)
     Once the node is blacklisted no further activity on the node is
     possible.
  */
contract NodeManager {
    PermissionsUpgradable private permUpgradable;

    struct NodeDetails {
        string enodeId;
        string ip;
        uint16 port;
        uint16 raftPort;
        string orgId;
        uint256 status;
    }
    // use an array to store node details
    // if we want to list all node one day, mapping is not capable
    NodeDetails[] private nodeList;
    // mapping of enode id to array index to track node
    mapping(bytes32 => uint256) private nodeIdToIndex;
    // mapping of enodeId to array index to track node
    mapping(bytes32 => uint256) private enodeIdToIndex;
    // tracking total number of nodes in network
    uint256 private numberOfNodes;


    // node permission events for new node propose
    event NodeProposed(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);
    event NodeApproved(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    // node permission events for node deactivation
    event NodeDeactivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    // node permission events for node activation
    event NodeActivated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    // node permission events for node blacklist
    event NodeBlacklisted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    // node permission events for initiating the recovery of blacklisted
    // node
    event NodeRecoveryInitiated(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    // node permission events for completing the recovery of blacklisted
    // node
    event NodeRecoveryCompleted(string _enodeId, string _ip, uint16 _port, uint16 _raftport, string _orgId);

    /** @notice confirms that the caller is the address of implementation
        contract
    */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /** @notice  checks if the node exists in the network
      * @param _enodeId full enode id
      */
    modifier enodeExists(string memory _enodeId) {
        require(enodeIdToIndex[keccak256(abi.encode(_enodeId))] != 0,
            "passed enode id does not exist");
        _;
    }

    /** @notice  checks if the node does not exist in the network
      * @param _enodeId full enode id
      */
    modifier enodeDoesNotExists(string memory _enodeId) {
        require(enodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0,
            "passed enode id exists");
        _;
    }

    /** @notice constructor. sets the permissions upgradable address
      */
    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice fetches the node details given an enode id
      * @param _enodeId full enode id
      * @return org id
      * @return enode id
      * @return status of the node
      */
    function getNodeDetails(string calldata enodeId) external view
    returns (string memory _orgId, string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus) {
        if (nodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0) {
            return ("", "", "", 0, 0, 0);
        }
        uint256 nodeIndex = _getNodeIndex(enodeId);
        return (nodeList[nodeIndex].orgId, nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ip,
        nodeList[nodeIndex].port, nodeList[nodeIndex].raftPort,
        nodeList[nodeIndex].status);
    }

    /** @notice fetches the node details given the index of the enode
      * @param _nodeIndex node index
      * @return org id
      * @return enode id
      * @return ip of the node
      * @return port of the node
      * @return raftport of the node
      * @return status of the node
      */
    function getNodeDetailsFromIndex(uint256 _nodeIndex) external view
    returns (string memory _orgId, string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, uint256 _nodeStatus) {
        return (nodeList[_nodeIndex].orgId, nodeList[_nodeIndex].enodeId, nodeList[_nodeIndex].ip,
        nodeList[_nodeIndex].port, nodeList[_nodeIndex].raftPort,
        nodeList[_nodeIndex].status);
    }

    /** @notice returns the total number of enodes in the network
      * @return number of nodes
      */
    function getNumberOfNodes() external view returns (uint256) {
        return numberOfNodes;
    }

    /** @notice called at the time of network initialization for adding
        admin nodes
      * @param _enodeId enode id
      * @param _ip IP of node
      * @param _port tcp port of node
      * @param _raftport raft port of node
      * @param _orgId org id to which the enode belongs
      */
    function addAdminNode(string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, string memory _orgId) public
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        enodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ip, _port, _raftport, _orgId, 2));
        emit NodeApproved(_enodeId, _ip, _port, _raftport, _orgId);
    }

    /** @notice called at the time of new org creation to add node to org
      * @param _enodeId enode id
      * @param _ip IP of node
      * @param _port tcp port of node
      * @param _raftport raft port of node
      * @param _orgId org id to which the enode belongs
      */
    function addNode(string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, string memory _orgId) public
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        enodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ip, _port, _raftport, _orgId, 1));
        emit NodeProposed(_enodeId, _ip, _port, _raftport, _orgId);
    }

    /** @notice called org admins to add new enodes to the org or sub orgs
      * @param _enodeId enode id
      * @param _ip IP of node
      * @param _port tcp port of node
      * @param _raftport raft port of node
      * @param _orgId org or sub org id to which the enode belongs
      */
    function addOrgNode(string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, string memory _orgId) public
    onlyImplementation
    enodeDoesNotExists(_enodeId) {
        numberOfNodes++;
        enodeIdToIndex[keccak256(abi.encode(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ip, _port, _raftport, _orgId, 2));
        emit NodeApproved(_enodeId, _ip, _port, _raftport, _orgId);
    }

    /** @notice function to approve the node addition. only called at the time
        master org creation by network admin
      * @param _enodeId enode id
      * @param _ip IP of node
      * @param _port tcp port of node
      * @param _raftport raft port of node
      * @param _orgId org or sub org id to which the enode belongs
      */
    function approveNode(string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, string memory _orgId) public
    onlyImplementation
    enodeExists(_enodeId) {
        // node should belong to the passed org
        require(_checkOrg(_enodeId, _orgId), "enode id does not belong to the passed org id");
        require(_getNodeStatus(_enodeId) == 1, "nothing pending for approval");
        uint256 nodeIndex = _getNodeIndex(_enodeId);
        if (keccak256(abi.encode(nodeList[nodeIndex].ip)) != keccak256(abi.encode(_ip)) || nodeList[nodeIndex].port != _port || nodeList[nodeIndex].raftPort != _raftport) {
            return;
        }
        nodeList[nodeIndex].status = 2;
        emit NodeApproved(nodeList[nodeIndex].enodeId, _ip, _port, _raftport, nodeList[nodeIndex].orgId);
    }

    /** @notice updates the node status. can be called for deactivating/
        blacklisting  and reactivating a deactivated node
      * @param _enodeId enode id
      * @param _ip IP of node
      * @param _port tcp port of node
      * @param _raftport raft port of node
      * @param _orgId org or sub org id to which the enode belong
      * @param _action action being performed
      * @dev action can have any of the following values
            1 - Suspend the node
            2 - Revoke suspension of a suspended node
            3 - blacklist a node
            4 - initiate the recovery of a blacklisted node
            5 - blacklisted node recovery fully approved. mark to active
      */
    function updateNodeStatus(string memory _enodeId, string memory _ip, uint16 _port, uint16 _raftport, string memory _orgId, uint256 _action) public
    onlyImplementation
    enodeExists(_enodeId) {
        // node should belong to the org
        require(_checkOrg(_enodeId, _orgId), "enode id does not belong to the passed org");
        require((_action == 1 || _action == 2 || _action == 3 || _action == 4 || _action == 5),
            "invalid operation. wrong action passed");

        uint256 nodeIndex = _getNodeIndex(_enodeId);
        if (keccak256(abi.encode(nodeList[nodeIndex].ip)) != keccak256(abi.encode(_ip)) || nodeList[nodeIndex].port != _port || nodeList[nodeIndex].raftPort != _raftport) {
            return;
        }

        if (_action == 1) {
            require(_getNodeStatus(_enodeId) == 2, "operation cannot be performed");
            nodeList[nodeIndex].status = 3;
            emit NodeDeactivated(_enodeId, _ip, _port, _raftport, _orgId);
        }
        else if (_action == 2) {
            require(_getNodeStatus(_enodeId) == 3, "operation cannot be performed");
            nodeList[nodeIndex].status = 2;
            emit NodeActivated(_enodeId, _ip, _port, _raftport, _orgId);
        }
        else if (_action == 3) {
            nodeList[nodeIndex].status = 4;
            emit NodeBlacklisted(_enodeId, _ip, _port, _raftport, _orgId);
        } else if (_action == 4) {
            // node should be in blacklisted state
            require(_getNodeStatus(_enodeId) == 4, "operation cannot be performed");
            nodeList[nodeIndex].status = 5;
            emit NodeRecoveryInitiated(_enodeId, _ip, _port, _raftport, _orgId);
        } else {
            // node should be in initiated recovery state
            require(_getNodeStatus(_enodeId) == 5, "operation cannot be performed");
            nodeList[nodeIndex].status = 2;
            emit NodeRecoveryCompleted(_enodeId, _ip, _port, _raftport, _orgId);
        }
    }

    // private functions
    /** @notice returns the node index for given enode id
      * @param _enodeId enode id
      * @return trur or false
      */
    function _getNodeIndex(string memory _enodeId) internal view
    returns (uint256) {
        return enodeIdToIndex[keccak256(abi.encode(_enodeId))] - 1;
    }

    /** @notice checks if enode id is linked to the org id passed
      * @param _enodeId enode id
      * @param _orgId org or sub org id to which the enode belongs
      * @return true or false
      */
    function _checkOrg(string memory _enodeId, string memory _orgId) internal view
    returns (bool) {
        return (keccak256(abi.encode(nodeList[_getNodeIndex(_enodeId)].orgId)) == keccak256(abi.encode(_orgId)));
    }

    /** @notice returns the node status for a given enode id
      * @param _enodeId enode id
      * @return node status
      */
    function _getNodeStatus(string memory _enodeId) internal view returns (uint256) {
        if (enodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0) {
            return 0;
        }
        return nodeList[_getNodeIndex(_enodeId)].status;
    }

    /** @notice checks if the node is allowed to connect or not
    * @param _enodeId enode id
    * @param _ip IP of node
    * @param _port tcp port of node
    * @return bool indicating if the node is allowed to connect or not
    */
    function connectionAllowed(string memory _enodeId, string memory _ip, uint16 _port) public view onlyImplementation
    returns (bool){
        if (enodeIdToIndex[keccak256(abi.encode(_enodeId))] == 0) {
            return false;
        }
        uint256 nodeIndex = _getNodeIndex(_enodeId);
        if (nodeList[nodeIndex].status == 2 && keccak256(abi.encode(nodeList[nodeIndex].ip)) == keccak256(abi.encode(_ip))) {
            return true;
        }

        return false;
    }
}
