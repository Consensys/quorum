package plugin

import (
	"errors"

	"github.com/hashicorp/go-plugin"
)

const (
	DefaultProtocolVersion = 1
)

var (
	DefaultHandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  DefaultProtocolVersion,
		MagicCookieKey:   "QUORUM_PLUGIN_MAGIC_COOKIE",
		MagicCookieValue: "CB9F51969613126D93468868990F77A8470EB9177503C5A38D437FEFF7786E0941152E05C06A9A3313391059132A7F9CED86C0783FE63A8B38F01623C8257664",
	}

	ErrNotSupported = errors.New("not supported")
)
