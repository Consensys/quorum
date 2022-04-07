package p2p

import (
	"crypto/ecdsa"
	"crypto/tls"
	"net"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
)

var qlightTLSConfig *tls.Config

func SetQLightTLSConfig(config *tls.Config) {
	qlightTLSConfig = config
}

type tlsErrorTransport struct {
	err error
}

func (tr *tlsErrorTransport) doEncHandshake(prv *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	return nil, tr.err
}
func (tr *tlsErrorTransport) doProtoHandshake(our *protoHandshake) (*protoHandshake, error) {
	return nil, tr.err
}
func (tr *tlsErrorTransport) ReadMsg() (Msg, error) { return Msg{}, tr.err }
func (tr *tlsErrorTransport) WriteMsg(Msg) error    { return tr.err }
func (tr *tlsErrorTransport) close(err error)       {}

func NewQlightClientTransport(conn net.Conn, dialDest *ecdsa.PublicKey) transport {
	log.Info("Setting up qlight client transport")
	if qlightTLSConfig != nil {
		tlsConn := tls.Client(conn, qlightTLSConfig)
		err := tlsConn.Handshake()
		if err != nil {
			log.Error("Failure setting up qlight client transport", "err", err)
			return &tlsErrorTransport{err}
		}
		log.Info("Qlight client tls transport established successfully")
		return &rlpxTransport{conn: rlpx.NewConn(tlsConn, dialDest)}
	}
	return &rlpxTransport{conn: rlpx.NewConn(conn, dialDest)}
}

func NewQlightServerTransport(conn net.Conn, dialDest *ecdsa.PublicKey) transport {
	log.Info("Setting up qlight server transport")
	if qlightTLSConfig != nil {
		tlsConn := tls.Server(conn, qlightTLSConfig)
		err := tlsConn.Handshake()
		if err != nil {
			log.Error("Failure setting up qlight server transport", "err", err)
			return &tlsErrorTransport{err}
		}
		log.Info("Qlight server tls transport established successfully")
		return &rlpxTransport{conn: rlpx.NewConn(tlsConn, dialDest)}
	}
	return &rlpxTransport{conn: rlpx.NewConn(conn, dialDest)}
}
