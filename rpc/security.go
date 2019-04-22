package rpc

import "net/http"

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
	return 0, nil
}

func (ctx *SecurityContext)  ProcessWSRequest(r *http.Request) (int, error){
	return 0, nil
}


// GetDefaultSecurityContext returns a disabled context
func GetDefaultSecurityContext() SecurityContext {
	return SecurityContext{Enabled:false}
}


