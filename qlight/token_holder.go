package qlight

var token string

func GetCurrentToken() string {
	return token
}

func SetCurrentToken(newToken string) {
	token = newToken
}
