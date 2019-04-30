package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)

//IsClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *EnterpriseSecurityProvider) IsClientAuthorized(request rpcRequest) bool {
	fmt.Println("Send Introspect for:")

	if request.token == "cucrisis" {
		 return true
	}

	return false
}

//SetType the clients to memory and writes it to file.
func (l *EnterpriseSecurityProvider) SetType(typeName string) {
	l.providerTypeName = typeName
}

//getType the clients to memory and writes it to file.
func (l *EnterpriseSecurityProvider) GetType() string {
	return l.providerTypeName
}

func (l *EnterpriseSecurityProvider) Init() error {
	return nil
}

//IsClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results
func (l *LocalSecurityProvider) IsClientAuthorized(request rpcRequest) bool {
	fmt.Println("Send Introspect for:")
	if request.token == "cucrisis" {
		return true
	}
	return false
}

//init local provider
func (l *LocalSecurityProvider) Init() error {
	if l.clientsFile == nil {
		return fmt.Errorf("security provider file is not set in config")
	}

	// create file if not present
	if _, err := os.Stat(*l.clientsFile); os.IsNotExist(err) {
		f , err :=os.OpenFile(*l.clientsFile, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}else{
			err = f.Close()
			if err != nil {
				return err
			}
		}

		// write empty json
		clients := make([]*ClientInfo, 0)
		clientsJson, _ := json.Marshal(clients)
		return ioutil.WriteFile(*l.clientsFile, clientsJson, 0644)
	}

	return l.AddClientsFromFile(l.clientsFile)
}

func (l *LocalSecurityProvider) GetClientByToken(clientSecret *string) *ClientInfo {
	if c, ok := l.tokensToClients[*clientSecret]; ok {
		return c
	}
	return nil

}

func (l *LocalSecurityProvider) GetClientById(clientId *string) *ClientInfo {

	for _, c := range l.clientsToTokens {
		if c.ClientId == *clientId {
			return c
		}
	}

	return nil
}

func (l *LocalSecurityProvider) GetClientByName(clientName *string) *ClientInfo {

	for _, c := range l.clientsToTokens {
		if c.Username == *clientName {
			return c
		}
	}

	return nil
}

//addClientToFile add clientsToTokens to json file. Assumes clientsToTokens are well formed
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
		return ioutil.WriteFile(*path, jsonClients, 0644)
	}

	return nil
}

//addClientsFromFile add clientsToTokens from json file
func (l *LocalSecurityProvider) AddClientsFromFile(path *string) error {
	// ensure to fall back on provider config if path is not provided
	if path == nil {
		if l.clientsFile == nil {
			return fmt.Errorf("local security provider must be set in config")
		} else {
			path = l.clientsFile
		}
	}

	// check we can read from file
	if _, err := os.Stat(*path); err != nil {
		return fmt.Errorf("error reading from file %v", err)
	}

	// Read & unmarshall the json file
	clients, err := ioutil.ReadFile(*path)
	if err != nil {
		return err
	}

	var clientToTokenList []ClientInfo
	if err := json.Unmarshal(clients, &clientToTokenList); err != nil {
		return err
	}

	// Adding process most not fail on one entry
	// but it must report error tell it solved
	var addingError *error
	for _, client := range clientToTokenList {
		if err := l.AddClient(&client); err != nil {
			addingError = &err
		}
	}
	if addingError == nil {
		return nil
	} else {

		return *addingError
	}
}

//NewClient creates a new client struct
func (l *LocalSecurityProvider) NewClient(clientName string, clientId string, secret string, scope string, exp int) (ClientInfo, error) {
	clientName, err := cleanString(clientName)

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
		ClientId:   clientId,
		Scope:      scope,
		Secret:     secret,
		Username:   clientName,
		Expiration: exp,
	}, nil

}

//SetType the clients to memory and writes it to file.
func (l *LocalSecurityProvider) SetType(typeName string) {
	l.providerTypeName = typeName
}

//getType the clients to memory and writes it to file.
func (l *LocalSecurityProvider) GetType() string {
	return l.providerTypeName
}

//AddClient adds the clients to memory and writes it to file.
func (l *LocalSecurityProvider) AddClient(client *ClientInfo) error {
	if client == nil {
		return fmt.Errorf("client must be provided")
	}

	if l.GetClientByName(&client.Username) != nil || l.GetClientByToken(&client.Secret) != nil || l.GetClientById(&client.ClientId) != nil {
		return fmt.Errorf("client with same username, secret or identifier exists")
	}

	// Add clients to memory
	l.clientsToTokens[client.ClientId] = client
	l.tokensToClients[client.Secret] = client

	return l.AddClientsToFile(l.GetClientsList(), l.clientsFile)
}

//listClients return list of clients
func (l *LocalSecurityProvider) GetClientsList() []*ClientInfo {
	// Write client to file
	clients := make([]*ClientInfo, len(l.clientsToTokens))
	counter := 0

	for _, c := range l.clientsToTokens {
		clients[counter] = c
		counter++
	}
	return clients
}

//removeClient remove clients from memory and file.
func (l *LocalSecurityProvider) RemoveClient(clientName *string) error {
	if l.GetClientByName(clientName) == nil {
		return fmt.Errorf("client doesnt exist")
	}

	client := l.GetClientByName(clientName)

	delete(l.clientsToTokens, client.Username)
	delete(l.tokensToClients, client.Secret)

	return l.AddClientsToFile(l.GetClientsList(), l.clientsFile)
}

//regenerateClientSecret regenerate client secret.
func (l *LocalSecurityProvider) RegenerateClientSecret(clientName *string) (*ClientInfo, error) {
	if l.GetClientByName(clientName) == nil {
		return nil, fmt.Errorf("client doesnt exist")
	}

	client := l.GetClientByName(clientName)

	secGuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	client.Secret = secGuid.String()
	return client, nil
}
