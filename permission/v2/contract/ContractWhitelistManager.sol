
pragma solidity ^0.8.4;

import "./PermissionsUpgradable.sol";

/** @title Contract whitelist manager contract
  * @notice This contract holds implementation logic for all contract whitelisting
    functionality. This can be called only by the implementation contract only.
    there are few view functions exposed as public and can be called directly.
    these are invoked by quorum for populating permissions data in cache
  */
contract ContractWhitelistManager {
    PermissionsUpgradable private permUpgradable;
    struct ContractWhitelistDetails {
        address contractAddress;
    }

    ContractWhitelistDetails[] private contractWhitelist;
    mapping(address => uint) private contractIndex;
    uint private numContracts;

    // contract whitelist events
    event ContractWhitelistModified(address _contract);

    /** @notice confirms that the caller is the address of implementation
        contract
      */
    modifier onlyImplementation {
        require(msg.sender == permUpgradable.getPermImpl(), "invalid caller");
        _;
    }

    /// @notice constructor. sets the permissions upgradable address
    constructor (address _permUpgradable) public {
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
      */
    function getContractWhitelistDetailsFromIndex(uint _cIndex) external view returns
    (address) {
        return contractWhitelist[_cIndex].contractAddress;
    }


    /** @notice function to add a new whitelisted contract
      * @param _contract - contract address
      */
    function addNewContract(address _contract) external
    onlyImplementation {
        contractIndex[_contract] = numContracts;
        emit ContractWhitelistModified(_contract);
        // Check if contract already exists
        uint256 cIndex = _getContractIndex(_contract);
        if (contractIndex[_contract] != 0) {
            contractWhitelist[cIndex].contractAddress = _contract;
        }
        else {
            numContracts ++;
            contractIndex[_contract] = numContracts;
            contractWhitelist.push(ContractWhitelistDetails(_contract));
        }
        emit ContractWhitelistModified(_contract);
    }

    /** @notice returns the index for a given contract address
      * @param _contract contract address
      * @return contract index
      */
    function _getContractIndex(address _contract) internal view returns (uint256) {
        return contractIndex[_contract] - 1;
    }
}
