// Quorum
package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type securityContextKey string
type SecurityContext context.Context

const (
	HttpAuthorizationHeader              = "Authorization"
	HttpPrivateStateIdentifierHeader     = "Quorum-PSI"
	QueryPrivateStateIdentifierParamName = "PSI"
	EnvVarPrivateStateIdentifier         = "QUORUM_PSI"
	// this key is set by server to indicate if server supports mulitenancy
	ctxIsMultitenant = securityContextKey("IS_MULTITENANT")
	// this key is set into the secured context to indicate
	// the authorized private state being operated on for the request.
	// the value MUST BE OF TYPE types.PrivateStateIdentifier
	ctxPrivateStateIdentifier = securityContextKey("PRIVATE_STATE_IDENTIFIER")
	// this key is set into the request context to indicate
	// the private state being operated on for the request
	ctxRequestPrivateStateIdentifier = securityContextKey("REQUEST_PRIVATE_STATE_IDENTIFIER")
	// this key is exported for WS transport
	ctxCredentialsProvider = securityContextKey("CREDENTIALS_PROVIDER") // key to save reference to rpc.HttpCredentialsProviderFunc
	ctxPSIProvider         = securityContextKey("PSI_PROVIDER")         // key to save reference to rpc.PSIProviderFunc
	// keys used to save values in request context
	ctxAuthenticationError   = securityContextKey("AUTHENTICATION_ERROR")   // key to save error during authentication before processing the request body
	ctxPreauthenticatedToken = securityContextKey("PREAUTHENTICATED_TOKEN") // key to save the preauthenticated token once authenticated
)

// WithIsMultitenant populates ctx with ctxIsMultitenant key and provided value
func WithIsMultitenant(ctx context.Context, isMultitenant bool) SecurityContext {
	return context.WithValue(ctx, ctxIsMultitenant, isMultitenant)
}

// IsMultitenantFromContext returns bool value from ctx with ctxIsMultitenant key
// and returns false if value does not exist in the ctx
func IsMultitenantFromContext(ctx SecurityContext) bool {
	if f, ok := ctx.Value(ctxIsMultitenant).(bool); ok {
		return f
	}
	return false
}

// WithPrivateStateIdentifier populates ctx with ctxPrivateStateIdentifier key and provided value
func WithPrivateStateIdentifier(ctx context.Context, psi types.PrivateStateIdentifier) SecurityContext {
	return context.WithValue(ctx, ctxPrivateStateIdentifier, psi)
}

// PrivateStateIdentifierFromContext returns types.PrivateStateIdentifier value from ctx with ctxPrivateStateIdentifier key
func PrivateStateIdentifierFromContext(ctx SecurityContext) (types.PrivateStateIdentifier, bool) {
	psi, found := ctx.Value(ctxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	return psi, found
}

// WithCredentialsProvider populates ctx with ctxCredentialsProvider key and provided value
func WithCredentialsProvider(ctx context.Context, f HttpCredentialsProviderFunc) SecurityContext {
	return context.WithValue(ctx, ctxCredentialsProvider, f)
}

// CredentialsProviderFromContext returns HttpCredentialsProviderFunc value from ctx with ctxCredentialsProvider key
// and returns nil if value does not exist in the ctx
func CredentialsProviderFromContext(ctx SecurityContext) HttpCredentialsProviderFunc {
	if f, ok := ctx.Value(ctxCredentialsProvider).(HttpCredentialsProviderFunc); ok {
		return f
	}
	return nil
}

// WithPSIProvider populates ctx with ctxPSIProvider key and provided value
func WithPSIProvider(ctx context.Context, f PSIProviderFunc) SecurityContext {
	return context.WithValue(ctx, ctxPSIProvider, f)
}

// PSIProviderFromContext returns PSIProviderFunc value from ctx with ctxPSIProvider key
// and returns nil if value does not exist in the ctx
func PSIProviderFromContext(ctx SecurityContext) PSIProviderFunc {
	if f, ok := ctx.Value(ctxPSIProvider).(PSIProviderFunc); ok {
		return f
	}
	return nil
}

// WithPreauthenticatedToken populates ctx with ctxPreauthenticatedToken key and provided value
func WithPreauthenticatedToken(ctx context.Context, token *proto.PreAuthenticatedAuthenticationToken) SecurityContext {
	return context.WithValue(ctx, ctxPreauthenticatedToken, token)
}

// PreauthenticatedTokenFromContext returns *proto.PreAuthenticatedAuthenticationToken value from ctx with ctxPreauthenticatedToken key
// and returns nil if value does not exist in the ctx
func PreauthenticatedTokenFromContext(ctx SecurityContext) *proto.PreAuthenticatedAuthenticationToken {
	if t, ok := ctx.Value(ctxPreauthenticatedToken).(*proto.PreAuthenticatedAuthenticationToken); ok {
		return t
	}
	return nil
}
