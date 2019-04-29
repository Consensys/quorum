package rpc

import ("github.com/pkg/errors")

//ProcessRequestSecurity RPC WS/WSS/HTTPS/HTTP request security.
// Deny all policy by default
func (ctx *SecurityContext) ProcessRequestSecurity(request rpcRequest) error {
	// Ensure Deny By Default Policy
	if ctx.Enabled && ctx.Config == nil {
		return errors.New("Request requires valid token")
	}
	// Check if token/scope using provider.
	if !ctx.Provider.isClientAuthorized(request) {
		return errors.New("Request requires valid token")
	}

	return nil
}

// GetDenyAllPolicy returns a disabled context
func GetDenyAllPolicy() SecurityContext {
	return SecurityContext{Enabled: true}
}

// GetDefaultAllowAllPolicy returns a disabled context
func GetDefaultAllowAllPolicy() SecurityContext {
	return SecurityContext{Enabled: false}
}
