package lc

import "github.com/ethereum/go-ethereum/plugin/gen/proto"

type PluginGateway struct {
	client proto.LCProtocolServiceClient
}
