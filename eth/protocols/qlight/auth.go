package qlight

import (
	"time"

	"github.com/ethereum/go-ethereum/p2p"
)

func (p *Peer) PeriodicAuthCheck() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			p.Log().Debug("Performing periodic auth check")
			err := p.QLightPeriodicAuthFunc()
			if err != nil {
				p.Log().Error("Disconnecting peer due to periodic auth check", "err", err)
				p.Disconnect(p2p.DiscAuthError)
			}
		case <-p.term:
			return
		}
	}
}
