package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru"
	"io/ioutil"
	"os"
	"strings"
)

//IsClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results.
//TODO (cucrisis): Cache is required here w ecant be querying the server every time
func (l *EnterpriseSecurityProvider) IsClientAuthorized(request rpcRequest) bool {
	token := request.token
	if token == "" {
		return false
	}

	// check cache first
	if entry, ok := l.tokensCache.Get(token); ok {
		introspectResponse := entry.(IntrospectResponse)
		if IsTokenExpired(introspectResponse.Created, introspectResponse.Expiration) {
			l.tokensCache.Remove(token)
		} else {
			scopes, err := parseScopeStr(introspectResponse.Scope, " ")
			if err != nil {
				return false
			}

			// search through scope -Optimize lookup
			for _, scope := range scopes {
				if isRequestAuthorized(&scope, request) {
					return true
				}
			}
			return false
		}
	}

	// issue introspect request
	response, err := getIntrospectResponse(&IntrospectRequest{
		Token:         token,
		TokenTypeHint: "access_token",
	}, &l.client, l.SecurityConfig)

	if err != nil {
		return false
	}
	if response.Active {
		// save active tokens
		l.tokensCache.Add(token, *response)
		scopes, err := parseScopeStr(response.Scope, " ")
		if err != nil {
			return false
		}

		// search through scope -Optimize lookup
		for _, scope := range scopes {
			if isRequestAuthorized(&scope, request) {
				return true
			}
		}
		return false

	}
	return false
}

//IsClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results
func (l *LocalSecurityProvider) IsClientAuthorized(request rpcRequest) bool {
	// Authenticate token
	token := request.token
	client := l.GetClientByToken(&token)

	if client == nil {
		return false
	}

	// check request scope token scope
	scopes, err := parseScopeStr(client.Scope, ",")
	if err != nil {
		return false
	}

	// search through scope -Optimize lookup
	for _, scope := range scopes {
		if isRequestAuthorized(&scope, request) {
			return true
		}
	}

	return false
}

func (l *EnterpriseSecurityProvider) Init() error {
	if l.SecurityConfig == nil {
		return fmt.Errorf("security provider confignot provided")
	}
	defaultCacheLimit := 50
	l.tokensCache, _ = lru.New(defaultCacheLimit)
	// build client
	if client, err := buildHttpClient(l.SecurityConfig); err == nil {
		l.client = *client
	} else {
		return err
	}

	return nil
}

//init local provider
func (l *LocalSecurityProvider) Init() error {
	if l.clientsFile == nil {
		return fmt.Errorf("security provider file is not set in config")
	}

	// Init structures
	if l.TokensToClients == nil || l.ClientsToTokens == nil {
		l.TokensToClients = make(map[string]ClientInfo)
		l.ClientsToTokens = make(map[string]ClientInfo)
	}
	// create file if not present
	if _, err := os.Stat(*l.clientsFile); os.IsNotExist(err) {
		f, err := os.OpenFile(*l.clientsFile, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		} else {
			err = f.Close()
			if err != nil {
				return err
			}
		}

		// write empty json
		clients := make([]ClientInfo, 0)
		clientsJson, err := json.Marshal(clients)

		if err != nil {
			return err
		}

		return ioutil.WriteFile(*l.clientsFile, clientsJson, os.ModePerm)
	}

	_, err := l.AddClientsFromFile(l.clientsFile)
	return err
}

func (l *EnterpriseSecurityProvider) GetClientByToken(clientSecret *string) *ClientInfo {
	panic("not implemented")
	return nil
}

func (l *LocalSecurityProvider) GetClientByToken(clientSecret *string) *ClientInfo {
	if c, ok := l.TokensToClients[*clientSecret]; ok {
		return &c
	}
	return nil
}

func (l *EnterpriseSecurityProvider) GetClientById(clientId *string) *ClientInfo {
	panic("not implemented")

	return nil
}

func (l *LocalSecurityProvider) GetClientById(clientId *string) *ClientInfo {
	for k := range l.ClientsToTokens {
		c := l.ClientsToTokens[k]
		if c.ClientId == *clientId {
			return &c
		}
	}

	return nil
}

func (l *EnterpriseSecurityProvider) GetClientByName(clientName *string) *ClientInfo {
	panic("not implemented")
	return nil
}

