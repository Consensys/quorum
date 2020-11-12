package security

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type TLSConfigurationSource interface {
	Get(ctx context.Context) (*tls.Config, error)
}

type AuthenticationManager interface {
	Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error)
	IsEnabled(ctx context.Context) (bool, error)
}

type AuthenticationManagerDeferFunc func() (AuthenticationManager, error)

type DeferredAuthenticationManager struct {
	deferFunc AuthenticationManagerDeferFunc
}

func (d *DeferredAuthenticationManager) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	am, err := d.deferFunc()
	if err != nil {
		return nil, err
	}
	return am.Authenticate(ctx, token)
}

func (d *DeferredAuthenticationManager) IsEnabled(ctx context.Context) (bool, error) {
	am, err := d.deferFunc()
	if err != nil {
		return false, err
	}
	return am.IsEnabled(ctx)
}

func NewDeferredAuthenticationManager(deferFunc AuthenticationManagerDeferFunc) *DeferredAuthenticationManager {
	return &DeferredAuthenticationManager{
		deferFunc: deferFunc,
	}
}

type DisabledAuthenticationManager struct {
}

func (*DisabledAuthenticationManager) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	return nil, errors.New("not supported operation")
}

func (*DisabledAuthenticationManager) IsEnabled(ctx context.Context) (bool, error) {
	return false, nil
}

func NewDisabledAuthenticationManager() AuthenticationManager {
	return &DisabledAuthenticationManager{}
}
