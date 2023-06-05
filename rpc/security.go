package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/multitenancy"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type securityContextSupport interface {
	securityContextConfigurer
	SecurityContextResolver
}

type securityContextConfigurer interface {
	Configure(secCtx SecurityContext)
}

type SecurityContextResolver interface {
	Resolve() SecurityContext
}

type securityError struct{ message string }

// Provider function to return token being injected in Authorization http request header
type HttpCredentialsProviderFunc func(ctx context.Context) (string, error)

// Provider function to return a string value which will be
// 1. injected in HttpPrivateStateIdentifierHeader http request header for HTTP/WS transports
// 2. encoded in JSON MessageID for IPC/InProc transports
type PSIProviderFunc func(ctx context.Context) (types.PrivateStateIdentifier, error)

func (e *securityError) ErrorCode() int { return -32001 }

func (e *securityError) Error() string { return e.message }

func extractToken(req *http.Request) (string, bool) {
	token := req.Header.Get(HttpAuthorizationHeader)
	return token, token != ""
}

func verifyExpiration(token *proto.PreAuthenticatedAuthenticationToken) error {
	if token == nil {
		return nil
	}
	err := token.ExpiredAt.CheckValid()
	if err != nil {
		return fmt.Errorf("invalid timestamp in token: %s", err)
	}
	if time.Now().Before(token.ExpiredAt.AsTime()) {
		return nil
	}
	return &securityError{"token expired"}
}

func verifyAccess(service, method string, authorities []*proto.GrantedAuthority) error {
	for _, authority := range authorities {
		if authority.Service == "*" && authority.Method == "*" {
			return nil
		}
		if authority.Service == "*" && authority.Method == method {
			return nil
		}
		if authority.Service == service && authority.Method == "*" {
			return nil
		}
		if authority.Service == service && authority.Method == method {
			return nil
		}
	}
	return &securityError{fmt.Sprintf("%s%s%s - access denied", service, serviceMethodSeparator, method)}
}

// verify if a call is authorized using information available in the security context
// it also checks for token expiration. That means if this is called multiple times (batch processing),
// token expiration is checked multiple times.
//
// It returns the verfied security context for caller to use.
func SecureCall(resolver SecurityContextResolver, method string) (context.Context, error) {
	secCtx := resolver.Resolve()
	if secCtx == nil {
		return context.Background(), nil
	}
	if err, hasError := secCtx.Value(ctxAuthenticationError).(error); hasError {
		return nil, err
	}
	if authToken := PreauthenticatedTokenFromContext(secCtx); authToken != nil {
		if err := verifyExpiration(authToken); err != nil {
			return nil, err
		}
		elem := strings.SplitN(method, serviceMethodSeparator, 2)
		if len(elem) != 2 {
			log.Warn("unsupported method when performing authorization check", "method", method)
		} else if err := verifyAccess(elem[0], elem[1], authToken.Authorities); err != nil {
			return nil, err
		}
		// authorization check for PSI when multitenancy is enabled
		if isMultitenant := IsMultitenantFromContext(secCtx); isMultitenant {
			var authorizedPSI types.PrivateStateIdentifier
			var err error
			// does user provide PSI in the request
			if requestPSI, ok := secCtx.Value(ctxRequestPrivateStateIdentifier).(types.PrivateStateIdentifier); !ok {
				// let's try to extract from token
				authorizedPSI, err = multitenancy.ExtractPSI(authToken)
				if err != nil {
					return nil, err
				}
			} else {
				isAuthorized, err := multitenancy.IsPSIAuthorized(authToken, requestPSI)
				if err != nil {
					return nil, err
				}
				if !isAuthorized {
					return nil, multitenancy.ErrNotAuthorized
				}
				authorizedPSI = requestPSI
			}
			secCtx = WithPrivateStateIdentifier(secCtx, authorizedPSI)
			log.Debug("Determined authorized PSI", "psi", authorizedPSI)
		}
	}
	return secCtx, nil
}

