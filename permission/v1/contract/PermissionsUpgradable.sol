pragma solidity ^0.5.3;

import "./PermissionsInterface.sol";

/** @title Permissions Upgradable Contract
  * @notice This contract holds the address of current permissions implementation
    contract. The contract is owned by a guardian account. Only the
    guardian account can change the implementation contract address as
    business needs.
  */
contract PermissionsUpgradable {

    address private guardian;
    address private permImpl;
    address private permInterface;
    // initDone ensures that init can be called only once
    bool private initDone;

    /** @notice constructor
      * @param _guardian account address
      */
    constructor (address _guardian) public{
        guardian = _guardian;
        initDone = false;
    }

    /** @notice confirms that the caller is the guardian account
    */
    modifier onlyGuardian {
        require(msg.sender == guardian, "invalid caller");
        _;
    }

    /** @notice executed by guardian. Links interface and implementation contract
        addresses. Can be executed by guardian account only
      * @param _permInterface permissions interface contract address
      * @param _permImpl implementation contract address
      */
    function init(address _permInterface, address _permImpl) external
    onlyGuardian {
        require(!initDone, "can be executed only once");
        permImpl = _permImpl;
        permInterface = _permInterface;
        _setImpl(permImpl);
        initDone = true;
    }

    /** @notice changes the implementation contract address to the new address
        address passed. Can be executed by guardian account only
      * @param _proposedImpl address of the new permissions implementation contract
      */
    function confirmImplChange(address _proposedImpl) public
    onlyGuardian {
        // The policy details needs to be carried forward from existing
        // implementation to new. So first these are read from existing
        // implementation and then updated in new implementation
        (string memory adminOrg, string memory adminRole, string memory orgAdminRole, bool bootStatus) = PermissionsImplementation(permImpl).getPolicyDetails();
        _setPolicy(_proposedImpl, adminOrg, adminRole, orgAdminRole, bootStatus);
        permImpl = _proposedImpl;
        _setImpl(permImpl);
    }

    /** @notice function to fetch the guardian account address
      * @return _guardian guardian account address
      */
    function getGuardian() public view returns (address) {
        return guardian;
    }

    /** @notice function to fetch the current implementation address
      * @return permissions implementation contract address
      */
    function getPermImpl() public view returns (address) {
        return permImpl;
    }
    /** @notice function to fetch the interface address
      * @return permissions interface contract address
      */
    function getPermInterface() public view returns (address) {
        return permInterface;
    }

    /** @notice function to set the permissions policy details in the
        permissions implementation contract
      * @param _permImpl permissions implementation contract address
      * @param _adminOrg name of admin organization
      * @param _adminRole name of the admin role
      * @param _orgAdminRole name of default organization admin role
      * @param _bootStatus network boot status
      */
    function _setPolicy(address _permImpl, string memory _adminOrg, string memory _adminRole, string memory _orgAdminRole, bool _bootStatus) private {
        PermissionsImplementation(_permImpl).setMigrationPolicy(_adminOrg, _adminRole, _orgAdminRole, _bootStatus);
    }

    /** @notice function to set the permissions implementation contract address
        in the permissions interface contract
      * @param _permImpl permissions implementation contract address
      */
    function _setImpl(address _permImpl) private {
        PermissionsInterface(permInterface).setPermImplementation(_permImpl);
    }

}