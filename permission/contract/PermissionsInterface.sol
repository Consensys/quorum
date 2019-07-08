pragma solidity ^0.5.3;

import "./PermissionsImplementation.sol";
import "./PermissionsUpgradable.sol";


contract PermissionsInterface {
    PermissionsImplementation private permImplementation;
    PermissionsUpgradable private permUpgradable;
    address private permImplUpgradeable;

    constructor(address _permImplUpgradeable) public {
        permImplUpgradeable = _permImplUpgradeable;
    }

    modifier onlyUpgradeable {
        require(msg.sender == permImplUpgradeable);
        _;
    }

    function setPermImplementation(address _permImplementation) public
    onlyUpgradeable
    {
        permImplementation = PermissionsImplementation(_permImplementation);
    }

    function getPermissionsImpl() public view returns (address)
    {
        return address(permImplementation);
    }

    function setPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole, string calldata _oAdminRole) external
    {
        permImplementation.setPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole);
    }

    function init(uint _breadth, uint _depth) external
    {
        permImplementation.init(_breadth, _depth);
    }

    function addAdminNodes(string calldata _enodeId) external
    {
        permImplementation.addAdminNodes(_enodeId);
    }

    function addAdminAccounts(address _acct) external
    {
        permImplementation.addAdminAccounts(_acct);
    }

    // update the network boot status as true
    function updateNetworkBootStatus() external
    returns (bool)
    {
        permImplementation.updateNetworkBootStatus();
    }

    //    // Get network boot status
    function getNetworkBootStatus() external view returns (bool)
    {
        return permImplementation.getNetworkBootStatus();
    }

    // function for adding a new master org
    function addOrg(string calldata _orgId, string calldata _enodeId, address _account) external
    {
        permImplementation.addOrg(_orgId, _enodeId, _account, msg.sender);
    }

    function approveOrg(string calldata _orgId, string calldata _enodeId, address _account) external
    {
        permImplementation.approveOrg(_orgId, _enodeId, _account, msg.sender);
    }

    // function for adding a new master org
    function addSubOrg(string calldata _pOrg, string calldata _orgId, string calldata _enodeId) external
    {
        permImplementation.addSubOrg(_pOrg, _orgId, _enodeId, msg.sender);
    }

    function updateOrgStatus(string calldata _orgId, uint _action) external
    {
        permImplementation.updateOrgStatus(_orgId, _action, msg.sender);
    }

    function approveOrgStatus(string calldata _orgId, uint _action) external
    {
        permImplementation.approveOrgStatus(_orgId, _action, msg.sender);
    }

    // Role related functions
    function addNewRole(string calldata _roleId, string calldata _orgId, uint _access, bool _voter, bool _admin) external
    {
        permImplementation.addNewRole(_roleId, _orgId, _access, _voter, _admin, msg.sender);
    }

    function removeRole(string calldata _roleId, string calldata _orgId) external
    {
        permImplementation.removeRole(_roleId, _orgId, msg.sender);
    }

    function assignAdminRole(string calldata _orgId, address _account, string calldata _roleId) external
    {
        permImplementation.assignAdminRole(_orgId, _account, _roleId, msg.sender);

    }

    function approveAdminRole(string calldata _orgId, address _account) external
    {
        permImplementation.approveAdminRole(_orgId, _account, msg.sender);

    }

    function assignAccountRole(address _acct, string memory _orgId, string memory _roleId) public
    {
        permImplementation.assignAccountRole(_acct, _orgId, _roleId, msg.sender);

    }

    function updateAccountStatus(string calldata _orgId, address _account, uint _status) external
    {
        permImplementation.updateAccountStatus(_orgId, _account, _status, msg.sender);
    }

    function addNode(string calldata _orgId, string calldata _enodeId) external
    {
        permImplementation.addNode(_orgId, _enodeId, msg.sender);

    }

    function updateNodeStatus(string calldata _orgId, string calldata _enodeId, uint _action) external
    {
        permImplementation.updateNodeStatus(_orgId, _enodeId, _action, msg.sender);
    }

    function isNetworkAdmin(address _account) public view returns (bool)
    {
        return permImplementation.isNetworkAdmin(_account);
    }

    function isOrgAdmin(address _account, string memory _orgId) public view returns (bool)
    {

        return permImplementation.isOrgAdmin(_account, _orgId);
    }

    function validateAccount(address _account, string memory _orgId) public view returns (bool)
    {
        return permImplementation.validateAccount(_account, _orgId);
    }

    function getPendingOp(string calldata _orgId) external view returns (string memory, string memory, address, uint)
    {
        return permImplementation.getPendingOp(_orgId);
    }

}