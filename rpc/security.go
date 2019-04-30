package rpc

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

//ProcessRequestSecurity RPC WS/WSS/HTTPS/HTTP request security.
// Deny all policy by default
func (ctx *SecurityContext) ProcessRequestSecurity(request rpcRequest) error {
	// Ensure Deny By Default Policy
	if ctx.Enabled && ctx.Config == nil {
		return errors.New("Request requires valid token")
	}
	// Check if token/scope using provider.
	if !ctx.Provider.IsClientAuthorized(request) {
		return errors.New("Request requires valid token")
	}

	return nil
}

// return true if is local security provider
func (ctx *SecurityContext) IsLocalSecurityProviderAvailable() (bool, error) {
	if ctx.Provider == nil {
		return false, fmt.Errorf("security provider not set")
	}

	if strings.ToLower(ctx.Provider.GetType()) == LocalSecProvider {
			return true , nil
		}else {
			return false, nil
		}
}


// GetDenyAllPolicy returns a disabled context
func GetDenyAllPolicy() SecurityContext {
	return SecurityContext{Enabled: true}
}

// GetDefaultAllowAllPolicy returns a disabled context
func GetDefaultAllowAllPolicy() SecurityContext {
	return SecurityContext{Enabled: false}
}
