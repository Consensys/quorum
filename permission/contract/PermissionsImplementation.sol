pragma solidity ^0.5.3;

import "./RoleManager.sol";
import "./AccountManager.sol";
import "./VoterManager.sol";
import "./NodeManager.sol";
import "./OrgManager.sol";
import "./PermissionsUpgradable.sol";

/// @title Permissions Implementation Contract
/// @notice This contract holds implementation logic for all permissions
/// @notice related functionality. This can be called only by the interface
/// @notice contract.
contract PermissionsImplementation {
    AccountManager private accounts;
    RoleManager private roles;
    VoterManager private voter;
    NodeManager private nodes;
    OrgManager private org;
    PermissionsUpgradable private permUpgradable;

    string private adminOrg;
    string private adminRole;
    string private orgAdminRole;


    uint private fullAccess = 3;

    /// @dev this variable is meant for tracking the initial network boot up
    /// @dev once the network boot up is done the value is set to true
    bool private networkBoot = false;

    // Modifiers
    /// @notice modifier to confirm that caller is the interface contract
    modifier onlyInterface{
        require(msg.sender == permUpgradable.getPermInterface(),
            "can be called by interface contract only");
        _;
    }
    /// @notice modifier to confirm that caller is the upgradable contract
    modifier onlyUpgradeable {
        require(msg.sender == address(permUpgradable));
        _;
    }

    /// @notice confirms if the network boot status is equal to passed value
    /// @param _status true/false
    modifier networkBootStatus(bool _status){
        require(networkBoot == _status, "Incorrect network boot status");
        _;
    }

    /// @notice confirms that the account passed is network admin account
    /// @param _account account id
    modifier networkAdmin(address _account) {
        require(isNetworkAdmin(_account) == true, "account is not a network admin account");
        _;
    }

    /// @notice confirms that the account passed is org admin account
    /// @param _account account id
    /// @param _orgId org id to which the account belongs
    modifier orgAdmin(address _account, string memory _orgId) {
        require(isOrgAdmin(_account, _orgId) == true, "account is not a org admin account");
        _;
    }

    /// @notice confirms that org does not exist
    /// @param _orgId org id
    modifier orgNotExists(string memory _orgId) {
        require(_checkOrgExists(_orgId) != true, "org exists");
        _;
    }

    /// @notice confirms that org exists
    /// @param _orgId org id
    modifier orgExists(string memory _orgId) {
        require(_checkOrgExists(_orgId) == true, "org does not exist");
        _;
    }

    /// @notice checks of the passed org id is in approved status
    /// @param _orgId org id
    modifier orgApproved(string memory _orgId) {
        require(checkOrgApproved(_orgId) == true, "org not in approved status");
        _;
    }

    /// @notice constructor accepts the contracts addresses of other deployed
    /// @notice contracts of the permissions model
    /// @param _permUpgradable - address of permissions upgradable contract
    /// @param _orgManager - address of org manager contract
    /// @param _rolesManager - address of role manager contract
    /// @param _acctManager - address of account manager contract
    /// @param _voterManager - address of voter manager contract
    /// @param _nodeManager - address of node manager contract
    constructor (address _permUpgradable, address _orgManager, address _rolesManager,
        address _acctManager, address _voterManager, address _nodeManager) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
        org = OrgManager(_orgManager);
        roles = RoleManager(_rolesManager);
        accounts = AccountManager(_acctManager);
        voter = VoterManager(_voterManager);
        nodes = NodeManager(_nodeManager);
    }

    // initial set up related functions
    /// @notice for permissions its necessary to define the initial admin org
    /// @notice id, network admin role id and default org admin role id. this
    /// @notice sets these values at the time of network boot up
    /// @param _nwAdminOrg - address of permissions upgradable contract
    /// @param _nwAdminRole - address of org manager contract
    /// @param _oAdminRole - address of role manager contract
    /// @dev this function will be executed only once as part of the boot up
    function setPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole,
        string calldata _oAdminRole) external onlyInterface
    networkBootStatus(false) {
        adminOrg = _nwAdminOrg;
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
    }

    /// @notice when migrating implementation contract, the values of these
    /// @notice key values need to be set from the previous implementation
    /// @notice contract. this function allows these values to be set
    /// @param _nwAdminOrg - address of permissions upgradable contract
    /// @param _nwAdminRole - address of org manager contract
    /// @param _oAdminRole - address of role manager contract
    /// @param _networkBootStatus - network boot status true/false
    function setMigrationPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole,
        string calldata _oAdminRole, bool _networkBootStatus) external onlyUpgradeable
    networkBootStatus(false) {
        adminOrg = _nwAdminOrg;
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
        networkBoot = _networkBootStatus;
    }

    /// @notice called at the time of network initialization. sets up
    /// @notice network admin org with allowed sub org depth and breadth
    /// @notice creates the network admin for the network admin org
    /// @notice sets the default values required by account manager contract
    /// @param _breadth - number of sub orgs allowed at parent level
    /// @param _depth - levels of sub org nesting allowed at parent level
    function init(uint _breadth, uint _depth) external
    onlyInterface
    networkBootStatus(false) {
        org.setUpOrg(adminOrg, _breadth, _depth);
        roles.addRole(adminRole, adminOrg, fullAccess, true, true);
        accounts.setDefaults(adminRole, orgAdminRole);
    }
    /// @notice as a part of network initialization add all nodes which
    /// @notice are part of static-nodes.json as nodes belonging to
    /// @notice network admin org
    /// @param _enodeId - full enode id
    function addAdminNode(string calldata _enodeId) external
    onlyInterface
    networkBootStatus(false) {
        nodes.addAdminNode(_enodeId, adminOrg);
    }

    /// @notice as a part of network initialization add all accounts which are
    /// @notice passed via permission-config.json as network administrator
    /// @notice accounts
    /// @param _account - account id
    function addAdminAccount(address _account) external
    onlyInterface
    networkBootStatus(false) {
        updateVoterList(adminOrg, _account, true);
        accounts.assignAdminRole(_account, adminOrg, adminRole, 2);
    }

    /// @notice once the network initialization is complete, sets the network
    /// @notice boot status to true
    /// @return network boot status
    /// @dev this will be called only once from geth as a part of
    /// @dev network initialization
    function updateNetworkBootStatus() external
    onlyInterface
    networkBootStatus(false)
    returns (bool){
        networkBoot = true;
        return networkBoot;
    }

    /// @notice function to add a new organization to the network. creates org
    /// @notice record and marks it as pending approval. adds the passed node
    /// @notice node manager contract. adds the account with org admin role to
    /// @notice account manager contracts. creates voting record for approval
    /// @notice by other network admin accounts
    /// @param _orgId unique organization id
    /// @param _enodeId full enode id linked to the organization
    /// @param _account account id. this will have the org admin privileges
    function addOrg(string calldata _orgId, string calldata _enodeId,
        address _account, address _caller) external
    onlyInterface
    networkBootStatus(true)
    networkAdmin(_caller) {
        voter.addVotingItem(adminOrg, _orgId, _enodeId, _account, 1);
        org.addOrg(_orgId);
        nodes.addNode(_enodeId, _orgId);
        require(validateAccount(_account, _orgId) == true,
            "Operation cannot be performed");
        accounts.assignAdminRole(_account, _orgId, orgAdminRole, 1);
    }

    /// @notice functions to approve a pending approval org record by networ
    /// @notice admin account. once majority votes are received the org is
    /// @notice marked as approved
    /// @param _orgId unique organization id
    /// @param _enodeId full enode id linked to the organization
    /// @param _account account id this will have the org admin privileges
    function approveOrg(string calldata _orgId, string calldata _enodeId,
        address _account, address _caller) external onlyInterface networkAdmin(_caller) {
        require(_checkOrgStatus(_orgId, 1) == true, "Nothing to approve");
        if ((processVote(adminOrg, _caller, 1))) {
            org.approveOrg(_orgId);
            roles.addRole(orgAdminRole, _orgId, fullAccess, true, true);
            nodes.approveNode(_enodeId, _orgId);
            accounts.addNewAdmin(_orgId, _account);
        }
    }

    /// @notice function to create a sub org under a given parent org.
    /// @param _pOrgId parent org id under which the sub org is being added
    /// @param _orgId unique id for the sub organization
    /// @param _enodeId full enode id linked to the sjb organization
    /// @dev _enodeId is optional. parent org id should contain the complete
    /// @dev org hierarchy from master org id to the immediate parent. The org
    /// @dev hierarchy is separated by .. For example, if master org ABC has a
    /// @dev sub organization SUB1, then while creating the sub organization at
    /// @dev SUB1 level, the parent org should be given as ABC.SUB1
    function addSubOrg(string calldata _pOrgId, string calldata _orgId,
        string calldata _enodeId, address _caller) external onlyInterface
    orgExists(_pOrgId) orgAdmin(_caller, _pOrgId) {
        org.addSubOrg(_pOrgId, _orgId);
        string memory pOrgId = string(abi.encode(_pOrgId, ".", _orgId));
        if (bytes(_enodeId).length > 0) {
            nodes.addOrgNode(_enodeId, pOrgId);
        }
    }

    /// @notice function to update the org status. it updates the org status
    /// @notice and adds a voting item for network admins to approve
    /// @param _orgId unique id of the organization
    /// @param _action 1 for suspending an org and 2 for revoke of suspension
    function updateOrgStatus(string calldata _orgId, uint _action, address _caller)
    external onlyInterface networkAdmin(_caller) {
        uint pendingOp;
        pendingOp = org.updateOrg(_orgId, _action);
        voter.addVotingItem(adminOrg, _orgId, "", address(0), pendingOp);
    }

    /// @notice function to approve org status change. the org status is
    /// @notice changed once the majority votes are received from network
    /// @notice admin accounts.
    /// @param _orgId unique id for the sub organization
    /// @param _action 1 for suspending an org and 2 for revoke of suspension
    function approveOrgStatus(string calldata _orgId, uint _action, address _caller)
    external onlyInterface networkAdmin(_caller) {
        require((_action == 1 || _action == 2), "Operation not allowed");
        uint pendingOp;
        uint orgStatus;
        if (_action == 1) {
            pendingOp = 2;
            orgStatus = 3;
        }
        else if (_action == 2) {
            pendingOp = 3;
            orgStatus = 5;
        }
        require(_checkOrgStatus(_orgId, orgStatus) == true, "operation not allowed");
        if ((processVote(adminOrg, _caller, pendingOp))) {
            org.approveOrgStatusUpdate(_orgId, _action);
        }
    }

    // Role related functions

    /// @notice function to add new role definition to an organization
    /// @notice can be executed by the org admin account only
    /// @param _roleId unique id for the role
    /// @param _orgId unique id of the organization to which the role belongs
    /// @param _access 0-ReadOnly, 1-Transact, 2-ContractDeploy, 3-FullAccess
    /// @param _voter bool indicates if the role is voter role or not
    /// @param _admin bool indicates if the role is an admin role
    function addNewRole(string calldata _roleId, string calldata _orgId,
        uint _access, bool _voter, bool _admin, address _caller) external
    onlyInterface orgApproved(_orgId) orgAdmin(_caller, _orgId) {
        //add new roles can be created by org admins only
        roles.addRole(_roleId, _orgId, _access, _voter, _admin);
    }

    /// @notice function to remove a role definition from an organization
    /// @notice can be executed by the org admin account only
    /// @param _roleId unique id for the role
    /// @param _orgId unique id of the organization to which the role belongs
    function removeRole(string calldata _roleId, string calldata _orgId,
        address _caller) external onlyInterface orgApproved(_orgId)
    orgAdmin(_caller, _orgId) {
        require(((keccak256(abi.encode(_roleId)) != keccak256(abi.encode(adminRole))) &&
        (keccak256(abi.encode(_roleId)) != keccak256(abi.encode(orgAdminRole)))),
            "admin roles cannot be removed");
        roles.removeRole(_roleId, _orgId);
    }

    // Account related functions
    /// @notice function to assign network admin/org admin role to an account
    /// @notice this can be executed by network admin accounts only. it assigns
    /// @notice the role to the accounts and creates voting record for network
    /// @notice admin accounts
    /// @param _orgId unique id of the organization to which the account belongs
    /// @param _account account id
    /// @param _roleId role id to be assigned to the account
    function assignAdminRole(string calldata _orgId, address _account,
        string calldata _roleId, address _caller) external
    onlyInterface orgExists(_orgId) networkAdmin(_caller) {
        accounts.assignAdminRole(_account, _orgId, _roleId, 1);
        //add voting item
        voter.addVotingItem(adminOrg, _orgId, "", _account, 4);
    }

    /// @notice function to approve network admin/org admin role assigment
    /// @notice this can be executed by network admin accounts only.
    /// @param _orgId unique id of the organization to which the account belongs
    /// @param _account account id
    function approveAdminRole(string calldata _orgId, address _account,
        address _caller) external onlyInterface networkAdmin(_caller) {
        if ((processVote(adminOrg, _caller, 4))) {
            (bool ret, address acct) = accounts.removeExistingAdmin(_orgId);
            if (ret) {
                updateVoterList(adminOrg, acct, false);
            }
            bool ret1 = accounts.addNewAdmin(_orgId, _account);
            if (ret1) {
                updateVoterList(adminOrg, _account, true);
            }
        }
    }

    /// @notice function to update account status. can be executed by org admin
    /// @notice account only.
    /// @param _orgId unique id of the organization to which the account belongs
    /// @param _account account id
    /// @param _status 1-suspending 2-activating back 3-blacklisting
    function updateAccountStatus(string calldata _orgId, address _account,
        uint _status, address _caller) external onlyInterface
    orgAdmin(_caller, _orgId) {
        accounts.updateAccountStatus(_orgId, _account, _status);
    }

    // Node related functions

    /// @notice function to add a new node to the organization. can be invoked
    /// @notice org admin account only
    /// @param _orgId unique id of the organization to which the account belongs
    /// @param _enodeId full enode id being dded to the org
    function addNode(string calldata _orgId, string calldata _enodeId, address _caller)
    external onlyInterface orgApproved(_orgId) orgAdmin(_caller, _orgId) {
        // check that the node is not part of another org
        nodes.addOrgNode(_enodeId, _orgId);
    }

    /// @notice function to update node status. can be invoked by org admin
    /// @notice account only
    /// @param _orgId unique id of the organization to which the account belongs
    /// @param _enodeId full enode id being dded to the org
    /// @param _action 1-deactivate, 2-activate back, 3-blacklist the node
    function updateNodeStatus(string calldata _orgId, string calldata _enodeId,
        uint _action, address _caller) external onlyInterface
    orgAdmin(_caller, _orgId) {
        nodes.updateNodeStatus(_enodeId, _orgId, _action);
    }

    /// @notice function to fetch network boot status
    /// @return bool network boot status
    function getNetworkBootStatus() external view
    returns (bool){
        return networkBoot;
    }

    /// @notice function to fetch detail of any pending approval activities
    /// @notice for network admin organization
    /// @param _orgId unique id of the organization to which the account belongs
    function getPendingOp(string calldata _orgId) external view
    returns (string memory, string memory, address, uint){
        return voter.getPendingOpDetails(_orgId);
    }

    /// @notice function to assigns a role id to the account given account
    /// @notice can be executed by org admin account only
    /// @param _account account id
    /// @param _orgId organization id to which the account belongs
    /// @param _roleId role id to be assigned to the account
    function assignAccountRole(address _account, string memory _orgId,
        string memory _roleId, address _caller) public
    onlyInterface
    orgAdmin(_caller, _orgId)
    orgApproved(_orgId) {
        require(validateAccount(_account, _orgId) == true, "Operation cannot be performed");
        require(_roleExists(_roleId, _orgId) == true, "role does not exists");
        bool admin = roles.isAdminRole(_roleId, _orgId, _getUltimateParent(_orgId));
        accounts.assignAccountRole(_account, _orgId, _roleId, admin);
    }

    /// @notice function to check if passed account is an network admin account
    /// @param _account account id
    /// @return true/false
    function isNetworkAdmin(address _account) public view
    returns (bool){
        return (keccak256(abi.encode(accounts.getAccountRole(_account))) == keccak256(abi.encode(adminRole)));
    }

    /// @notice function to check if passed account is an org admin account
    /// @param _account account id
    /// @param _orgId organization id
    /// @return true/false
    function isOrgAdmin(address _account, string memory _orgId) public view
    returns (bool){
        if (accounts.checkOrgAdmin(_account, _orgId, _getUltimateParent(_orgId))) {
            return true;
        }
        return roles.isAdminRole(accounts.getAccountRole(_account), _orgId,
            _getUltimateParent(_orgId));
    }

    /// @notice function to validate the account for access change operation
    /// @param _account account id
    /// @param _orgId organization id
    /// @return true/false
    function validateAccount(address _account, string memory _orgId) public view
    returns (bool){
        return (accounts.validateAccount(_account, _orgId));
    }

    /// @notice function to update the voter list at network level. this will
    /// @notice be called whenever an account is assigned a network admin role
    /// @notice or an account having network admin role is being assigned
    /// @notice different role
    /// @param _orgId org id to which the account belongs
    /// @param _account account which needs to be added/removed as voter
    /// @param _add bool indicating if its an add or delete operation
    function updateVoterList(string memory _orgId, address _account, bool _add) internal {
        if (_add) {
            voter.addVoter(_orgId, _account);
        }
        else {
            voter.deleteVoter(_orgId, _account);
        }
    }

    /// @notice whenever a network admin account votes on a pending item, this
    /// @notice function processes the vote.
    /// @param _orgId org id of the caller
    /// @param _caller account which approving the operation
    /// @param _pendingOp operation for which the approval is being done
    /// @dev the list of pending ops are managed in voter manager contract
    function processVote(string memory _orgId, address _caller, uint _pendingOp) internal
    returns (bool){
        return voter.processVote(_orgId, _caller, _pendingOp);
    }

    /// @notice returns various permissions policy related parameters
    /// @return adminOrg admin org id
    /// @return adminRole default network admin role
    /// @return orgAdminRole default org admin role
    /// @return networkBoot network boot status
    function getPolicyDetails() external view
    returns (string memory, string memory, string memory, bool){
        return (adminOrg, adminRole, orgAdminRole, networkBoot);
    }

    /// @notice checks if the passed org exists or not
    /// @param _orgId org id
    /// @return true/false
    function _checkOrgExists(string memory _orgId) internal view
    returns (bool){
        return org.checkOrgExists(_orgId);
    }

    /// @notice checks if the passed org is in approved status
    /// @param _orgId org id
    /// @return true/false
    function checkOrgApproved(string memory _orgId) internal view
    returns (bool){
        return org.checkOrgStatus(_orgId, 2);
    }

    /// @notice checks if the passed org is in the status passed
    /// @param _orgId org id
    /// @param _status status to be checked for
    /// @return true/false
    function _checkOrgStatus(string memory _orgId, uint _status) internal view
    returns (bool){
        return org.checkOrgStatus(_orgId, _status);
    }

    /// @notice checks if org admin account exists for the passed org id
    /// @param _orgId org id
    /// @return true/false
    function _checkOrgAdminExists(string memory _orgId) internal view
    returns (bool){
        return accounts.orgAdminExists(_orgId);
    }

    /// @notice checks if role id exists for the passed org_id
    /// @param _roleId role id
    /// @param _orgId org id
    /// @return true/false
    function _roleExists(string memory _roleId, string memory _orgId) internal view
    returns (bool){
        return roles.roleExists(_roleId, _orgId, _getUltimateParent(_orgId));
    }

    /// @notice checks if the role id for the org is a voter role
    /// @param _roleId role id
    /// @param _orgId org id
    /// @return true/false
    function _isVoterRole(string memory _roleId, string memory _orgId) internal view
    returns (bool){
        return roles.isVoterRole(_roleId, _orgId, _getUltimateParent(_orgId));
    }

    /// @notice returns the ultimate parent for a given org id
    /// @param _orgId org id
    /// @return ultimate parent org id
    function _getUltimateParent(string memory _orgId) internal view
    returns (string memory){
        return org.getUltimateParent(_orgId);
    }

}