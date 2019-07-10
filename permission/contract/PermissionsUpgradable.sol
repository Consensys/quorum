pragma solidity ^0.5.3;

import "./PermissionsInterface.sol";

contract PermissionsUpgradable {

    address private guardian;
    address private permImpl;
    address private permInterface;

    // sets the guardian account as part of constructor
    // only this account will be able to change the implementation contract address
    constructor (address _guardian) public
    {
        guardian = _guardian;
    }

    modifier onlyCustodian {
        require(msg.sender == guardian);
        _;
    }

    // executed by guardian, links interface and implementation contract addresses
    function init(address _permInterface, address _permImpl) external
    onlyCustodian
    {
        permImpl = _permImpl;
        permInterface = _permInterface;
        setImpl(permImpl);
    }


    // guardian can potentially become a contract
    // implementation change and guardian change are sending from guardian
    function confirmImplChange(address _proposedImpl) public
    onlyCustodian
    {
        // read the details from current implementation
        (string memory adminOrg, string memory adminRole, string memory orgAdminRole, bool bootStatus) = PermissionsImplementation(permImpl).getPolicyDetails();
        setPolicy(_proposedImpl, adminOrg, adminRole, orgAdminRole, bootStatus);
        // set these values in new implementation
        permImpl = _proposedImpl;
        setImpl(permImpl);
    }

    function getCustodian() public view returns (address)
    {
        return guardian;
    }

    function getPermImpl() public view returns (address)
    {
        return permImpl;
    }

    function getPermInterface() public view returns (address)
    {
        return permInterface;
    }

    function setPolicy(address _permImpl, string memory _adminOrg, string memory _adminRole, string memory _orgAdminRole, bool _bootStatus) private
    {
        PermissionsImplementation(_permImpl).setMigrationPolicy(_adminOrg, _adminRole, _orgAdminRole, _bootStatus);
    }

    function setImpl(address _permImpl) private
    {
        PermissionsInterface(permInterface).setPermImplementation(_permImpl);
    }

}