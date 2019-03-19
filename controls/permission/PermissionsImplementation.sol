pragma solidity ^0.5.3;

import "./RoleManager.sol";
import "./AccountManager.sol";
import "./VoterManager.sol";
import "./NodeManager.sol";
import "./OrgManager.sol";
import "./PermissionsUpgradable.sol";

contract PermissionsImplementation {
    AccountManager private accounts;
    RoleManager private roles;
    VoterManager private voter;
    NodeManager private nodes;
    OrgManager private org;
    PermissionsUpgradable private permUpgradable;

    event Dummy(string _msg);

    string private adminOrg;
    string private adminRole;
    string private orgAdminRole;

    uint private fullAccess = 3;

    // checks if first time network boot up has happened or not
    bool private networkBoot = false;

    modifier onlyProxy
    {
        require(msg.sender == permUpgradable.getPermInterface(), "can be called by proxy only");
        _;
    }

    // Checks if the given network boot up is pending exists
    modifier networkBootUpPending()
    {
        require(networkBoot == false, "Network boot up completed");
        _;
    }

    // Checks if the given network boot up is pending exists
    modifier networkBootUpDone()
    {
        require(networkBoot == true, "Network boot not complete");
        _;
    }

    modifier networkAdmin(address _account) {
        require(isNetworkAdmin(_account) == true, "Not an network admin");
        _;
    }

    modifier orgAdmin(address _account, string memory _orgId) {
        require(isOrgAdmin(_account, _orgId) == true, "Not an org admin");
        _;
    }

    modifier orgNotExists(string memory _orgId) {
        require(org.checkOrgExists(_orgId) == false, "Org already exists");
        _;
    }

    modifier orgExists(string memory _orgId) {
        require(org.checkOrgExists(_orgId) == true, "Org does not exists");
        _;
    }

    modifier orgApproved(string memory _orgId) {
        require(org.checkOrgStatus(_orgId, 2) == true, "Org not approved");
        _;
    }

    constructor (address _permUpgradable) public {
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    function setPolicy(string calldata _nwAdminOrg, string calldata _nwAdminRole, string calldata _oAdminRole) external
    onlyProxy
    networkBootUpPending()
    {
        adminOrg = _nwAdminOrg;
        adminRole = _nwAdminRole;
        orgAdminRole = _oAdminRole;
    }

    function init(address _orgManager, address _rolesManager, address _acctManager, address _voterManager, address _nodeManager) external
    onlyProxy
    networkBootUpPending()
    {
        org = OrgManager(_orgManager);
        roles = RoleManager(_rolesManager);
        accounts = AccountManager(_acctManager);
        voter = VoterManager(_voterManager);
        nodes = NodeManager(_nodeManager);

        org.addAdminOrg(adminOrg);
        roles.addRole(adminRole, adminOrg, fullAccess, true);
        accounts.setDefaults(adminRole, orgAdminRole);
    }

    function addAdminNodes(string calldata _enodeId) external
    onlyProxy
    networkBootUpPending()
    {
        nodes.addNode(_enodeId, adminOrg);
        nodes.approveNode(_enodeId);
    }

    function addAdminAccounts(address _acct) external
    onlyProxy
    networkBootUpPending()
    {
        // add the account as a voter for the admin org
        voter.addVoter(adminOrg, _acct);
        // add the account as an account with full access into the admin org
        accounts.addNWAdminAccount(_acct, adminOrg);
    }

    // update the network boot status as true
    function updateNetworkBootStatus() external
    onlyProxy
    networkBootUpPending()
    returns (bool)
    {
        networkBoot = true;
        return networkBoot;
    }

//    Get network boot status
    function getNetworkBootStatus() external view
    returns (bool)
    {
        return networkBoot;
    }

    // function for adding a new master org
    function addOrg(string calldata _orgId, string calldata _enodeId, address _caller) external
    onlyProxy
    networkBootUpDone()
    orgNotExists(_orgId)
    networkAdmin(_caller)
    {
        voter.addVotingItem(adminOrg, _orgId, _enodeId, address(0), 1);
        org.addOrg(_orgId);
        nodes.addNode(_enodeId, _orgId);
    }

    function approveOrg(string calldata _orgId, string calldata _enodeId, address _caller) external
    onlyProxy
    networkBootUpDone()
    networkAdmin(_caller)
    {
        require(org.checkOrgStatus(_orgId, 1) == true, "Nothing to approve");
        if ((voter.processVote(adminOrg, _caller, 1))) {
            org.approveOrg(_orgId);
            nodes.approveNode(_enodeId);
        }
    }

//    function updateOrgStatus(string calldata _orgId, uint _status) external
//    onlyProxy
//    networkBootUpDone()
//    orgExists(_orgId)
//    networkAdmin(msg.sender)
//    {
//        require ((_status == 3 || _status == 5), "Operation not allowed");
//        uint reqStatus;
//        uint pendingOp;
//        if (_status == 3) {
//            reqStatus = 2;
//            pendingOp = 2;
//        }
//        else if (_status == 5) {
//            reqStatus = 4;
//            pendingOp = 3;
//        }
//        require(org.checkOrgStatus(_orgId, reqStatus) == true, "Operation not allowed");
//        org.updateOrg(_orgId, _status);
//        voter.addVotingItem(adminOrg, _orgId, "", address(0), pendingOp);
//    }
//
//    function approveOrgStatus(string calldata _orgId, uint _status) external
//    onlyProxy
//    networkBootUpDone()
//    orgExists(_orgId)
//    networkAdmin(msg.sender)
//    {
//        require ((_status == 3 || _status == 5), "Operation not allowed");
//        uint pendingOp;
//        if (_status == 3) {
//            pendingOp = 2;
//        }
//        else if (_status == 5) {
//            pendingOp = 3;
//        }
//        require(org.checkOrgStatus(_orgId, _status) == true, "Operation not allowed");
//        if ((voter.processVote(adminOrg, msg.sender, pendingOp))) {
//            org.approveOrgStatusUpdate(_orgId, _status);
//        }
//    }
    // returns org and master org details based on org index
    function getOrgInfo(uint _orgIndex) external view
    returns (string memory, uint)

    {
        return org.getOrgInfo(_orgIndex);
    }

    // Role related functions
    function addNewRole(string calldata _roleId, string calldata _orgId, uint _access, bool _voter) external
    onlyProxy
    orgApproved(_orgId)
    orgAdmin(msg.sender, _orgId)
    {
        //add new roles can be created by org admins only
        roles.addRole(_roleId, _orgId, _access, _voter);
    }

    function removeRole(string calldata _roleId, string calldata _orgId) external
    onlyProxy
    orgApproved(_orgId)
    orgAdmin(msg.sender, _orgId)
    {
        roles.removeRole(_roleId, _orgId);
    }

    function getRoleDetails(string calldata _roleId, string calldata _orgId) external view
    returns (string memory, string memory, uint, bool, bool)
    {
        return roles.getRoleDetails(_roleId, _orgId);

    }

    // Org voter related functions
    function getNumberOfVoters(string calldata _orgId) external view
    returns (uint){

        return voter.getNumberOfValidVoters(_orgId);
    }

    function checkIfVoterExists(string calldata _orgId, address _acct) external view
    returns (bool)
    {
        return voter.checkIfVoterExists(_orgId, _acct);
    }

    function getVoteCount(string calldata _orgId) external view returns (uint, uint)
    {
        return voter.getVoteCount(_orgId);
    }

    function getPendingOp(string calldata _orgId) external view
    returns (string memory, string memory, address, uint)
    {
        return voter.getPendingOpDetails(_orgId);
    }

    function assignOrgAdminAccount(string calldata _orgId, address _account, address _caller) external
    onlyProxy
    networkBootUpDone()
    networkAdmin(_caller)
    orgExists(_orgId)
    {
        // check if orgAdmin already exists if yes then op cannot be performed
        require(accounts.orgAdminExists(_orgId) != true, "org admin exists");
        // assign the account org admin role and propose voting
        accounts.assignAccountRole(_account, _orgId, orgAdminRole);
        //add voting item
        voter.addVotingItem(adminOrg, _orgId, "", _account, 4);
    }

    function approveOrgAdminAccount(address _account, address _caller) external
    onlyProxy
    networkBootUpDone()
    networkAdmin(_caller)
    {
        require(isNetworkAdmin(_caller) == true, "can be called from network admin only");
        if ((voter.processVote(adminOrg, _caller, 4))) {
            accounts.approveOrgAdminAccount(_account);
        }
    }


    function assignAccountRole(address _acct, string memory _orgId, string memory _roleId) public
    onlyProxy
    networkBootUpDone()
    orgApproved(_orgId)
    orgAdmin(msg.sender, _orgId)
    {
        // check if the account is part of another org. If yes then op cannot be done
        require(validateAccount(_acct, _orgId) == true, "Operation cannot be performed");
        // check if role is existing for the org. if yes the op can be done
        require(roles.roleExists(_roleId, _orgId) == true, "role does not exists");
        bool newRoleVoter = roles.isVoterRole(_roleId, _orgId);
        // check the role of the account. if the current role is voter and new role is also voter
        // voterlist change is not required. else voter list needs to be changed
        string memory acctRole = accounts.getAccountRole(_acct);
        if (keccak256(abi.encodePacked(acctRole)) == keccak256(abi.encodePacked("NONE"))) {
            //new account
            if (newRoleVoter) {
                // add to voter list
                voter.addVoter(_orgId, _acct);
            }
        }
        else {
            bool currRoleVoter = roles.isVoterRole(acctRole, _orgId);
            if (!(currRoleVoter && newRoleVoter)) {
                if (newRoleVoter) {
                    // add to voter list
                    voter.addVoter(_orgId, _acct);
                }
                else {
                    // delete from voter list
                    voter.deleteVoter(_orgId, _acct);
                }
            }
        }
        accounts.assignAccountRole(_acct, _orgId, _roleId);
    }

    function addNode(string calldata _orgId, string calldata _enodeId) external
    onlyProxy
    networkBootUpDone()
    orgApproved(_orgId)
    orgAdmin(msg.sender, _orgId)
    {
        // check that the node is not part of another org
        require(getNodeStatus(_enodeId) == 0, "Node present already");
        nodes.addOrgNode(_enodeId, _orgId);
    }

    function getNodeStatus(string memory _enodeId) public view
    returns (uint)
    {
        return (nodes.getNodeStatus(_enodeId));
    }

    function isNetworkAdmin(address _account) public view
    returns (bool)
    {
        return (keccak256(abi.encodePacked(accounts.getAccountRole(_account))) == keccak256(abi.encodePacked(adminRole)));
    }

    function isOrgAdmin(address _account, string memory _orgId) public view
    returns (bool)
    {
        return (accounts.checkOrgAdmin(_account, _orgId));
    }

    function validateAccount(address _account, string memory _orgId) public view
    returns (bool)
    {
        return (accounts.valAcctAccessChange(_account, _orgId));
    }

    function getAccountDetails(address _acct) external view
    returns (address, string memory, string memory, uint, bool)
    {
        return  accounts.getAccountDetails(_acct);
    }

}