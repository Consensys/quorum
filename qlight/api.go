package qlight

import "github.com/ethereum/go-ethereum/rpc"

type RunningPeerAuthUpdater interface {
	UpdateTokenForRunningQPeers(token string) error
}

type PrivateQLightAPI struct {
	tokenHolder *TokenHolder
	peerUpdater RunningPeerAuthUpdater
	rpcClient   *rpc.Client
}

// NewPublicEthereumAPI creates a new Ethereum protocol API for full nodes.
func NewPrivateQLightAPI(peerUpdater RunningPeerAuthUpdater, rpcClient *rpc.Client) *PrivateQLightAPI {
	return &PrivateQLightAPI{peerUpdater: peerUpdater, rpcClient: rpcClient}
}

func (p *PrivateQLightAPI) SetCurrentToken(token string) {
	p.tokenHolder.SetCurrentToken(token)
	p.peerUpdater.UpdateTokenForRunningQPeers(token)
	if p.rpcClient != nil {
		if len(token) > 0 {
			p.rpcClient.WithHTTPCredentials(p.tokenHolder.HttpCredentialsProvider)
		} else {
			p.rpcClient.WithHTTPCredentials(nil)
		}
	}
}

func (p *PrivateQLightAPI) GetCurrentToken() string {
	return p.tokenHolder.CurrentToken()
}

func (p *PrivateQLightAPI) ReloadPlugin() error {
	return p.tokenHolder.ReloadPlugin()
}
