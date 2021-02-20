package message

import "github.com/ethereum/go-ethereum/rlp"

func DecodeMessage(code uint64, data []byte) (QBFTMessage, error){
	switch code {
	case preprepareMsgCode:
		var preprepare PreprepareMsg
		if err := rlp.DecodeBytes(data, &preprepare); err != nil {
			return nil, ErrFailedDecodePreprepare
		}
		preprepare.code = preprepareMsgCode
		return &preprepare, nil
	case prepareMsgCode:
		var prepare Prepare
		if err := rlp.DecodeBytes(data, &prepare); err != nil {
			return nil, ErrFailedDecodeCommit
		}
		prepare.code = prepareMsgCode
		return &prepare, nil
	case commitMsgCode:
		var commit Commit
		if err := rlp.DecodeBytes(data, &commit); err != nil {
			return nil, ErrFailedDecodeCommit
		}
		commit.code = commitMsgCode
		return &commit, nil
	case roundChangeMsgCode:
		var roundChange RoundChangeMsg
		if err := rlp.DecodeBytes(data, &roundChange); err != nil {
			return nil, ErrFailedDecodeRoundChange
		}
		roundChange.code = roundChangeMsgCode
		return &roundChange, nil
	}

	return nil, ErrInvalidMessage
}