// AuthenticateHttpRequest uses the provided authManager to authenticate an http request and populates
// the provided ctx with additional information useful for consumers
func AuthenticateHttpRequest(ctx context.Context, r *http.Request, authManager security.AuthenticationManager) (securityContext context.Context) {
	securityContext = ctx
	userProvidedPSI, found := extractPSI(r)
	if found {
		securityContext = context.WithValue(securityContext, ctxRequestPrivateStateIdentifier, userProvidedPSI)
	}
	if isAuthEnabled, err := authManager.IsEnabled(context.Background()); err != nil {
		// this indicates a failure in the plugin. We don't want any subsequent request unchecked
		log.Error("failure when checking if authentication manager is enabled", "err", err)
		securityContext = context.WithValue(securityContext, ctxAuthenticationError, &securityError{"internal error"})
		return
	} else if !isAuthEnabled {
		// node is not configured to be multitenant but MPS is enabled
		securityContext = WithPrivateStateIdentifier(securityContext, userProvidedPSI)
		return
	}
	if token, hasToken := extractToken(r); hasToken {
		if authToken, err := authManager.Authenticate(context.Background(), token); err != nil {
			securityContext = context.WithValue(securityContext, ctxAuthenticationError, &securityError{err.Error()})
		} else {
			securityContext = WithPreauthenticatedToken(securityContext, authToken)
		}
	} else {
		securityContext = context.WithValue(securityContext, ctxAuthenticationError, &securityError{"missing access token"})
	}
	return
}

// construct JSON RPC error message which has the ID of the request
func securityErrorMessage(forMsg *jsonrpcMessage, err error) *jsonrpcMessage {
	msg := &jsonrpcMessage{Version: vsn, ID: forMsg.ID, Error: &jsonError{
		Code:    defaultErrorCode,
		Message: err.Error(),
	}}
	ec, ok := err.(Error)
	if ok {
		msg.Error.Code = ec.ErrorCode()
	}
	return msg
}

// extractPSI tries to extract the PSI from the HTTP Header then the URL
// otherwise return the default value but still signal the caller
// that user doesn't provide PSI
func extractPSI(r *http.Request) (types.PrivateStateIdentifier, bool) {
	psi := r.Header.Get(HttpPrivateStateIdentifierHeader)
	if len(psi) == 0 {
		psi = r.URL.Query().Get(QueryPrivateStateIdentifierParamName)
	}
	if len(psi) == 0 {
		return types.DefaultPrivateStateIdentifier, false
	}
	return types.PrivateStateIdentifier(psi), true
}

// resolvePSIProvider enriches the given context with PSIProviderFunc if PSI value found
// in URL Query or env variable
func resolvePSIProvider(ctx context.Context, endpoint string) (newCtx context.Context) {
	newCtx = ctx
	var rawPSI string
	// first take from endpoint
	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return
	}
	switch parsedUrl.Scheme {
	case "http", "https", "ws", "wss":
		rawPSI = parsedUrl.Query().Get(QueryPrivateStateIdentifierParamName)
	default:
	}
	// then from the env variable
	if value := os.Getenv(EnvVarPrivateStateIdentifier); len(value) > 0 {
		rawPSI = value
	}
	if len(rawPSI) > 0 {
		// must declare type here so the context value reflects the same
		var f PSIProviderFunc = func(_ context.Context) (types.PrivateStateIdentifier, error) {
			return types.PrivateStateIdentifier(rawPSI), nil
		}
		newCtx = WithPSIProvider(ctx, f)
	}
	return
}

// encodePSI includes counter and PSI value in an JSON message ID.
// i.e.: <counter> becomes "<psi>/32"
func encodePSI(idCounterBytes []byte, psi types.PrivateStateIdentifier) json.RawMessage {
	if len(psi) == 0 {
		return idCounterBytes
	}
	newID := make([]byte, len(idCounterBytes)+len(psi)+3) // including 2 double quotes and '@'
	newID[0], newID[len(newID)-1] = '"', '"'
	copy(newID[1:len(psi)+1], psi)
	copy(newID[len(psi)+1:], append([]byte("/"), idCounterBytes...))
	return newID
}

// decodePSI extracts PSI value from an encoded JSON message ID. Return DefaultPrivateStateIdentifier
// if not found
// i.e.: "<counter>/<psi>" returns <psi>
func decodePSI(id json.RawMessage) types.PrivateStateIdentifier {
	idStr := string(id)
	if !strings.HasPrefix(idStr, "\"") || !strings.HasSuffix(idStr, "\"") {
		return types.DefaultPrivateStateIdentifier
	}
	sepIdx := strings.Index(idStr, "/")
	if sepIdx == -1 {
		return types.DefaultPrivateStateIdentifier
	}
	return types.PrivateStateIdentifier(id[1:sepIdx])
}
