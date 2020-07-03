package rpc

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
)

type securityContextKey string
type securityContext context.Context

const (
	HttpAuthorizationHeader = "Authorization"
	// this key is exported for WS transport
	CtxCredentialsProvider = securityContextKey("CREDENTIALS_PROVIDER") // key to save reference to rpc.HttpCredentialsProviderFunc
	// keys used to save values in request context
	ctxAuthenticationError   = securityContextKey("AUTHENTICATION_ERROR")   // key to save error during authentication before processing the request body
	ctxPreauthenticatedToken = securityContextKey("PREAUTHENTICATED_TOKEN") // key to save the preauthenticated token once authenticated
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
func secureCall(resolver securityContextResolver, msg *jsonrpcMessage) error {
	secCtx := resolver.Resolve()
	if secCtx == nil {
		return nil
	}
	if err, hasError := secCtx.Value(ctxAuthenticationError).(error); hasError {
		return err
	}
	if authToken, isPreauthenticated := secCtx.Value(ctxPreauthenticatedToken).(*proto.PreAuthenticatedAuthenticationToken); isPreauthenticated {
		if err := verifyExpiration(authToken); err != nil {
			return err
		}
		elem := strings.SplitN(msg.Method, serviceMethodSeparator, 2)
		if len(elem) != 2 {
			log.Warn("unsupported method when performing authorization check", "method", msg.Method)
		} else if err := verifyAccess(elem[0], elem[1], authToken.Authorities); err != nil {
			return err
		}
	}
	return nil
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
