package multitenancy

import (
	"net/url"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(false)))
}

type testCase struct {
	msg          string
	granted      []string
	ask          *PrivateStateSecurityAttribute
	isAuthorized bool
}

func TestMatch_whenTypical(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")
	ask, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")

	assert.True(t, match(ask, granted))
}

func TestMatch_whenNoEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1")
	ask, _ := url.Parse("psi://arbitrary.psi1")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenAskWithNoEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")
	ask, _ := url.Parse("psi://arbitrary.psi1")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenGrantWithNoEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1")
	ask, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenGrantWithDifferentEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")
	ask, _ := url.Parse("psi://arbitrary.psi1?self.eoa=0xaaa")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenAskMultipleEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa")
	ask, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa&node.eoa=0xbbb")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenGrantMultipleEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0x111&self.eoa=0xaaa&self.eoa=0xbbb&self.eoa=0xccc")
	ask, _ := url.Parse("psi://arbitrary.psi1?self.eoa=0xaaa&self.eoa=0xbbb")

	assert.True(t, match(ask, granted))
}

func TestMatch_whenGrantWithWildCardEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0x0")
	ask, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa&node.eoa=0xbbb")

	assert.True(t, match(ask, granted))
}

func TestMatch_whenDiffScheme(t *testing.T) {
	granted, _ := url.Parse("rpc://eth_sendTransaction")
	ask, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0xaaa&node.eoa=0xbbb")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenDiffPSI(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1?node.eoa=0x0")
	ask, _ := url.Parse("psi://arbitrary.psi2?node.eoa=0xaaa&node.eoa=0xbbb")

	assert.False(t, match(ask, granted))
}

func TestMatch_whenDiffPSIAndNoEOA(t *testing.T) {
	granted, _ := url.Parse("psi://arbitrary.psi1")
	ask, _ := url.Parse("psi://arbitrary.psi2")

	assert.False(t, match(ask, granted))
}

func TestAuthorizePSI(t *testing.T) {
	testCases := []struct {
		msg          string
		granted      []string
		ask          types.PrivateStateIdentifier
		isAuthorized bool
	}{
		{
			msg: "Granting PSI with no EOA",
			granted: []string{
				"psi://arbitrary.ps1",
			},
			ask:          "arbitrary.ps1",
			isAuthorized: true,
		},
		{
			msg: "Granting PSI with EOA",
			granted: []string{
				"psi://arbitrary.ps1?node.eoa=0x0",
			},
			ask:          "arbitrary.ps1",
			isAuthorized: true,
		},
		{
			msg: "Different scheme",
			granted: []string{
				"rpc://arbitrary.ps1",
			},
			ask:          "arbitrary.ps1",
			isAuthorized: false,
		},
	}

	for _, tc := range testCases {
		log.Debug("Test case :: " + tc.msg)
		actual, err := IsPSIAuthorized(toToken(tc.granted), tc.ask)
		assert.NoError(t, err, tc.msg)
		assert.Equal(t, tc.isAuthorized, actual, tc.msg)
	}
}

func TestAuthorize(t *testing.T) {
	testCases := []testCase{
		{
			msg: "Granting PSI with no EOA",
			granted: []string{
				"psi://arbitrary.ps1",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1").
				WithNodeEOA(common.HexToAddress("0x000000000000000000000000000000000000aaaa")),
			isAuthorized: false,
		},
		{
			msg: "Granted with default wild card EOA, inadequate ask",
			granted: []string{
				"psi://arbitrary.ps1",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1"),
			isAuthorized: false,
		},
		{
			msg: "Node-managed: Granted with wild card EOA, ask for specific",
			granted: []string{
				"psi://arbitrary.ps1?node.eoa=0x0&self.eoa=0x000000000000000000000000000000000000aaaa",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1").
				WithNodeEOA(common.StringToAddress("0xc")),
			isAuthorized: true,
		},
		{
			msg: "Different EOA grant",
			granted: []string{
				"psi://arbitrary.ps1?self.eoa=0x000000000000000000000000000000000000aaaa",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1").
				WithNodeEOA(common.StringToAddress("0xc")),
			isAuthorized: false,
		},
		{
			msg: "Self-managed: Granted with wild card EOA, ask for specific",
			granted: []string{
				"psi://arbitrary.ps1?self.eoa=0x0",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1").
				WithSelfEOA(common.StringToAddress("0xc")),
			isAuthorized: true,
		},
		{
			msg: "Not granted to a PSI",
			granted: []string{
				"psi://arbitrary.ps2?node.eoa=0x0&self.eoa=0x0",
			},
			ask: (&PrivateStateSecurityAttribute{}).
				WithPSI("arbitrary.ps1").WithNodeEOA(common.StringToAddress("arbitrary")),
			isAuthorized: false,
		},
	}

	for _, tc := range testCases {
		log.Debug("Test case :: " + tc.msg)
		actual, err := IsAuthorized(toToken(tc.granted), tc.ask)
		assert.NoError(t, err, tc.msg)
		assert.Equal(t, tc.isAuthorized, actual, tc.msg)
	}
}

func toToken(granted []string) *proto.PreAuthenticatedAuthenticationToken {
	values := make([]*proto.GrantedAuthority, len(granted))
	for i, g := range granted {
		values[i] = &proto.GrantedAuthority{
			Raw: g,
		}
	}
	return &proto.PreAuthenticatedAuthenticationToken{
		Authorities: values,
	}
}

func TestExtractPSI_whenTypical(t *testing.T) {
	psi, err := ExtractPSI(toToken([]string{
		"psi://arbitrary.psi1",
		"psi://arbitrary.psi1?node.eoa=0x0",
		"rpc://eth_call",
	}))

	assert.NoError(t, err)
	assert.Equal(t, types.ToPrivateStateIdentifier("arbitrary.psi1"), psi)
}

func TestExtractPSI_whenNotFound(t *testing.T) {
	_, err := ExtractPSI(toToken([]string{
		"rpc://eth_call",
	}))

	assert.EqualError(t, err, ErrPSINotFound.Error())
}

func TestExtractPSI_whenFoundMultiple(t *testing.T) {
	_, err := ExtractPSI(toToken([]string{
		"psi://arbitrary.psi1",
		"psi://arbitrary.psi2",
		"psi://arbitrary.psi3",
	}))

	assert.EqualError(t, err, ErrPSIFoundMultiple.Error())
}
