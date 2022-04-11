// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package ethapi

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type ProxyAPISupport interface {
	ProxyEnabled() bool
	ProxyClient() *rpc.Client
}

// PublicTransactionPoolAPI exposes methods for the RPC interface
type PublicTransactionPoolProxyAPI struct {
	PublicTransactionPoolAPI
	proxyClient *rpc.Client
}

// NewPublicTransactionPoolAPI creates a new RPC service with methods specific for the transaction pool.
func NewPublicTransactionPoolProxyAPI(b Backend, nonceLock *AddrLocker) interface{} {
	apiSupport, ok := b.(ProxyAPISupport)
	if ok && apiSupport.ProxyEnabled() {
		signer := types.LatestSigner(b.ChainConfig())
		return &PublicTransactionPoolProxyAPI{
			PublicTransactionPoolAPI{b, nonceLock, signer},
			apiSupport.ProxyClient(),
		}
	}
	return NewPublicTransactionPoolAPI(b, nonceLock)
}

func (s *PublicTransactionPoolProxyAPI) SendTransaction(ctx context.Context, args SendTxArgs) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "eth_sendTransaction", args)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) SendRawTransaction(ctx context.Context, encodedTx hexutil.Bytes) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "eth_sendRawTransaction", encodedTx)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) SendRawPrivateTransaction(ctx context.Context, encodedTx hexutil.Bytes, args SendRawTxArgs) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "eth_sendRawPrivateTransaction", encodedTx, args)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) FillTransaction(ctx context.Context, args SendTxArgs) (*SignTransactionResult, error) {
	log.Info("QLight - proxy enabled")
	var result SignTransactionResult
	err := s.proxyClient.CallContext(ctx, &result, "eth_fillTransaction", args)
	return &result, err
}

func (s *PublicTransactionPoolProxyAPI) DistributePrivateTransaction(ctx context.Context, encodedTx hexutil.Bytes, args SendRawTxArgs) (string, error) {
	log.Info("QLight - proxy enabled")
	var result string
	err := s.proxyClient.CallContext(ctx, &result, "eth_distributePrivateTransaction", encodedTx, args)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) Resend(ctx context.Context, sendArgs SendTxArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "eth_resend", sendArgs, gasPrice, gasLimit)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) SendTransactionAsync(ctx context.Context, args AsyncSendTxArgs) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "eth_sendTransactionAsync", args)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) Sign(addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	log.Info("QLight - proxy enabled")
	var result hexutil.Bytes
	err := s.proxyClient.Call(&result, "eth_sign", addr, data)
	return result, err
}

func (s *PublicTransactionPoolProxyAPI) SignTransaction(ctx context.Context, args SendTxArgs) (*SignTransactionResult, error) {
	log.Info("QLight - proxy enabled")
	var result SignTransactionResult
	err := s.proxyClient.CallContext(ctx, &result, "eth_signTransaction", args)
	return &result, err
}

type PrivateAccountProxyAPI struct {
	PrivateAccountAPI
	proxyClient *rpc.Client
}

func NewPrivateAccountProxyAPI(b Backend, nonceLock *AddrLocker) interface{} {
	apiSupport, ok := b.(ProxyAPISupport)
	if ok && apiSupport.ProxyEnabled() {
		return &PrivateAccountProxyAPI{
			PrivateAccountAPI{
				am:        b.AccountManager(),
				nonceLock: nonceLock,
				b:         b,
			},
			apiSupport.ProxyClient(),
		}
	}
	return NewPrivateAccountAPI(b, nonceLock)
}

func (s *PrivateAccountProxyAPI) SendTransaction(ctx context.Context, args SendTxArgs, passwd string) (common.Hash, error) {
	log.Info("QLight - proxy enabled")
	var result common.Hash
	err := s.proxyClient.CallContext(ctx, &result, "personal_sendTransaction", args, passwd)
	return result, err
}

func (s *PrivateAccountProxyAPI) SignTransaction(ctx context.Context, args SendTxArgs, passwd string) (*SignTransactionResult, error) {
	log.Info("QLight - proxy enabled")
	var result SignTransactionResult
	err := s.proxyClient.CallContext(ctx, &result, "personal_signTransaction", args, passwd)
	return &result, err
}

func (s *PrivateAccountProxyAPI) Sign(ctx context.Context, data hexutil.Bytes, addr common.Address, passwd string) (hexutil.Bytes, error) {
	log.Info("QLight - proxy enabled")
	var result hexutil.Bytes
	err := s.proxyClient.CallContext(ctx, &result, "personal_sign", data, addr, passwd)
	return result, err
}

func (s *PrivateAccountProxyAPI) SignAndSendTransaction(ctx context.Context, args SendTxArgs, passwd string) (common.Hash, error) {
	return s.SendTransaction(ctx, args, passwd)
}
