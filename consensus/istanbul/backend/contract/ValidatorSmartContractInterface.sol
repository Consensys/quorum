// Interface for contracts used to select validators

pragma solidity >=0.5.0;

interface ValidatorSmartContractInterface {
    function getValidators() external view returns (address[] memory);
}