pragma solidity ^0.5.3;

contract Permissions {
  address[] initialAcctList;
  // enum and struct declaration
  enum NodeStatus {NotInList, PendingApproval, Approved, PendingDeactivation, Deactivated, PendingActivation, PendingBlacklisting, Blacklisted }
  struct NodeDetails {
    string enodeId; //e.g. 127.0.0.1:20005
    string ipAddrPort;
    string discPort;
    string raftPort;
    NodeStatus status;
  }
  // use an array to store node details
  NodeDetails[] private nodeList;
  // use a mapping of enodeid to array index to track node
  mapping (bytes32 => uint) private nodeIdToIndex;
  // keep track of node number
  uint private numberOfNodes;

  enum AccountAccess { ReadOnly, Transact, ContractDeploy, FullAccess}
  struct AccountAccessDetails {
    address acctId;
    AccountAccess acctAccess;
  }
  AccountAccessDetails[] private acctAccessList;
  mapping (address => uint) private acctToIndex;
  uint private numberOfAccts;
  uint private numFullAccessAccts;

  // use an array to store account details
  enum VoterStatus { Active, Inactive }
  struct VoterAcctDetails {
    address voterAcct;
    VoterStatus voterStatus;
  }
  VoterAcctDetails[] private voterAcctList;
  mapping (address => uint) private voterAcctIndex;
  uint private numberOfVoters;
  uint private numberOfValidVoters;

  // store pre pending status, use for cancelPendingOperation
  mapping(uint => NodeStatus) private prependingStatus;
  // store node approval, deactivation and blacklisting vote status (prevent double vote)
  mapping (uint => mapping (address => bool)) private voteStatus;
  // valid vote count
  mapping (uint => uint) private voteCount;

  // checks if first time network boot up has happened or not
  bool private networkBoot = false;

  // node permission events for new node propose
  event NodeProposed(string _enodeId);
  event NodeApproved(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);

  // node permission events for node decativation
  event NodePendingDeactivation (string _enodeId);
  event NodeDeactivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);

  // node permission events for node activation
  event NodePendingActivation(string _enodeId);
  event NodeActivated(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);

  // node permission events for node blacklist
  event NodePendingBlacklist(string _enodeId);
  event NodeBlacklisted(string _enodeId, string _ipAddrPort, string _discPort, string _raftPort);

  // pending operation cancelled
  event PendingOperationCancelled(string _enodeId);

  // account permission events
  event AccountAccessModified(address _address, AccountAccess _access);

  // Checks if the given enode exists
  modifier enodeInList(string memory _enodeId)
  {
    require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0, "Enode is not in the list");
    _;
  }

  // Checks if the given enode does not exists
  modifier enodeNotInList(string memory _enodeId)
  {
    require(nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] == 0 || getNodeStatus(_enodeId) == NodeStatus.NotInList, "Enode is in the list");
    _;
  }

  // Checks if the account can vote
  modifier canVote()
  {
    bool flag = false;
    uint voterIndex = getVoterIndex(msg.sender);
    if ((voterAcctIndex[msg.sender] != 0) && (voterAcctList[voterIndex].voterStatus == VoterStatus.Active)) {
      flag = true;
    }
    require(flag, "Account can not vote");
    _;
  }

  /* public and external functions */
  // view functions

  // get number of accounts in the init list given as per genesis.json
  function getInitAccountsCount() external view returns (uint){
    return initialAcctList.length;
  }

  // returns the numbers of accounts which will have full access 
  function getFullAccessAccountCount() external view returns (uint){
    return numFullAccessAccts;
  }
  // Get number of voters
  function getNumberOfVoters() external view returns (uint)
  {
    return numberOfVoters;
  }

  // Get number of valid voters
  function getNumberOfValidVoters() external view returns (uint)
  {
    return numberOfValidVoters;
  }
  // Get voter details given the voter index
  function getVoter(uint i) external view returns (address _addr, VoterStatus _voterStatus)
  {
  	return (voterAcctList[i].voterAcct, voterAcctList[i].voterStatus);
  }

  // Get network boot status
  function getNetworkBootStatus() external view returns (bool)
  {
    return networkBoot;
  }

  // Get node details given enode Id
  function getNodeDetails(string calldata enodeId) external view returns (string memory _enodeId, string memory _ipAddrPort, string memory _discPort, string memory _raftPort, NodeStatus _nodeStatus)
  {
    uint nodeIndex = getNodeIndex(enodeId);
    if (nodeIdToIndex[keccak256(abi.encodePacked(enodeId))] != 0){
      return (nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort, nodeList[nodeIndex].status);
    }
    else {
      return (enodeId, "", "", "", NodeStatus.NotInList);
    }
  }

  // Get node details given index
  function getNodeDetailsFromIndex(uint nodeIndex) external view returns (string memory _enodeId, string memory _ipAddrPort, string memory _discPort, string memory _raftPort, NodeStatus _nodeStatus)
  {
    return (nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort, nodeList[nodeIndex].status);
  }

  // Get number of nodes
  function getNumberOfNodes() external view returns (uint)
  {
    return numberOfNodes;
  }

  // Get account details given index
  function getAccountDetails(uint acctIndex) external view returns (address _acct, AccountAccess _acctAccess)
  {
    return (acctAccessList[acctIndex].acctId, acctAccessList[acctIndex].acctAccess);
  }

  // Get number of accounts 
  function getNumberOfAccounts() external view returns (uint)
  {
    return acctAccessList.length;
  }

  // Get node status by enode id
  function getNodeStatus(string memory _enodeId) public view enodeInList(_enodeId) returns (NodeStatus)
  {
    return nodeList[getNodeIndex(_enodeId)].status;
  }

  // checks if the given account is a voter account
  function isVoter(address _acctid) external view returns (bool)
  {
    return ((voterAcctIndex[_acctid] != 0) &&
            (voterAcctList[getVoterIndex(_acctid)].voterStatus == VoterStatus.Active));
  }

  // update the networ boot status as true
  function updateNetworkBootStatus() external returns (bool)
  {
    require (networkBoot == false, "Invalid call: Network boot up completed");
    networkBoot = true;
    return networkBoot;
  }

  // initializes the voting status for each voting account to false
  function initNodeVoteStatus(uint nodeIndex) internal {
    voteCount[nodeIndex] = 0;
    for (uint i = 0; i < voterAcctList.length; i++){
      if (voterAcctList[i].voterStatus == VoterStatus.Active){
        voteStatus[nodeIndex][voterAcctList[i].voterAcct] = false;
      }
    }
  }

  // updates the vote status and increses the vote count
  function updateVoteStatus(uint nodeIndex) internal {
    voteCount[nodeIndex]++;
    voteStatus[nodeIndex][msg.sender] = true;
  }

  // checks if enough votes are received for the approval
  function checkEnoughVotes(uint nodeIndex) internal view returns (bool) {
    bool approvalStatus = false;
    if (voteCount[nodeIndex] > numberOfValidVoters/2){
      approvalStatus = true;
    }
    return approvalStatus;
  }

  // propose a new node to the network
  function proposeNode(string calldata _enodeId, string calldata _ipAddrPort, string calldata _discPort, string calldata _raftPort) external enodeNotInList(_enodeId)
  {
    if (!(networkBoot)){
      numberOfNodes++;
      nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
      nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.Approved));
    }
    else {
      if (checkVotingAccountExist()){
        if (nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0){
          nodeList[getNodeIndex(_enodeId)].status = NodeStatus.PendingApproval;
          prependingStatus[getNodeIndex(_enodeId)] = NodeStatus.NotInList;
        } else {
          // increment node number, add node to the list
          numberOfNodes++;
          nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
          nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.PendingApproval));
          prependingStatus[numberOfNodes] = NodeStatus.NotInList;
        }

        // add voting status, numberOfNodes is the index of current proposed node
        initNodeVoteStatus(numberOfNodes);
        // emit event
        emit NodeProposed(_enodeId);
      }
    }
  }

  // Adds a node to the nodeList mapping and emits node approved event if successful 
  function approveNode(string calldata _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingApproval, "Node need to be in PendingApproval status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    updateVoteStatus(nodeIndex);
    // emit event
    // check if node vote reach majority
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Approved;
      emit NodeApproved(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  // Propose a node for deactivation from network
  function proposeDeactivation(string calldata _enodeId) external enodeInList(_enodeId)
  {
    if (checkVotingAccountExist()){
      require(getNodeStatus(_enodeId) == NodeStatus.Approved, "Node need to be in Approved status");
      uint nodeIndex = getNodeIndex(_enodeId);
      prependingStatus[nodeIndex] = NodeStatus.Approved;
      nodeList[nodeIndex].status = NodeStatus.PendingDeactivation;
      // add voting status, numberOfNodes is the index of current proposed node
      initNodeVoteStatus(nodeIndex);
      // emit event
      emit NodePendingDeactivation(_enodeId);
    }
  }

  //deactivates a given Enode and emits the node decativation event
  function deactivateNode(string calldata _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation, "Node need to be in PendingDeactivation status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    updateVoteStatus(nodeIndex);
    // check if node vote reachead majority and emit event
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Deactivated;
      emit NodeDeactivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  // Propose activation of a deactivated node
  function proposeNodeActivation(string calldata _enodeId) external
  {
    if (checkVotingAccountExist()){
      require(getNodeStatus(_enodeId) == NodeStatus.Deactivated, "Node need to be in Deactivated status");
      uint nodeIndex = getNodeIndex(_enodeId);
      prependingStatus[nodeIndex] = NodeStatus.Deactivated;
      nodeList[nodeIndex].status = NodeStatus.PendingActivation;
      // add voting status, numberOfNodes is the index of current proposed node
      initNodeVoteStatus(nodeIndex);
      // emit event
      emit NodePendingActivation(_enodeId);
    }
  }

  // Activates a given Enode and emits the node activated event
  function activateNode(string calldata _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingActivation, "Node need to be in PendingActivation status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    updateVoteStatus(nodeIndex);
    // check if node vote reachead majority and emit event
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Approved;
      emit NodeActivated(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  // Propose node for blacklisting
  function proposeNodeBlacklisting(string calldata _enodeId, string calldata _ipAddrPort, string calldata _discPort, string calldata _raftPort) external
  {
    if (checkVotingAccountExist()){
      uint nodeIndex = getNodeIndex(_enodeId);
      // check if node is in the nodeList
      if (nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] != 0){
        // no matter what status the node is in, vote will reset and node status change to PendingBlacklisting
        if (!(nodeList[nodeIndex].status == NodeStatus.PendingApproval || nodeList[nodeIndex].status == NodeStatus.PendingActivation || nodeList[nodeIndex].status == NodeStatus.PendingDeactivation || nodeList[nodeIndex].status == NodeStatus.PendingBlacklisting)){
          prependingStatus[nodeIndex] = nodeList[nodeIndex].status;
        }
        nodeList[nodeIndex].status = NodeStatus.PendingBlacklisting;
        nodeIndex = getNodeIndex(_enodeId);
      } else {
        // increment node number, add node to the list
        numberOfNodes++;
        nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] = numberOfNodes;
        nodeList.push(NodeDetails(_enodeId, _ipAddrPort,_discPort, _raftPort, NodeStatus.PendingBlacklisting));
        prependingStatus[nodeIndex] = NodeStatus.NotInList;
        nodeIndex = numberOfNodes;
      }
      initNodeVoteStatus(nodeIndex);
      // emit event
      emit NodePendingBlacklist(_enodeId);
    }
  }

  //Approve node blacklisting
  function blacklistNode(string calldata _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingBlacklisting, "Node need to be in PendingBlacklisting status");
    uint nodeIndex = getNodeIndex(_enodeId);
    require(voteStatus[nodeIndex][msg.sender] == false, "Node can not double vote");
    // vote node
    voteStatus[nodeIndex][msg.sender] = true;
    voteCount[nodeIndex]++;
    // emit event
    // check if node vote reach majority
    if (checkEnoughVotes(nodeIndex)) {
      nodeList[nodeIndex].status = NodeStatus.Blacklisted;
      emit NodeBlacklisted(nodeList[nodeIndex].enodeId, nodeList[nodeIndex].ipAddrPort, nodeList[nodeIndex].discPort, nodeList[nodeIndex].raftPort);
    }
  }

  // Cancel current pending node operation
  function cancelPendingOperation(string calldata _enodeId) external canVote
  {
    require(getNodeStatus(_enodeId) == NodeStatus.PendingApproval ||
            getNodeStatus(_enodeId) == NodeStatus.PendingActivation ||
            getNodeStatus(_enodeId) == NodeStatus.PendingDeactivation ||
            getNodeStatus(_enodeId) == NodeStatus.PendingBlacklisting,
            "Node status must be in pending");

    uint nodeIndex = getNodeIndex(_enodeId);
    nodeList[nodeIndex].status = prependingStatus[nodeIndex];
    emit PendingOperationCancelled(_enodeId);
  }

  // sets the account access to full access for the initial list of accounts
  // given as a part of genesis.json
  function initAccounts() external
  {
    require(networkBoot == false, "network accounts already boot up");
    for (uint i=0; i<initialAcctList.length; i++){
      if (acctToIndex[initialAcctList[i]] == 0){
        numberOfAccts ++;
        numFullAccessAccts ++;
        acctToIndex[initialAcctList[i]] = numberOfAccts;
        acctAccessList.push(AccountAccessDetails(initialAcctList[i], AccountAccess.FullAccess));
        emit AccountAccessModified(initialAcctList[i], AccountAccess.FullAccess);
      }
    }
  }

  // updates accounts access
  function updateAccountAccess(address _address, AccountAccess _accountAccess) external
  {
    // Check if account already exists
    uint acctIndex = getAcctIndex(_address);
    if (acctToIndex[_address] != 0){
      if (acctAccessList[acctIndex].acctAccess == AccountAccess.FullAccess &&
          _accountAccess != AccountAccess.FullAccess &&
          numFullAccessAccts > 1){
        numFullAccessAccts --;
        acctAccessList[acctIndex].acctAccess = _accountAccess;
      }
    }
    else{
      numberOfAccts ++;
      acctToIndex[_address] = numberOfAccts;
      if (_accountAccess == AccountAccess.FullAccess) {
        numFullAccessAccts ++;
      }
      acctAccessList.push(AccountAccessDetails(_address, _accountAccess));
    }
    emit AccountAccessModified(_address, _accountAccess);
  }

  // Add voting account to the network
  function addVoter(address _address) external
  {
    uint vId = getVoterIndex(_address);
    if (voterAcctIndex[_address] != 0) {
      if (voterAcctList[vId].voterStatus == VoterStatus.Inactive){
        voterAcctList[vId].voterStatus = VoterStatus.Active;
        numberOfValidVoters ++;
      }
    }
    else {
      numberOfVoters ++;
      voterAcctIndex[_address] = numberOfVoters;
      voterAcctList.push(VoterAcctDetails(_address, VoterStatus.Active));
      numberOfValidVoters ++;
    }
  }

  // Remove voting account from the network
  function removeVoter(address _address) external
  {
    uint vId = getVoterIndex(_address);
    if (voterAcctIndex[_address] != 0) {
      voterAcctList[vId].voterStatus = VoterStatus.Inactive;
      numberOfValidVoters --;
    }
  }

  // returns total voter count and number of valid voter count
  function getVoterCount() public view returns (uint, uint)
  {
    return (numberOfVoters,numberOfValidVoters);
  }

  /* private functions */

  // Returns the node index based on enode id
  function getNodeIndex(string memory _enodeId) internal view returns (uint)
  {
    return nodeIdToIndex[keccak256(abi.encodePacked(_enodeId))] - 1;
  }

  // Returns the account index based on account id
  function getAcctIndex(address _acct) internal view returns (uint)
  {
    return acctToIndex[_acct] - 1;
  }

  // Returns the voter index based on account id
  function getVoterIndex(address _acct) internal view returns (uint)
  {
    return voterAcctIndex[_acct] - 1;
  }

  // checks if voting account exists
  function checkVotingAccountExist() internal view returns (bool)
  {
    return (!(numberOfValidVoters == 0));
  }

}
