pragma solidity ^0.8.4;

import "./openzeppelin-v5/Initializable.sol";
import "./openzeppelin-v5/EnumerableSet.sol";
import "./PermissionsUpgradable.sol";

/** @title Contract whitelist manager contract
  * @notice This contract holds implementation logic for all contract whitelisting
    functionality. This can be called only by the implementation contract only.
    there are few view functions exposed as public and can be called directly.
    these are invoked by quorum for populating permissions data in cache
  */
contract ContractWhitelistManager is Initializable {
    bytes32 private constant CONTRACT_WHITELIST_STORAGE_POSITION =
        keccak256(
            abi.encode(
                uint256(keccak256("quorum.storage.contractwhitelist")) - 1
            )
        ) & ~bytes32(uint256(0xff));

    using EnumerableSet for EnumerableSet.AddressSet;

    struct ContractWhitelistStorage {
        PermissionsUpgradable permUpgradable;
        EnumerableSet.AddressSet contractAddressSet;
    }

    // contract whitelist events
    event ContractWhitelistAdded(address indexed _contractAddr);
    event ContractWhitelistRevoked(address indexed _contractAddr);

    /** @notice confirms that the caller is the address of implementation
        contract
      */
    modifier onlyImplementation() {
        require(
            msg.sender ==
                contractWhitelistStorage().permUpgradable.getPermImpl(),
            "invalid caller"
        );
        _;
    }

    /// @notice initialized only once. sets the permissions upgradable address
    function initialize(address _permUpgradable) external initializer {
        require(_permUpgradable != address(0x0), "Cannot set to empty address");
        contractWhitelistStorage().permUpgradable = PermissionsUpgradable(
            _permUpgradable
        );
    }

    /** @notice returns the total number of whitelisted contracts
     * @return total number whitelisted contracts
     */
    function getNumberOfWhitelistedContracts() external view returns (uint) {
        return contractWhitelistStorage().contractAddressSet.length();
    }

    /** @notice returns the array of whitelisted contracts
     * @return whitelisted contracts
     */
    function getWhitelistedContracts() external view returns (address[] memory) {
        return contractWhitelistStorage().contractAddressSet.values();
    }

    /** @notice returns the contract whitelist details a given contract whitelist index
     * @param  _contract contract address
     * @return status bool whether contract is whitelisted
     */
    function isContractWhitelisted(
        address _contract
    ) external view returns (bool) {
        return
            contractWhitelistStorage().contractAddressSet.contains(_contract);
    }

    /** @notice function to add whitelisted contract
     * @param _contract - contract address
     */
    function addWhitelist(address _contract) external onlyImplementation {
        // Check if contract already exists
        ContractWhitelistStorage storage cs = contractWhitelistStorage();
        require(
            !cs.contractAddressSet.contains(_contract),
            "whitelist already exists"
        );
        cs.contractAddressSet.add(_contract);
        emit ContractWhitelistAdded(_contract);
    }

    /** @notice function to revoke whitelisted contract by address
     * @param _contract - contract address
     */
    function revokeWhitelist(address _contract) external onlyImplementation {
        // Check if contract already exists
        ContractWhitelistStorage storage cs = contractWhitelistStorage();
        require(
            cs.contractAddressSet.contains(_contract),
            "whitelist does not exist"
        );
        cs.contractAddressSet.remove(_contract);
        emit ContractWhitelistRevoked(_contract);
    }

    function contractWhitelistStorage()
        internal
        pure
        returns (ContractWhitelistStorage storage cs)
    {
        // Specifies a random position from a hash of a string
        bytes32 storagePosition = CONTRACT_WHITELIST_STORAGE_POSITION;
        // Set the position of our struct in contract storage
        assembly {
            cs.slot := storagePosition
        }
    }
}