package rpc

import (
	"encoding/json"
	"fmt"
)

//isClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *EnterpriseSecurityProvider) isClientAuthorized(request rpcRequest) bool {
	fmt.Println("Send Introspect for:")
	fmt.Println(request)
	if request.token == "cucrisis" {
		if request.service != "admin" {
			return true
		}
	}

	return false
}

func (l *EnterpriseSecurityProvider) init() error {
	 return nil
}


//isClientAuthorized Parse the RPC Request, Call send Introspect Request & Parse results
func (ctx *LocalSecurityProvider) isClientAuthorized(request rpcRequest) bool {
	fmt.Println("Send Introspect for:")
	fmt.Println(request)
	if request.token == "cucrisis" {
		if request.service != "admin" {
			return true
		}
	}

	return false
}

func (l *LocalSecurityProvider) init() error {
	 db , err := newDatabase(*l.securityProviderFile)
	 if err != nil {
	 	l.securityDatabase = db
	 	defer l.securityDatabase.Close()
	 	return nil
	 } else {
		 return err
	 }


}

func FindClient(clientName *string, l *LocalSecurityProvider) (LocalProviderClient, error) {
	data, err := l.securityDatabase.Get([]byte(*clientName), nil)
	if err != nil {
		return LocalProviderClient{}, err
	}

	var client  LocalProviderClient
	err = json.Unmarshal(data, &client)
	if err != nil {
		return LocalProviderClient{}, err
	}
	return  client, nil
}

func (l *LocalSecurityProvider) addClientsFromFile(fileName *string) {

}

func AddLocalRpcClient(client LocalProviderClient, l *LocalSecurityProvider) error {
	clientJson, err := json.Marshal(client)
	if err != nil {
		return err
	}
	return l.securityDatabase.Put([]byte(client.ClientName),clientJson, nil)
}

func (l *LocalSecurityProvider) listClients() {

}

func (l *LocalSecurityProvider) removeClient(clientName *string) {

}

func (l *LocalSecurityProvider) regenerateClient(clientName *string) {

}

