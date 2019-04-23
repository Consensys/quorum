package rpc

import (
	"github.com/pkg/errors"
	"net/http"
)

// RPC Security Configuration
type SecurityConfig struct {
	 ProviderType string `json:"providerType"`
}

// RPC Security Context
type SecurityContext struct {
	 Enabled bool
	 Config *SecurityConfig
}

func (ctx *SecurityContext) ProcessHttpRequest(r *http.Request) (int, error){
	if ctx.Enabled && ctx.Config == nil {
		return http.StatusUnauthorized, errors.New("Unauthorized")
	}



	return http.StatusOK, nil
}

func (ctx *SecurityContext)  ProcessWSRequest(r *http.Request) (int, error){
	if ctx.Enabled && ctx.Config == nil {
		return http.StatusUnauthorized, errors.New("Unauthorized")
	}

}

// RPC Security Console APIs
type SecurityApi struct {

}


func test(){

}


// GetDefaultDenyAllSecurityContext returns a disabled context
func GetDefaultDenyAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled:true}
}
// GetDefaultAllowAllSecurityContext returns a disabled context
func GetDefaultAllowAllSecurityContext() SecurityContext {
	return SecurityContext{Enabled:false}
}


