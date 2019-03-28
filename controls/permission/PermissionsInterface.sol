pragma solidity ^0.5.3;

import "./PermissionsImplementation.sol";
import "./PermissionsUpgradable.sol";


contract PermissionsInterface {
    PermissionsImplementation private permImplementation;
    PermissionsUpgradable private permUpgradable;
    address private permImplUpgradeable;

    event Dummy(string _msg);

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

    function getPermissionsImpl() public view returns(address)
    {
        return address(permImplementation);
    }

    function setPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole, string calldata _oAdminRole) external
    {
        permImplementation.setPolicy(_nwAdminOrg, _nwAdminRole, _oAdminRole);
    }

    function init(address _orgManager, address _rolesManager, address _acctManager, address _voterManager, address _nodeManager) external
    {
        permImplementation.init(_orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager);
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
    function addOrg(string calldata _orgId, string calldata _enodeId) external
    {
        permImplementation.addOrg(_orgId, _enodeId, msg.sender);
    }

    function approveOrg(string calldata _orgId, string calldata _enodeId) external
    {
        permImplementation.approveOrg(_orgId, _enodeId, msg.sender);
    }

    function updateOrgStatus(string calldata _orgId, uint _status) external
    {
        permImplementation.updateOrgStatus(_orgId, _status);
    }

    function approveOrgStatus(string calldata _orgId, uint _status) external
    {
        permImplementation.approveOrgStatus(_orgId, _status);
    }
    // returns org and master org details based on org index
    function getOrgInfo(uint _orgIndex) external view returns (string memory, uint)
    {
        return permImplementation.getOrgInfo(_orgIndex);
    }

    // Role related functions
    function addNewRole(string calldata _roleId, string calldata _orgId, uint _access, bool _voter) external
    {
        permImplementation.addNewRole(_roleId, _orgId, _access, _voter, msg.sender);
    }

    function removeRole(string calldata _roleId, string calldata _orgId) external
    {
        permImplementation.removeRole(_roleId, _orgId);
    }

    function getRoleDetails(string calldata _roleId, string calldata _orgId) external view returns (string memory, string memory, uint, bool, bool)
    {
        return permImplementation.getRoleDetails(_roleId, _orgId);
    }

    // Org voter related functions
    function getNumberOfVoters(string calldata _orgId) external view returns (uint)
    {
        return permImplementation.getNumberOfVoters(_orgId);
    }


    function checkIfVoterExists(string calldata _orgId, address _acct) external view returns (bool)
    {
        return permImplementation.checkIfVoterExists(_orgId, _acct);
    }


    function getVoteCount(string calldata _orgId) external view returns (uint, uint)
    {
        return permImplementation.getVoteCount(_orgId);
    }

    function getPendingOp(string calldata _orgId) external view returns (string memory, string memory, address, uint)
    {
        return permImplementation.getPendingOp(_orgId);
    }

    function assignOrgAdminAccount(string calldata _orgId, address _account) external
    {
        permImplementation.assignOrgAdminAccount(_orgId, _account, msg.sender);

    }

    function approveOrgAdminAccount(address _account) external
    {
        permImplementation.approveOrgAdminAccount(_account, msg.sender);

    }

    function assignAccountRole(address _acct, string memory _orgId, string memory _roleId) public
    {
        permImplementation.assignAccountRole(_acct, _orgId, _roleId, msg.sender);

    }

    function addNode(string calldata _orgId, string calldata _enodeId) external
    {
        permImplementation.addNode(_orgId, _enodeId, msg.sender);

    }

    function updateNodeStatus(string calldata _orgId, string calldata _enodeId, uint _status) external
    {
        permImplementation.updateNodeStatus(_orgId, _enodeId, _status, msg.sender);
    }

    function getNodeStatus(string memory _enodeId) public view returns (uint)
    {
        return permImplementation.getNodeStatus(_enodeId);
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

    function getAccountDetails(address _acct) external view returns (address, string memory, string memory, uint, bool)
    {
        return permImplementation.getAccountDetails(_acct);
    }

}