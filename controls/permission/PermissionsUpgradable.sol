pragma solidity ^0.5.3;

import "./PermissionsInterface.sol";

contract PermissionsImplUpgradeable {

    address private custodian;
    address private permImpl;
    // store the instances in the contract because upgradeable will setCoinImpl for them
    PermissionsInterface private permInterface;

    constructor (address _custodian, address _permInterface, address _permImpl) public {
        custodian = _custodian;
        permImpl = _permImpl;
        permInterface = PermissionsInterface(_permInterface);
        setImpl(_permImpl);
    }

    modifier onlyCustodian {
        require(msg.sender == custodian);
        _;
    }

    // custodian can potentially become a contract
    // implementation change and custodian change are sending from custodian
    function confirmImplChange(address _proposedImpl) public onlyCustodian {
        permImpl = _proposedImpl;
        setImpl(permImpl);
    }

    function getCustodian() public view returns(address) {
        return custodian;
    }

    function getPermImpl() public view returns(address) {
        return permImpl;
    }

    function setImpl(address _permImpl) private {
        permInterface.setPermImplementation(_permImpl);
    }

}