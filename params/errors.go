package params

import "errors"

var (
	ErrTransitionAlgorithm = errors.New("transition algorithm is invalid, should be either `ibft` or `qbft`")
	ErrBlockNumberMissing  = errors.New("block number not given in transitions data")
	ErrBlockOrder          = errors.New("block order should be ascending")
	ErrTransition          = errors.New("can't transition from qbft to ibft")
)
