package graphql

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
)

func TestAddQuorumHTML_whenTypical(t *testing.T) {
	testObject := GraphiQL{
		authManagerFunc: func() (security.AuthenticationManager, error) {
			return security.NewDisabledAuthenticationManager(), nil
		},
		isMPS: false,
	}
	out, err := testObject.addQuorumHTML(graphiql)

	assert.NoError(t, err)
	assert.NotContains(t, string(out), "access-token")
	assert.NotContains(t, string(out), "psi")
}

func TestAddQuorumHTML_whenMPS(t *testing.T) {
	testObject := GraphiQL{
		authManagerFunc: func() (security.AuthenticationManager, error) {
			return security.NewDisabledAuthenticationManager(), nil
		},
		isMPS: true,
	}
	out, err := testObject.addQuorumHTML(graphiql)

	assert.NoError(t, err)
	html := string(out)
	assert.NotContains(t, html, "access-token")
	assert.Contains(t, html, "psi")
}

func TestAddQuorumHTML_whenRPCSecured(t *testing.T) {
	testObject := GraphiQL{
		authManagerFunc: func() (security.AuthenticationManager, error) {
			return &StubAuthenticationManager{}, nil
		},
		isMPS: false,
	}
	out, err := testObject.addQuorumHTML(graphiql)

	assert.NoError(t, err)
	html := string(out)
	assert.Contains(t, html, "access-token")
	assert.NotContains(t, html, "psi")
}

func TestAddQuorumHTML_whenMPSAndRPCSecured(t *testing.T) {
	testObject := GraphiQL{
		authManagerFunc: func() (security.AuthenticationManager, error) {
			return &StubAuthenticationManager{}, nil
		},
		isMPS: true,
	}
	out, err := testObject.addQuorumHTML(graphiql)

	assert.NoError(t, err)
	html := string(out)
	assert.Contains(t, html, "access-token")
	assert.Contains(t, html, "psi")
}

type StubAuthenticationManager struct {
}

func (s *StubAuthenticationManager) Authenticate(_ context.Context, _ string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	panic("implement me")
}

func (s *StubAuthenticationManager) IsEnabled(_ context.Context) (bool, error) {
	return true, nil
}
