package params

import (
	"errors"
	"fmt"
)

var (
	ErrTransitionAlgorithm             = errors.New("transition algorithm is invalid, should be either `ibft` or `qbft`")
	ErrBlockNumberMissing              = errors.New("block number not given in transitions data")
	ErrBlockOrder                      = errors.New("block order should be ascending")
	ErrTransition                      = errors.New("can't transition from qbft to ibft")
	ErrTestQBFTBlockAndTransitions     = errors.New("can't use transition algorithm and testQBFTBlock at the same time")
	ErrMaxCodeSizeConfigAndTransitions = errors.New("can't use transition ContractSizeLimit and MaxCodeSizeConfig at the same time")
	ErrContractSizeLimit               = errors.New("transition contract code size must be between 24 and 128")
	ErrValidatorSelectionMode          = errors.New("validator selection mode is invalid, should be either `contract` or `blockheader`")
	ErrMissingValidatorSelectionMode   = errors.New("validator selection mode is missing, should specify `contract` when using validatorcontractaddress")
	ErrTransactionSizeLimit            = errors.New("genesis transaction size limit must be between 32 and 128")
	ErrBeneficiaryMode                 = errors.New("beneficiary mode is not valid")
)

func ErrTransitionIncompatible(field string) error {
	return fmt.Errorf("transitions.%s data incompatible. %s historical data does not match", field, field)
}
