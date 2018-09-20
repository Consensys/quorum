pragma solidity ^0.4.23;

contract Permissions {

  // enum and struct declaration
  enum NodeStatus { NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated, PendingBlacklisting, Blacklisted }
  enum AccountAccess { FullAccess, ReadOnly, Transact, ContractDeploy }
  struct NodeDetails {
    string enodeId; //e.g. 127.0.0.1:20005
    string ipAddrPort;
    string discPort;
    string raftPort;
    NodeStatus status;
  }

  // use an array to store node details
  // if we want to list all node one day, mapping is not capable
  NodeDetails[] private nodeList;
  // use a mapping of enodeid to array index to track node
  mapping (bytes32 => uint) private nodeIdToIndex;
  // keep track of node number
  uint private numberOfNodes;

  // use an array to store account details
  // if we want to list all account one day, mapping is not capable
  address[] private accountList;

  // store node approval, deactivation and blacklisting vote status (prevent double vote)
  mapping (uint => mapping (address => bool)) private voteStatus;
  // valid vote count
  mapping (uint => uint) private voteCount;

  // node permission events for new node propose
  event NodeProposed(string _enodeId);
  event NodeApproved(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeApproval(string _enodeId, address _accountAddress);

  // node permission events for node decativation
  event NodePendingDeactivation (string _enodeId);
  event NodeDeactivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeDeactivation(string _enodeId, address _accountAddress);

  // node permission events for node blacklist 
  event NodePendingBlacklist(string _enodeId);
  event NodeBlacklisted(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeBlacklist(string _enodeId, address _accountAddress);

  // account permission events
  event AccountAccessModified(address _address, AccountAccess _access);

  // events related to voting accounts for majority voting
  event NoVotingAccount();
  event VoterAdded(address _address);
  event VoterRemoved(address _address);

  // Checks if the given enode exists
  modifier enodeInList(string _enodeId)
  {
    require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0, "Enode is not in the list");
    _;
  }

  // Checks if the given enode does not exists
  modifier enodeNotInList(string _enodeId)
  {
    require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] == 0, "Enode is in the list");
    _;
  }

  // Checks if the account can vote 
  modifier canVote()
  {
    bool flag = false;
    for (uint i=0; i<accountList.length; i++){
      if (accountList[i] == msg.sender){
        flag = true;
        break;
      }
    }
    require(flag, "Account can not vote");
    _;
  }

  /* public and external functions */
  // view functions

  // Get number of nodes
  function getNumberOfNodes() public view returns (uint)
  {
    return numberOfNodes;
  }
  // Get number of accounts and voting accounts
  function getNumberOfAccounts() public view returns (uint)
  {
    return accountList.length;
  }
  // Get node status by enode id
  function getNodeStatus(string _enodeId) public view enodeInList(_enodeId) returns (NodeStatus)
  {
    return nodeList[getNodeIndex(_enodeId)].status;
  }
  // Get vote count by enode id
  function getVoteCount(string _enodeId) public view enodeInList(_enodeId) returns (uint)
  {
    return voteCount[getNodeIndex(_enodeId)];
  }
  // Get vote status by enode id and voter address
  function getVoteStatus(string _enodeId, address _voter) public view enodeInList(_enodeId) returns (bool)
  {
    return voteStatus[getNodeIndex(_enodeId)][_voter];
  }
  // for potential external use
  // Get enode id by index
  function getEnodeId(uint _index) external view returns (string)
  {
    if (_index <= numberOfNodes){
      return nodeList[_index].enodeId;
    } else {
      return "";
    }
  }
  // Get account address by index
  function getAccountAddress(uint _index) external view returns (address)
  {
    if (_index <= accountList.length){
      return accountList[_index];
    } else {
      return address(0);
    }
  }

  // state change functions

  // propose a new node to the network
  function proposeNode(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort) external enodeNotInList(_enodeId)
  {
    if (checkVotingAccountExist()){
      // increment node number, add node to the list
      numberOfNodes++;
      nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
      nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.PendingApproval));
      // add voting status, numberOfNodes is the index of current proposed node
      for (uint i = 0; i < accountList.length; i++){
        voteStatus[numberOfNodes][accountList[i]] = false;
      }
      voteCount[numberOfNodes] = 0;
      // emit event
      emit NodeProposed(_enodeId);
    }
  }

  // Adds a node to the nodeList mapping and emits node added event if successfully and node exists event of node is already present
  function approveNode(string _enodeId) external canVote
  {
      require(getNodeStatus(_enodeId) == NodeStatus.PendingApproval, "Node need to be in PendingApproval status");
      uint nodeIndex = getNodeIndex(_enodeId);
      require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
      // vote node
      voteStatus[nodeIndex][msg.sender] = true;
      voteCount[nodeIndex]++;
      // emit event
      emit VoteNodeApproval(_enodeId, msg.sender);
      // check if node vote reach majority
      checkNodeApproval(_enodeId);
  }

  // Propose a node for deactivation from network
  function proposeDeactivation(string _enodeId) external enodeInList(_enodeId)
  {
    if (checkVotingAccountExist()){
      require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
      uint nodeIndex = getNodeIndex(_enodeId);
      nodeList[nodeIndex].status = NodeStatus.PendingDeactivation;
      // add voting status, numberOfNodes is the index of current proposed node
      for (uint i = 0; i < accountList.length; i++){
        voteStatus[nodeIndex][accountList[i]] = false;
      }
      voteCount[nodeIndex] = 0;
      // emit event
      emit NodePendingDeactivation(_enodeId);
    }
  }

  //deactivates a given Enode and emits the decativation event
  function deactivateNode(string _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation, "Node need to be in PendingDeactivation status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    voteStatus[nodeIndex][msg.sender] = true;
    voteCount[nodeIndex]++;
    // emit event
    emit VoteNodeDeactivation(_enodeId, msg.sender);
    // check if node vote reach majority
    checkNodeDeactivation(_enodeId);
  }

  // Propose node for blacklisting
  function proposeNodeBlacklisting(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort) external
  {
    if (checkVotingAccountExist()){
      uint nodeIndex;
      // check if node is in the nodeList
      if (nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0){
        // no matter what status the node is in, vote will reset and node status change to PendingBlacklisting
        nodeList[nodeIndex].status = NodeStatus.PendingBlacklisting;
        nodeIndex = getNodeIndex(_enodeId);
      } else {
        // increment node number, add node to the list
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.PendingBlacklisting));
        nodeIndex = numberOfNodes;
      }
      // add voting status, numberOfNodes is the index of current proposed node
      for (uint i = 0; i < accountList.length; i++){
        voteStatus[nodeIndex][accountList[i]] = false;
      }
      voteCount[nodeIndex] = 0;
      // emit event
      emit NodePendingBlacklist(_enodeId);
    }
  }

  //Approve node blacklisting
  function blacklistNode(string _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingBlacklisting, "Node need to be in PendingBlacklisting status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    voteStatus[nodeIndex][msg.sender] = true;
    voteCount[nodeIndex]++;
    // emit event
    emit VoteNodeBlacklist(_enodeId, msg.sender);
    // check if node vote reach majority
    checkNodeBlacklisting(_enodeId);
  }

  // Checks if the Node is already added. If yes then returns true
  function updateAccountAccess(address _address, AccountAccess _accountAccess) external
  {
      emit AccountAccessModified(_address, _accountAccess);
  }

  // Add voting account
  function addVoter(address _address) external
  {
    // Check if account already exists
    for (uint i=0; i<accountList.length; i++){
      if (accountList[i] == _address){
        return;
      }
    }
    accountList.push(_address);
    emit VoterAdded(_address);
  }
  // Remove voting account
  function removeVoter(address _address) external
  {
    // Check if account already exists
    for (uint i=0; i<accountList.length; i++){
      if (accountList[i] == _address){
        for (uint j=i+1; j<accountList.length; j++){
          accountList[j-1] = accountList[j];
        }
        delete accountList[accountList.length];
        emit VoterRemoved(_address);
      }
    }
  }

  /* private functions */

  function getNodeIndex(string _enodeId) internal view returns (uint)
  {
    return nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] - 1;
  }

  function checkVotingAccountExist() internal returns (bool)
  {
    if (accountList.length == 0){
      emit NoVotingAccount();
      return false;
    } else {
      return true;
    }
  }

  function checkNodeApproval(string _enodeId) internal
  {
    uint nodeIndex = getNodeIndex(_enodeId);
    if (voteCount[nodeIndex] > accountList.length / 2){
      nodeList[nodeIndex].status = NodeStatus.Approved;
      emit NodeApproved(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  function checkNodeDeactivation(string _enodeId) internal
  {
    uint nodeIndex = getNodeIndex(_enodeId);
    if (voteCount[nodeIndex] > accountList.length / 2){
      nodeList[nodeIndex].status = NodeStatus.Deactivated;
      emit NodeDeactivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  function checkNodeBlacklisting(string _enodeId) internal
  {
    uint nodeIndex = getNodeIndex(_enodeId);
    if (voteCount[nodeIndex] > accountList.length / 2){
      nodeList[nodeIndex].status = NodeStatus.Blacklisted;
      emit NodeBlacklisted(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

}
