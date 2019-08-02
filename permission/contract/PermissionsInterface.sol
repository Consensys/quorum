pragma solidity ^0.5.3;

import "./PermissionsImplementation.sol";
import "./PermissionsUpgradable.sol";

/** @title Permissions Interface Contract
  * @notice This contract is the interface for permissions implementation
    contract. for any call, it forwards the call to the implementation
    contract
  */
contract PermissionsInterface {
    PermissionsImplementation private permImplementation;
    PermissionsUpgradable private permUpgradable;
    address private permImplUpgradeable;

    /** @notice constructor
      * @param _permImplUpgradeable permissions upgradable contract address
      */
    constructor(address _permImplUpgradeable) public {
        permImplUpgradeable = _permImplUpgradeable;
    }

    /** @notice confirms that the caller is the address of upgradable
        contract
      */
    modifier onlyUpgradeable {
        require(msg.sender == permImplUpgradeable, "invalid caller");
        _;
    }

    /** @notice interface for setting the permissions policy in implementation
      * @param _nwAdminOrg network admin organization id
      * @param _nwAdminRole default network admin role id
      * @param _oAdminRole default organization admin role id
      */
    function setPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole,
        string calldata _oAdminRole) external {
        permImplementation.setPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole);
    }

    /** @notice interface to initializes the breadth and depth values for
        sub organization management
      * @param _breadth controls the number of sub org a parent org can have
      * @param _depth controls the depth of nesting allowed for sub orgs
      */
    function init(uint256 _breadth, uint256 _depth) external {
        permImplementation.init(_breadth, _depth);
    }

    /** @notice interface to add new node to an admin organization
      * @param _enodeId full enode id of the node to be added
      */
    function addAdminNode(string calldata _enodeId) external {
        permImplementation.addAdminNode(_enodeId);
    }

    /** @notice interface to add accounts to an admin organization
      * @param _acct account address to be added
      */
    function addAdminAccount(address _acct) external {
        permImplementation.addAdminAccount(_acct);
    }

    /** @notice interface to update network boot up status
      * @return bool true or false
      */
    function updateNetworkBootStatus() external
    returns (bool)
    {
        return permImplementation.updateNetworkBootStatus();
    }

    /** @notice interface to fetch network boot status
      * @return bool network boot status
      */
    function getNetworkBootStatus() external view returns (bool){
        return permImplementation.getNetworkBootStatus();
    }

    /** @notice interface to add a new organization to the network
      * @param _orgId unique organization id
      * @param _enodeId full enode id linked to the organization
      * @param _account account id. this will have the org admin privileges
      */
    function addOrg(string calldata _orgId, string calldata _enodeId,
        address _account) external {
        permImplementation.addOrg(_orgId, _enodeId, _account, msg.sender);
    }

    /** @notice interface to approve a newly added organization
      * @param _orgId unique organization id
      * @param _enodeId full enode id linked to the organization
      * @param _account account id this will have the org admin privileges
      */
    function approveOrg(string calldata _orgId, string calldata _enodeId,
        address _account) external {
        permImplementation.approveOrg(_orgId, _enodeId, _account, msg.sender);
    }

    /** @notice interface to add sub org under an org
      * @param _pOrgId parent org id under which the sub org is being added
      * @param _orgId unique id for the sub organization
      * @param _enodeId full enode id linked to the sjb organization
      */
    function addSubOrg(string calldata _pOrgId, string calldata _orgId,
        string calldata _enodeId) external {
        permImplementation.addSubOrg(_pOrgId, _orgId, _enodeId, msg.sender);
    }

    /** @notice interface to update the org status
      * @param _orgId unique id of the organization
      * @param _action 1 for suspending an org and 2 for revoke of suspension
      */
    function updateOrgStatus(string calldata _orgId, uint256 _action) external {
        permImplementation.updateOrgStatus(_orgId, _action, msg.sender);
    }

    /** @notice interface to approve org status change
      * @param _orgId unique id for the sub organization
      * @param _action 1 for suspending an org and 2 for revoke of suspension
      */
    function approveOrgStatus(string calldata _orgId, uint256 _action) external {
        permImplementation.approveOrgStatus(_orgId, _action, msg.sender);
    }

    /** @notice interface to add a new role definition to an organization
      * @param _roleId unique id for the role
      * @param _orgId unique id of the organization to which the role belongs
      * @param _access account access type for the role
      * @param _voter bool indicates if the role is voter role or not
      * @param _admin bool indicates if the role is an admin role
      * @dev account access type can have of the following four values:
            0 - Read only
            1 - Transact access
            2 - Contract deployment access. Can transact as well
            3 - Full access
      */
    function addNewRole(string calldata _roleId, string calldata _orgId,
        uint256 _access, bool _voter, bool _admin) external {
        permImplementation.addNewRole(_roleId, _orgId, _access, _voter, _admin, msg.sender);
    }

    /** @notice interface to remove a role definition from an organization
      * @param _roleId unique id for the role
      * @param _orgId unique id of the organization to which the role belongs
      */
    function removeRole(string calldata _roleId, string calldata _orgId) external {
        permImplementation.removeRole(_roleId, _orgId, msg.sender);
    }

    /** @notice interface to assign network admin/org admin role to an account
        this can be executed by network admin accounts only
      * @param _orgId unique id of the organization to which the account belongs
      * @param _account account id
      * @param _roleId role id to be assigned to the account
      */
    function assignAdminRole(string calldata _orgId, address _account,
        string calldata _roleId) external {
        permImplementation.assignAdminRole(_orgId, _account, _roleId, msg.sender);

    }
    /** @notice interface to approve network admin/org admin role assigment
        this can be executed by network admin accounts only
      * @param _orgId unique id of the organization to which the account belongs
      * @param _account account id
      */
    function approveAdminRole(string calldata _orgId, address _account) external {
        permImplementation.approveAdminRole(_orgId, _account, msg.sender);

    }

    /** @notice interface to update account status
        this can be executed by org admin accounts only
      * @param _orgId unique id of the organization to which the account belongs
      * @param _account account id
      * @param _action 1-suspending 2-activating back 3-blacklisting
      */
    function updateAccountStatus(string calldata _orgId, address _account,
        uint256 _action) external {
        permImplementation.updateAccountStatus(_orgId, _account, _action, msg.sender);
    }

    /** @notice interface to add a new node to the organization
      * @param _orgId unique id of the organization to which the account belongs
      * @param _enodeId full enode id being dded to the org
      */
    function addNode(string calldata _orgId, string calldata _enodeId) external {
        permImplementation.addNode(_orgId, _enodeId, msg.sender);

    }

    /** @notice interface to update node status
      * @param _orgId unique id of the organization to which the account belongs
      * @param _enodeId full enode id being dded to the org
      * @param _action 1-deactivate, 2-activate back, 3-blacklist the node
      */
    function updateNodeStatus(string calldata _orgId, string calldata _enodeId,
        uint256 _action) external {
        permImplementation.updateNodeStatus(_orgId, _enodeId, _action, msg.sender);
    }

    /** @notice interface to initiate blacklisted node recovery
      * @param _orgId unique id of the organization to which the account belongs
      * @param _enodeId full enode id being recovered
      */
    function startBlacklistedNodeRecovery(string calldata _orgId, string calldata _enodeId)
    external {
        permImplementation.startBlacklistedNodeRecovery(_orgId, _enodeId, msg.sender);
    }

    /** @notice interface to approve blacklisted node recoevry
      * @param _orgId unique id of the organization to which the account belongs
      * @param _enodeId full enode id being recovered
      */
    function approveBlacklistedNodeRecovery(string calldata _orgId, string calldata _enodeId)
    external {
        permImplementation.approveBlacklistedNodeRecovery(_orgId, _enodeId, msg.sender);
    }

    /** @notice interface to initiate blacklisted account recovery
      * @param _orgId unique id of the organization to which the account belongs
      * @param _account account id being recovered
      */
    function startBlacklistedAccountRecovery(string calldata _orgId, address _account)
    external {
        permImplementation.startBlacklistedAccountRecovery(_orgId, _account, msg.sender);
    }

    /** @notice interface to approve blacklisted node recovery
      * @param _orgId unique id of the organization to which the account belongs
      * @param _account account id being recovered
      */
    function approveBlacklistedAccountRecovery(string calldata _orgId, address _account)
    external {
        permImplementation.approveBlacklistedAccountRecovery(_orgId, _account, msg.sender);
    }

    /** @notice interface to fetch detail of any pending approval activities
        for network admin organization
      * @param _orgId unique id of the organization to which the account belongs
      */
    function getPendingOp(string calldata _orgId) external view
    returns (string memory, string memory, address, uint256) {
        return permImplementation.getPendingOp(_orgId);
    }

    /** @notice sets the permissions implementation contract address
        can be called from upgradable contract only
      * @param _permImplementation permissions implementation contract address
      */
    function setPermImplementation(address _permImplementation) external
    onlyUpgradeable {
        permImplementation = PermissionsImplementation(_permImplementation);
    }

    /** @notice returns the address of permissions implementation contract
      * @return permissions implementation contract address
      */
    function getPermissionsImpl() external view returns (address) {
        return address(permImplementation);
    }

    /** @notice interface to assigns a role id to the account give
      * @param _account account id
      * @param _orgId organization id to which the account belongs
      * @param _roleId role id to be assigned to the account
      */
    function assignAccountRole(address _account, string calldata _orgId,
        string calldata _roleId) external {
        permImplementation.assignAccountRole(_account, _orgId, _roleId, msg.sender);

    }

    /** @notice interface to check if passed account is an network admin account
      * @param _account account id
      * @return true/false
      */
    function isNetworkAdmin(address _account) external view returns (bool) {
        return permImplementation.isNetworkAdmin(_account);
    }

    /** @notice interface to check if passed account is an org admin account
      * @param _account account id
      * @param _orgId organization id
      * @return true/false
      */
    function isOrgAdmin(address _account, string calldata _orgId)
    external view returns (bool) {
        return permImplementation.isOrgAdmin(_account, _orgId);
    }

    /** @notice interface to validate the account for access change operation
      * @param _account account id
      * @param _orgId organization id
      * @return true/false
      */
    function validateAccount(address _account, string calldata _orgId)
    external view returns (bool) {
        return permImplementation.validateAccount(_account, _orgId);
    }

}