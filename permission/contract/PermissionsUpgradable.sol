pragma solidity ^0.5.3;

import "./PermissionsInterface.sol";

contract PermissionsUpgradable {

    address private custodian;
    address private permImpl;
    address private permInterface;

    constructor (address _custodian) public
    {
        custodian = _custodian;
    }

    modifier onlyCustodian {
        require(msg.sender == custodian);
        _;
    }

    function init(address _permInterface, address _permImpl) external
    onlyCustodian
    {
        permImpl = _permImpl;
        permInterface = _permInterface;
        setImpl(permImpl);
    }

    // custodian can potentially become a contract
    // implementation change and custodian change are sending from custodian
    function confirmImplChange(address _proposedImpl) public
    onlyCustodian
    {
        permImpl = _proposedImpl;
        setImpl(permImpl);
    }

    function getCustodian() public view returns (address)
    {
        return custodian;
    }

    function getPermImpl() public view returns (address)
    {
        return permImpl;
    }

    function getPermInterface() public view returns (address)
    {
        return permInterface;
    }

    function setImpl(address _permImpl) private
    {
        PermissionsInterface(permInterface).setPermImplementation(_permImpl);
    }

}