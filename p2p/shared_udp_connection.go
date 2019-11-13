package p2p

import (
	"errors"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"net"
)

// UDPReader implements a shared connection. Write sends messages to the
// underlying connection while read returns messages that were found
// unprocessable and sent to the unhandled channel by the primary listener.
type UDPReader interface {
	ReadFromUDP([]byte) (int, *net.UDPAddr, error)

	WriteToUDP(b []byte, addr *net.UDPAddr) (n int, err error)
	Close() error
	LocalAddr() net.Addr
}

// baseSharedUDPConn implements a shared connection. Read/Write reads/sends
// messages to the underlying connection whilst messages that were found
// unprocessable are sent to the unhandled channel to be processed later.
type baseSharedUDPConn struct {
	*net.UDPConn
	unhandled chan discover.ReadPacket
}

// sharedUDPConn implements a shared connection. Write sends messages to the
// underlying connection while read returns messages that were found
// unprocessable and sent to the unhandled channel by the primary listener.
type sharedUDPConn struct {
	parent UDPReader
	unhandled chan discover.ReadPacket
}

func NewBaseSharedUDPConn(udpConn *net.UDPConn, unhandled chan discover.ReadPacket) *baseSharedUDPConn {
	return &baseSharedUDPConn{udpConn, unhandled,}
}

func NewSharedUDPConn(udpConn UDPReader, unhandled chan discover.ReadPacket) *sharedUDPConn {
	return &sharedUDPConn{udpConn, unhandled,}
}

func (s *baseSharedUDPConn) ReadFromUDP(b []byte) (n int, addr *net.UDPAddr, err error) {
	packet, ok := <-s.unhandled
	if !ok {
		return 0, nil, errors.New("connection was closed")
	}
	l := len(packet.Data)
	if l > len(b) {
		l = len(b)
	}
	copy(b[:l], packet.Data[:l])
	return l, packet.Addr, nil
}

func (s *baseSharedUDPConn) Close() error {
	return nil
}

func (s *sharedUDPConn) ReadFromUDP(b []byte) (n int, addr *net.UDPAddr, err error) {
	packet, ok := <-s.unhandled
	if !ok {
		return 0, nil, errors.New("connection was closed")
	}
	l := len(packet.Data)
	if l > len(b) {
		l = len(b)
	}
	copy(b[:l], packet.Data[:l])
	return l, packet.Addr, nil
}

func (s *sharedUDPConn) Close() error {
	return nil
}

func (s *sharedUDPConn) WriteToUDP(b []byte, addr *net.UDPAddr) (n int, err error) {
	return s.parent.WriteToUDP(b, addr)
}

func (s *sharedUDPConn) LocalAddr() net.Addr {
	return s.parent.LocalAddr()
}