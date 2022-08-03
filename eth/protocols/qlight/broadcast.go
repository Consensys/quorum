package qlight

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/qlight"
)

// blockPropagation is a block propagation event, waiting for its turn in the
// broadcast queue.
type blockPropagation struct {
	block            *types.Block
	td               *big.Int
	blockPrivateData *qlight.BlockPrivateData
}

// broadcastBlocks is a write loop that multiplexes blocks and block accouncements
// to the remote peer. The goal is to have an async writer that does not lock up
// node internals and at the same time rate limits queued data.
func (p *Peer) broadcastBlocksQLightServer() {
	for {
		select {
		case prop := <-p.queuedBlocks:
			if prop.blockPrivateData != nil {
				if prop.blockPrivateData.PSI.String() == p.qlightPSI {
					if err := p.SendBlockPrivateData([]qlight.BlockPrivateData{*prop.blockPrivateData}); err != nil {
						p.Log().Error("Error occurred while sending private data msg", "err", err)
						return
					}
				} else {
					p.Log().Error("PSI mismatch for block private data", "bpdPSI", prop.blockPrivateData.PSI, "peerPSI", p.qlightPSI)
				}
			}
			if err := p.SendNewBlock(prop.block, prop.td); err != nil {
				return
			}
			p.Log().Trace("Propagated block", "number", prop.block.Number(), "hash", prop.block.Hash(), "td", prop.td)
		case <-p.term:
			return
		}
	}
}
