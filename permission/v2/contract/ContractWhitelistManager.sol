
pragma solidity ^0.8.4;

import "./PermissionsUpgradable.sol";
import "./openzeppelin-v5/Initializable.sol";

/** @title Contract whitelist manager contract
  * @notice This contract holds implementation logic for all contract whitelisting
    functionality. This can be called only by the implementation contract only.
    there are few view functions exposed as public and can be called directly.
    these are invoked by quorum for populating permissions data in cache
  * @dev account status is denoted by a fixed integer value. The values are
    as below:
        0 - Inactive
        1 - Active
        2 - Revoked
  */
contract ContractWhitelistManager is Initializable {

    bytes32 private constant CONTRACT_WHITELIST_STORAGE_POSITION = keccak256(abi.encode(uint256(keccak256("quorum.storage.contractwhitelist")) - 1)) & ~bytes32(uint256(0xff));

    struct ContractWhitelistDetails {
        address contractAddress;
        string contractKey;
        uint8 whitelistStatus;
    }

    struct ContractWhitelistStorage {
        PermissionsUpgradable permUpgradable;
        ContractWhitelistDetails[] contractWhitelist;
        mapping(address => uint) contractIndexByAddress;
        mapping(string => uint) contractIndexByKey;
        uint numContracts;
    }

    // contract whitelist events
    event ContractWhitelistModified(address _contractAddr, string _contractKey, uint8 _status);
    event ContractWhitelistRevoked(address _contractAddr, string _contractKey, uint8 _status);

    /** @notice confirms that the caller is the address of implementation
        contract
      */
    modifier onlyImplementation {
        require(msg.sender == contractWhitelistStorage().permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /// @notice initialized only once. sets the permissions upgradable address
    function initialize(address _permUpgradable) external initializer {
        require(_permUpgradable != address(0x0), "Cannot set to empty address");
        contractWhitelistStorage().permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice returns the total number of whitelisted contracts
      * @return total number whitelisted contracts
      */
    function getNumberOfWhitelistedContracts() external view returns (uint) {
        return contractWhitelistStorage().contractWhitelist.length;
    }

    /** @notice returns the contract whitelist details a given contract whitelist index
      * @param  _cIndex contract index
      * @return contract contract address
      * @return key contract key
      * @return status whitelist status
      */
    function getContractWhitelistDetailsFromIndex(uint _cIndex) external view returns
    (address, string memory, uint status) {
        ContractWhitelistDetails memory details = contractWhitelistStorage().contractWhitelist[_cIndex];
        return (details.contractAddress, details.contractKey, details.whitelistStatus);
    }

    /** @notice function to add/update whitelisted contract
      * @param _key - contract key, i.e. human parsable string representation of an unique identifier for a contract
      * @param _contract - contract address
      */
    function addWhitelist(string calldata _key, address _contract) external
    onlyImplementation {
        // Check if contract already exists
        ContractWhitelistStorage storage cs = contractWhitelistStorage();
        if (cs.contractIndexByKey[_key] != 0) {
            uint256 cIndex = _getContractIndexByKey(_key);
            cs.contractWhitelist[cIndex].contractAddress = _contract;
            cs.contractWhitelist[cIndex].contractKey = _key;
            cs.contractWhitelist[cIndex].whitelistStatus = 1;
        }
        else {
            cs.numContracts++;
            cs.contractIndexByKey[_key] = cs.numContracts;
            cs.contractIndexByAddress[_contract] = cs.numContracts;
            cs.contractWhitelist.push(ContractWhitelistDetails(_contract, _key, 1));
        }
        emit ContractWhitelistModified(_contract, _key, 1);
    }

    /** @notice function to revoke whitelisted contract by address
      * @param _contract - contract address
      */
    function revokeWhitelistByAddress(address _contract) external
    onlyImplementation {
        // Check if contract already exists
        require((contractWhitelistStorage().contractIndexByAddress[_contract]) != 0, "whitelist does not exists");
        uint256 cIndex = _getContractIndexByAddress(_contract);
        contractWhitelistStorage().contractWhitelist[cIndex].whitelistStatus = 2;
        emit ContractWhitelistRevoked(_contract, contractWhitelistStorage().contractWhitelist[cIndex].contractKey, 2);
    }

    /** @notice function to revoke whitelisted contract by contract key
      * @param _key - contract key
      */
    function revokeWhitelistByKey(string calldata _key) external
    onlyImplementation {
        // Check if contract already exists
        require((contractWhitelistStorage().contractIndexByKey[_key]) != 0, "whitelist does not exists");
        uint256 cIndex = _getContractIndexByKey(_key);
        contractWhitelistStorage().contractWhitelist[cIndex].whitelistStatus = 2;
        emit ContractWhitelistRevoked(contractWhitelistStorage().contractWhitelist[cIndex].contractAddress, _key, 2);
    }

    /** @notice returns the index for a given contract address
      * @param _contract contract address
      * @return contract index
      */
    function _getContractIndexByAddress(address _contract) internal view returns (uint256) {
        return contractWhitelistStorage().contractIndexByAddress[_contract] - 1;
    }

    /** @notice returns the index for a given contract key
      * @param _key contract address
      * @return contract index
      */
    function _getContractIndexByKey(string calldata _key) internal view returns (uint256) {
        return contractWhitelistStorage().contractIndexByKey[_key] - 1;
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
