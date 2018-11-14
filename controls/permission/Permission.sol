pragma solidity ^0.4.23;

contract Permissions {

  // enum and struct declaration
  enum NodeStatus {NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated, PendingActivation, PendingBlacklisting, Blacklisted }
  struct NodeDetails {
    string enodeId; //e.g. 127.0.0.1:20005
    string ipAddrPort;
    string discPort;
    string raftPort;
    NodeStatus status;
  }

  enum AccountAccess { FullAccess, ReadOnly, Transact, ContractDeploy }
  struct AccountAccessDetails {
    address acctId;
    AccountAccess acctAccess;
  }
  // use an array to store node details
  // if we want to list all node one day, mapping is not capable
  NodeDetails[] private nodeList;

  // use a mapping of enodeid to array index to track node
  mapping (bytes32 => uint) private nodeIdToIndex;
  // keep track of node number
  uint private numberOfNodes;

  AccountAccessDetails[] private acctAccessList;
  mapping (address => uint) private acctToIndex;
  uint private numberOfAccts;
  // use an array to store account details
  // if we want to list all account one day, mapping is not capable
  address[] private voterAcctList;

  // store node approval, deactivation and blacklisting vote status (prevent double vote)
  mapping (uint => mapping (address => bool)) private voteStatus;
  // valid vote count
  mapping (uint => uint) private voteCount;

  // checks if firts time network boot up has happened or not
  bool private networkBoot = false;

  // node permission events for new node propose
  event NodeProposed(string _enodeId);
  event NodeApproved(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeApproval(string _enodeId, address _accountAddress);

  // node permission events for node decativation
  event NodePendingDeactivation (string _enodeId);
  event NodeDeactivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeDeactivation(string _enodeId, address _accountAddress);

  // node permission events for node activation 
  event NodePendingActivation(string _enodeId);
  event NodeActivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);
  event VoteNodeActivation(string _enodeId, address _accountAddress);

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
    for (uint i=0; i<voterAcctList.length; i++){
      if (voterAcctList[i] == msg.sender){
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
  function getNumberOfVoters() public view returns (uint)
  {
	return voterAcctList.length;
  }
  // Get number of nodes
  function getNetworkBootStatus() public view returns (bool)
  {
    return networkBoot;
  }

  // Get node details given index
  function getNodeDetails(uint nodeIndex) public view returns (string _enodeId, string _ipAddrPort, string _discPort, string _raftPort, NodeStatus _nodeStatus)
  {
    return (nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort, nodeList[nodeIndex].status);
  }
  // Get number of nodes
  function getNumberOfNodes() public view returns (uint)
  {
    return numberOfNodes;
  }
  // Get number of accounts and voting accounts
  function getNumberOfAccounts() public view returns (uint)
  {
    return acctAccessList.length;
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
  function isVoter(address _acctid) external view returns (bool)
  {
    bool flag = false;
    for (uint i=0; i<voterAcctList.length; i++){
      if (voterAcctList[i] == _acctid){
        flag = true;
        break;
      }
    }
    return flag;
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
    if (_index <= voterAcctList.length){
      return voterAcctList[_index];
    } else {
      return address(0);
    }
  }

  // state change functions
  // update the networ boot status as true
  function updateNetworkBootStatus() external returns (bool)
  {
    require (networkBoot == false, "Invalid call: Network boot up completed");
    networkBoot = true;
    return networkBoot;
  }

  function initNodeVoteStatus(uint nodeIndex) internal {
    voteCount[nodeIndex] = 0;
    for (uint i = 0; i < voterAcctList.length; i++){
      voteStatus[nodeIndex][voterAcctList[i]] = false;
    }
  }

  function updateVoteStatus(uint nodeIndex) internal {
    voteCount[nodeIndex]++;
    voteStatus[nodeIndex][msg.sender] = true;
  }

  function checkEnoughVotes(uint nodeIndex) internal view returns (bool) {
    bool approvalStatus = false;
    if (voteCount[nodeIndex] > voterAcctList.length / 2){
      approvalStatus = true;
    }
    return approvalStatus;
  }

  // propose a new node to the network
  function proposeNode(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort) external enodeNotInList(_enodeId)
  {
    if (!(networkBoot)){
      numberOfNodes++;
      nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
      nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.Approved));
    }
    else {
      if (checkVotingAccountExist()){
        // increment node number, add node to the list
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.PendingApproval));

        // add voting status, numberOfNodes is the index of current proposed node
        initNodeVoteStatus(numberOfNodes);
        // emit event
        emit NodeProposed(_enodeId);
      }
    }
  }

  // Adds a node to the nodeList mapping and emits node added event if successfully and node exists event of node is already present
  function approveNode(string _enodeId) external canVote
  {
      require(getNodeStatus(_enodeId) == NodeStatus.PendingApproval, "Node need to be in PendingApproval status");
      uint nodeIndex = getNodeIndex(_enodeId);
      require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
      // vote node
      updateVoteStatus(nodeIndex);
      // emit event
      emit VoteNodeApproval(_enodeId, msg.sender);
      // check if node vote reach majority
      if (checkEnoughVotes(nodeIndex)) {
        nodeList[nodeIndex].status = NodeStatus.Approved;
        emit NodeApproved(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
      }
  }

  // Propose a node for deactivation from network
  function proposeDeactivation(string _enodeId) external enodeInList(_enodeId)
  {
    if (checkVotingAccountExist()){
      require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
      uint nodeIndex = getNodeIndex(_enodeId);
      nodeList[nodeIndex].status = NodeStatus.PendingDeactivation;
      // add voting status, numberOfNodes is the index of current proposed node
      initNodeVoteStatus(nodeIndex);
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
    updateVoteStatus(nodeIndex);
    // emit event
    emit VoteNodeDeactivation(_enodeId, msg.sender);
    // check if node vote reach majority
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Deactivated;
      emit NodeDeactivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
      }
  }

  // Propose node for blacklisting
  function proposeNodeActivation(string _enodeId) external
  {
    if (checkVotingAccountExist()){
      require(getNodeStatus(_enodeId) == NodeStatus.Deactivated, "Node need to be in Deactivated status");
      uint nodeIndex = getNodeIndex(_enodeId);
      nodeList[nodeIndex].status = NodeStatus.PendingActivation;
      // add voting status, numberOfNodes is the index of current proposed node
      initNodeVoteStatus(nodeIndex);
      // emit event
      emit NodePendingActivation(_enodeId);
    }
  }

  //deactivates a given Enode and emits the decativation event
  function activateNode(string _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingActivation, "Node need to be in PendingActivation status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    updateVoteStatus(nodeIndex);
    // emit event
    emit VoteNodeDeactivation(_enodeId, msg.sender);
    // check if node vote reach majority
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Approved;
      emit NodeActivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  // Propose node for blacklisting
  function proposeNodeBlacklisting(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort) external
  {
    if (checkVotingAccountExist()){
      uint nodeIndex = getNodeIndex(_enodeId);
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
      initNodeVoteStatus(nodeIndex);
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
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Blacklisted;
      emit NodeBlacklisted(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }


  // Checks if the Node is already added. If yes then returns true
  function updateAccountAccess(address _address, AccountAccess _accountAccess) external
  {
    // Check if account already exists
    emit AccountAccessModified(_address, _accountAccess);
    uint acctIndex = getAcctIndex(_address);
    if (acctToIndex[_address] != 0){
      acctAccessList[acctIndex].acctAccess = _accountAccess;
    }
    else{
      numberOfAccts ++;
      acctToIndex[_address] = numberOfAccts;
      acctAccessList.push(AccountAccessDetails(_address, _accountAccess));
    }
  }

  // Add voting account
  function addVoter(address _address) external
  {
    // Check if account already exists
    for (uint i=0; i<voterAcctList.length; i++){
      if (voterAcctList[i] == _address){
        return;
      }
    }
    voterAcctList.push(_address);
    emit VoterAdded(_address);
  }
  // Remove voting account
  function removeVoter(address _address) external
  {
    // Check if account already exists
    for (uint i=0; i<voterAcctList.length; i++){
      if (voterAcctList[i] == _address){
        for (uint j=i; j<voterAcctList.length -1; j++){
          voterAcctList[j] = voterAcctList[j+1];
        }
        delete voterAcctList[voterAcctList.length - 1];
        emit VoterRemoved(_address);
      }
    }
  }

  /* private functions */
  function getNodeIndex(string _enodeId) internal view returns (uint)
  {
    return nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] - 1;
  }

  /* private functions */
  function getAcctIndex(address _acct) internal view returns (uint)
  {
    return acctToIndex[_acct] - 1;
  }

  function checkVotingAccountExist() internal returns (bool)
  {
    if (voterAcctList.length == 0){
      emit NoVotingAccount();
      return false;
    } else {
      return true;
    }
  }

}
