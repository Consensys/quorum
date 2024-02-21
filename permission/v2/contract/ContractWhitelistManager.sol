
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
    PermissionsUpgradable private permUpgradable;

    struct ContractWhitelistDetails {
        address contractAddress;
        string contractKey;
        uint8 whitelistStatus;
    }

    ContractWhitelistDetails[] private contractWhitelist;
    mapping(address => uint) private contractIndexByAddress;
    mapping(string => uint) private contractIndexByKey;
    uint private numContracts;

    // contract whitelist events
    event ContractWhitelistModified(address _contractAddr, string _contractKey, uint8 _status);
    event ContractWhitelistRevoked(address _contractAddr, string _contractKey, uint8 _status);

    /** @notice confirms that the caller is the address of implementation
        contract
      */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /// @notice initialized only once. sets the permissions upgradable address
    function initialize(address _permUpgradable) public initializer {
        require(_permUpgradable != address(0x0), "Cannot set to empty address");
        permUpgradable = PermissionsUpgradable(_permUpgradable);
    }

    /** @notice returns the total number of whitelisted contracts
      * @return total number whitelisted contracts
      */
    function getNumberOfWhitelistedContracts() external view returns (uint) {
        return contractWhitelist.length;
    }

    /** @notice returns the contract whitelist details a given contract whitelist index
      * @param  _cIndex contract index
      * @return contract contract address
      * @return key contract key
      * @return status whitelist status
      */
    function getContractWhitelistDetailsFromIndex(uint _cIndex) external view returns
    (address, string memory, uint status) {
        return (contractWhitelist[_cIndex].contractAddress, contractWhitelist[_cIndex].contractKey, contractWhitelist[_cIndex].whitelistStatus);
    }

    /** @notice function to add/update whitelisted contract
      * @param _key - contract key, i.e. human parsable string representation of an unique identifier for a contract
      * @param _contract - contract address
      */
    function addWhitelist(string calldata _key, address _contract) external
    onlyImplementation {
        // Check if contract already exists
        if (contractIndexByKey[_key] != 0) {
            uint256 cIndex = _getContractIndexByKey(_key);
            contractWhitelist[cIndex].contractAddress = _contract;
            contractWhitelist[cIndex].contractKey = _key;
            contractWhitelist[cIndex].whitelistStatus = 1;
        }
        else {
            numContracts++;
            contractIndexByKey[_key] = numContracts;
            contractIndexByAddress[_contract] = numContracts;
            contractWhitelist.push(ContractWhitelistDetails(_contract, _key, 1));
        }
        emit ContractWhitelistModified(_contract, _key, 1);
    }

    /** @notice function to revoke whitelisted contract by address
      * @param _contract - contract address
      */
    function revokeWhitelistByAddress(address _contract) external
    onlyImplementation {
        // Check if contract already exists
        require((contractIndexByAddress[_contract]) != 0, "whitelist does not exists");
        uint256 cIndex = _getContractIndexByAddress(_contract);
        contractWhitelist[cIndex].whitelistStatus = 2;
        emit ContractWhitelistRevoked(_contract, contractWhitelist[cIndex].contractKey, 2);
    }

    /** @notice function to revoke whitelisted contract by contract key
      * @param _key - contract key
      */
    function revokeWhitelistByKey(string calldata _key) external
    onlyImplementation {
        // Check if contract already exists
        require((contractIndexByKey[_key]) != 0, "whitelist does not exists");
        uint256 cIndex = _getContractIndexByKey(_key);
        contractWhitelist[cIndex].whitelistStatus = 2;
        emit ContractWhitelistRevoked(contractWhitelist[cIndex].contractAddress, _key, 2);
    }

    /** @notice returns the index for a given contract address
      * @param _contract contract address
      * @return contract index
      */
    function _getContractIndexByAddress(address _contract) internal view returns (uint256) {
        return contractIndexByAddress[_contract] - 1;
    }

    /** @notice returns the index for a given contract key
      * @param _key contract address
      * @return contract index
      */
    function _getContractIndexByKey(string calldata _key) internal view returns (uint256) {
        return contractIndexByKey[_key] - 1;
    }
}
