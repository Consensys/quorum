package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

type PrivateStateRepositoryWithSetPSRoot interface {
	mps.PrivateStateRepository
	SetPSRoot(psi types.PrivateStateIdentifier, hash common.Hash) error
}

/*
for block in blockChain {
    privateState = stateAt(block.Number) // existing private state
    privateStateRepository = privateStateRepositoryAt(block.Number-1)
    emptyState = privateStateRepository.get(“empty”)
    privateStateRepository.setPrivateState(<privateStateID>, privateState)
    receipts = loadReceipts(block)
    for tx in block {
        if tx.isPrivate() && tx.to == nil {
            receipt = receipts[tx.index]
            addr = receipt.contractAddress
            emptyContract = createEmptyContract(addr)
            emptyState.addContract(emptyContract)
            // build the empty receipt
            emptyReceipt = createEmptyReceipt(receipt)
            emptyReceipt.Versions[“private”] = receipt
            receipts[tx.index] = emptyReceipt
        }
    }

    writeReceipts(receipts, block)
    privateStateRepository.commitAndWrite(block.root)
}
setIsMPS(chainConfig,true)
*/
func UpgradeDB(db ethdb.Database, chain *BlockChain) error {

	currentBlockNumber := chain.CurrentBlock().Number().Int64()

	mpsMgr, err := newMultiplePrivateStateManager(db, nil, nil)
	if err != nil {
		return err
	}
	genesisHeader := chain.GetHeaderByNumber(0)
	mpsRepoIntf, err := mpsMgr.StateRepository(genesisHeader.Root)
	if err != nil {
		return err
	}
	var mpsRepo, ok = mpsRepoIntf.(PrivateStateRepositoryWithSetPSRoot)
	if !ok {
		return fmt.Errorf("Invalid mps repo")
	}

	for idx := 1; idx <= int(currentBlockNumber); idx++ {
		header := chain.GetHeaderByNumber(uint64(idx))
		// TODO consider periodic reports instead of logging about each block
		fmt.Printf("Processing block %v with hash %v\n", idx, header.Hash().Hex())
		block := chain.GetBlock(header.Hash(), header.Number.Uint64())
		emptyState, err := mpsRepo.DefaultState()
		if err != nil {
			return err
		}
		existingPrivateStateRoot := rawdb.GetPrivateStateRoot(db, header.Root)
		err = mpsRepo.SetPSRoot(types.DefaultPrivateStateIdentifier, existingPrivateStateRoot)
		if err != nil {
			return err
		}
		receipts := chain.GetReceiptsByHash(header.Hash())
		receiptsUpdated := false
		for txIdx, tx := range block.Transactions() {
			if tx.IsPrivate() && tx.To() == nil {
				// this is a contract creation transaction
				receipt := receipts[txIdx]
				accountAddress := receipt.ContractAddress
				emptyState.CreateAccount(accountAddress)
				emptyState.SetNonce(accountAddress, 1)

				emptyReceipt := &types.Receipt{
					PostState:         receipt.PostState,
					Status:            1,
					CumulativeGasUsed: receipt.CumulativeGasUsed,
					Bloom:             types.Bloom{},
					Logs:              nil,
					TxHash:            receipt.TxHash,
					ContractAddress:   receipt.ContractAddress,
					GasUsed:           receipt.GasUsed,
					BlockHash:         receipt.BlockHash,
					BlockNumber:       receipt.BlockNumber,
					TransactionIndex:  receipt.TransactionIndex,
				}
				emptyReceipt.Bloom = types.CreateBloom(types.Receipts{emptyReceipt})
				emptyReceipt.PSReceipts = map[types.PrivateStateIdentifier]*types.Receipt{types.DefaultPrivateStateIdentifier: receipt}
				receipts[txIdx] = emptyReceipt
				receiptsUpdated = true
			}
		}

		err = mpsRepo.CommitAndWrite(chain.chainConfig.IsEIP158(block.Number()), block)
		if err != nil {
			return err
		}

		if receiptsUpdated {
			batch := db.NewBatch()
			rawdb.WriteReceipts(batch, block.Hash(), block.NumberU64(), receipts)
			err := batch.Write()
			if err != nil {
				return err
			}
		}
	}
	// update isMPS in the chain config
	config := chain.Config()
	config.IsMPS = true
	rawdb.WriteChainConfig(db, rawdb.ReadCanonicalHash(db, 0), config)
	fmt.Printf("MPS DB upgrade finished successfully.\n")
	return nil
}
