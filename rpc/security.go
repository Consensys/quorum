package rpc

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/multitenancy"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type securityContextKey string
type securityContext context.Context

const (
	HttpAuthorizationHeader              = "Authorization"
	HttpPrivateStateIdentifierHeader     = "PSI"
	QueryPrivateStateIdentifierParamName = "PSI"
	// this key is set by server to indicate if server supports mulitenancy
	ctxIsMultitenant = securityContextKey("IS_MULTITENANT")
	// this key is set into the secured context to indicate
	// the authorized private state being operated on for the request
	ctxPrivateStateIdentifier = securityContextKey("PRIVATE_STATE_IDENTIFIER")
	// this key is set into the request context to indicate
	// the private state being operated on for the request
	ctxRequestPrivateStateIdentifier = securityContextKey("REQUEST_PRIVATE_STATE_IDENTIFIER")
	// this key is exported for WS transport
	CtxCredentialsProvider = securityContextKey("CREDENTIALS_PROVIDER") // key to save reference to rpc.HttpCredentialsProviderFunc
	// keys used to save values in request context
	ctxAuthenticationError   = securityContextKey("AUTHENTICATION_ERROR")   // key to save error during authentication before processing the request body
	CtxPreauthenticatedToken = securityContextKey("PREAUTHENTICATED_TOKEN") // key to save the preauthenticated token once authenticated
)

type securityContextConfigurer interface {
	Configure(secCtx securityContext)
}

type securityContextResolver interface {
	Resolve() securityContext
}

type securityError struct{ message string }

// Provider function to return token being injected in Authorization http request header
type HttpCredentialsProviderFunc func(ctx context.Context) (string, error)

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
	expiredAt, err := ptypes.Timestamp(token.ExpiredAt)
	if err != nil {
		return fmt.Errorf("invalid timestamp in token: %s", err)
	}
	if time.Now().Before(expiredAt) {
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
func secureCall(resolver securityContextResolver, msg *jsonrpcMessage) (context.Context, error) {
	secCtx := resolver.Resolve()
	if secCtx == nil {
		return context.Background(), nil
	}
	if err, hasError := secCtx.Value(ctxAuthenticationError).(error); hasError {
		return nil, err
	}
	if authToken, isPreauthenticated := secCtx.Value(CtxPreauthenticatedToken).(*proto.PreAuthenticatedAuthenticationToken); isPreauthenticated {
		if err := verifyExpiration(authToken); err != nil {
			return nil, err
		}
		elem := strings.SplitN(msg.Method, serviceMethodSeparator, 2)
		if len(elem) != 2 {
			log.Warn("unsupported method when performing authorization check", "method", msg.Method)
		} else if err := verifyAccess(elem[0], elem[1], authToken.Authorities); err != nil {
			return nil, err
		}
		// authorization check for PSI when multitenancy is enabled
		if isMultitenant, ok := secCtx.Value(ctxIsMultitenant).(bool); ok && isMultitenant {
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
				isAuthorized, err := multitenancy.Authorize(authToken, (&multitenancy.PrivateStateSecurityAttribute{}).WithPSI(requestPSI))
				if err != nil {
					return nil, err
				}
				if !isAuthorized {
					return nil, multitenancy.ErrNotAuthorized
				}
				authorizedPSI = requestPSI
			}
			secCtx = context.WithValue(secCtx, ctxPrivateStateIdentifier, authorizedPSI)
		}
	}
	return secCtx, nil
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

// try to extract the PSI from the HTTP Header then the URL
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