func (l *LocalSecurityProvider) GetClientByName(clientName *string) *ClientInfo {
	for k := range l.ClientsToTokens {
		c := l.ClientsToTokens[k]
		if c.Username == *clientName {
			return &c
		}
	}

	return nil
}

//addClientToFile add ClientsToTokens to json file. Assumes ClientsToTokens are well formed
func (l *EnterpriseSecurityProvider) AddClientsToFile(clients []*ClientInfo, path *string) error {
	panic("not implemented")
	return nil
}

//addClientToFile add ClientsToTokens to json file. Assumes ClientsToTokens are well formed
func (l *LocalSecurityProvider) AddClientsToFile(clients []*ClientInfo, path *string) error {
	// ensure to fall back on provider config if path is not provided
	if path == nil {
		if l.clientsFile == nil {
			return fmt.Errorf("local security provider must be set in config")
		} else {
			path = l.clientsFile
		}
	}

	// check we can read / write to file
	if _, err := os.Stat(*path); err != nil {
		return fmt.Errorf("error reading from file %v", err)
	}

	// Write to file
	jsonClients, err := json.Marshal(clients)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(*path, jsonClients, os.ModePerm)

}

//addClientsFromFile add ClientsToTokens from json file
func (l *EnterpriseSecurityProvider) AddClientsFromFile(path *string) ([]ClientInfo, error) {
	panic("not implemented")
	return nil, nil
}

//addClientsFromFile add ClientsToTokens from json file
func (l *LocalSecurityProvider) AddClientsFromFile(path *string) ([]ClientInfo, error) {
	// ensure to fall back on provider config if path is not provided
	if path == nil {
		if l.clientsFile == nil {
			return nil, fmt.Errorf("local security provider must be set in config")
		} else {
			path = l.clientsFile
		}
	}

	// check we can read from file
	if _, err := os.Stat(*path); err != nil {
		return nil, fmt.Errorf("error reading from file %v", err)
	}

	// Read & unmarshall the json file
	clients, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	var clientToTokenList []ClientInfo
	if err := json.Unmarshal(clients, &clientToTokenList); err != nil {
		return nil, err
	}

	// Adding process most not fail on one entry
	// but it must report error tell it solved
	var addingError *error
	for k := range clientToTokenList {
		client := clientToTokenList[k]
		if err := l.AddClient(&client); err != nil {
			addingError = &err
		}
	}
	if addingError == nil {
		return clientToTokenList, nil
	} else {

		return nil, *addingError
	}
}

func (l *EnterpriseSecurityProvider) SetClientScope(clientName string, scope string) error {
	panic("not implemented")
	return nil
}

func (l *LocalSecurityProvider) SetClientScope(clientName string, scope string) error {
	client := l.GetClientByName(&clientName)
	if client == nil {
		return fmt.Errorf("client not found")
	}

	scopeToSet, err := cleanScope(scope)
	if err != nil {
		return err
	}
	client.Scope = scopeToSet

	if err := l.RemoveClient(&client.Username); err != nil {
		return err
	}

	if err := l.AddClient(client); err != nil {
		return err
	}
	return nil
}

func (l *EnterpriseSecurityProvider) SetClientStatus(clientName string, status bool) error {
	panic("not implemented")
	return nil
}

func (l *LocalSecurityProvider) SetClientStatus(clientName string, status bool) error {
	client := l.GetClientByName(&clientName)
	if client == nil {
		return fmt.Errorf("client not found")
	}

	if err := l.RemoveClient(&client.Username); err != nil {
		return err
	}
	client.Active = status
	if err := l.AddClient(client); err != nil {
		return err
	}
	return nil
}

func (l *EnterpriseSecurityProvider) NewClient(clientName string, clientId string, secret string, scope string, active bool) (ClientInfo, error) {
	panic("not implemented")
	return ClientInfo{}, nil
}

//NewClient creates a new client struct
func (l *LocalSecurityProvider) NewClient(clientName string, clientId string, secret string, scope string, active bool) (ClientInfo, error) {
	clientName, err := cleanString(clientName)
	if err != nil {
		return ClientInfo{}, err
	}

	clientScope, err := cleanScope(scope)
	if err != nil {
		return ClientInfo{}, err
	}

	if clientName == "" {
		return ClientInfo{}, fmt.Errorf("clientName must be provided. Only alpha numeric is accepted")
	}

	// Generate Defaults
	if clientId == "" {
		secGuid, err := uuid.NewRandom()
		if err != nil {
			return ClientInfo{}, err
		}
		clientId = secGuid.String()
	}
	if secret == "" {
		secGuid, err := uuid.NewRandom()
		if err != nil {
			return ClientInfo{}, err
		}
		secret = secGuid.String()
	}

	return ClientInfo{
		ClientId: clientId,
		Scope:    clientScope,
		Secret:   secret,
		Username: clientName,
		Active:   active,
	}, nil

}

