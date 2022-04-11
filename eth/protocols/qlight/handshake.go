package qlight

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/p2p"
)

const (
	// handshakeTimeout is the maximum allowed time for the `eth` handshake to
	// complete before dropping the connection.= as malicious.
	handshakeTimeout = 5 * time.Second
)

func (p *Peer) QLightHandshake(server bool, psi string, token string) error {
	// Send out own handshake in a new thread
	errc := make(chan error, 2)

	var (
		status qLightStatusData // safe to read after two values have been received from errc
	)
	go func() {
		errc <- p2p.Send(p.rw, QLightStatusMsg, &qLightStatusData{
			ProtocolVersion: uint32(p.version),
			Server:          server,
			PSI:             psi,
			Token:           token,
		})
	}()
	go func() {
		errc <- p.readQLightStatus(&status)
	}()
	timeout := time.NewTimer(handshakeTimeout)
	defer timeout.Stop()
	for i := 0; i < 2; i++ {
		select {
		case err := <-errc:
			if err != nil {
				return err
			}
		case <-timeout.C:
			return p2p.DiscReadTimeout
		}
	}
	p.qlightServer, p.qlightPSI, p.qlightToken = status.Server, status.PSI, status.Token
	return nil
}

func (p *Peer) readQLightStatus(qligtStatus *qLightStatusData) error {
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Code != QLightStatusMsg {
		return fmt.Errorf("%w: second msg has code %x (!= %x)", errNoStatusMsg, msg.Code, QLightStatusMsg)
	}
	if msg.Size > maxMessageSize {
		return fmt.Errorf("Message too long: %v > %v", msg.Size, maxMessageSize)
	}
	// Decode the handshake and make sure everything matches
	if err := msg.Decode(&qligtStatus); err != nil {
		return fmt.Errorf("%w: msg %v: %v", errDecode, msg, err)
	}
	if !qligtStatus.Server && len(qligtStatus.PSI) == 0 {
		return fmt.Errorf("client connected without specifying PSI")
	}
	return nil
}