//AddClient adds the clients to memory and writes it to file.
func (l *EnterpriseSecurityProvider) AddClient(client *ClientInfo) error {
	panic("not implemented")
	return nil
}

//AddClient adds the clients to memory and writes it to file.
func (l *LocalSecurityProvider) AddClient(client *ClientInfo) error {
	if client == nil {
		return fmt.Errorf("client must be provided")
	}

	if l.GetClientByName(&client.Username) != nil || l.GetClientByToken(&client.Secret) != nil || l.GetClientById(&client.ClientId) != nil {
		return fmt.Errorf("client with same username, secret or identifier exists")
	}

	// Init structures
	l.ClientsToTokens[client.ClientId] = *client
	l.TokensToClients[client.Secret] = *client

	return l.AddClientsToFile(l.GetClientsList(), l.clientsFile)
}

func (l *EnterpriseSecurityProvider) GetClientsList() []*ClientInfo {
	panic("not implemented")
	return nil
}

//listClients return list of clients
func (l *LocalSecurityProvider) GetClientsList() []*ClientInfo {
	clients := make([]*ClientInfo, len(l.ClientsToTokens))
	var counter = 0
	for k := range l.ClientsToTokens {
		c := l.ClientsToTokens[k]
		clients[counter] = &c
		counter++
	}
	return clients
}

//removeClient remove clients from memory and file.
func (l *EnterpriseSecurityProvider) RemoveClient(clientName *string) error {
	panic("not implemented")
	return nil
}

//removeClient remove clients from memory and file.
func (l *LocalSecurityProvider) RemoveClient(clientName *string) error {
	if client := l.GetClientByName(clientName); client != nil {
		delete(l.ClientsToTokens, client.ClientId)
		delete(l.TokensToClients, client.Secret)

		return l.AddClientsToFile(l.GetClientsList(), l.clientsFile)

	} else {
		return fmt.Errorf("client doesnt exist")
	}
}

func (l *EnterpriseSecurityProvider) RegenerateClientSecret(clientName *string) (*ClientInfo, error) {
	panic("not implemented")
	return nil, nil
}

//regenerateClientSecret regenerate client secret.
func (l *LocalSecurityProvider) RegenerateClientSecret(clientName *string) (*ClientInfo, error) {
	client := l.GetClientByName(clientName)
	if client == nil {
		return nil, fmt.Errorf("client doesnt exist")
	}

	secGuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	client.Secret = secGuid.String()

	if err = l.RemoveClient(clientName); err != nil {
		return nil, err
	}

	if err = l.AddClient(client); err != nil {
		return nil, err
	}

	return client, nil
}

func RegisterProvider(ctx *SecurityContext, log log.Logger) {
	// Ensure Provider is created before APIs is invoked otherwise
	if ctx.Config != nil {
		switch provider := strings.ToLower(ctx.Config.ProviderType); provider {
		case LocalSecProvider:
			log.Info("register local security provider", "RPC security", "Enabled")

			ctx.Provider = &LocalSecurityProvider{
				clientsFile: ctx.Config.ProviderInformation.LocalProviderFile,
			}

			if err := ctx.Provider.Init(); err != nil {
				Fatalf(err.Error())
			}

		case EnterpriseSecProvider:
			log.Info("register enterprise security provider", "RPC security", "Enabled")

			ctx.Provider = &EnterpriseSecurityProvider{
				SecurityConfig:      ctx.Config,
				IntrospectURL:       ctx.Config.ProviderInformation.EnterpriseProviderIntrospectionURL,
				ProviderCertificate: ctx.Config.ProviderInformation.EnterpriseProviderCertificateInfo,
			}

			log.Info("security provided init")

			if err := ctx.Provider.Init(); err != nil {
				Fatalf(err.Error())
			}

		default:
			Fatalf("rpc security provider Type not supported. supported providers are [local, enterprise]. change config file")
		}
	}
}
