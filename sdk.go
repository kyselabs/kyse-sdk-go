package sdk

import (
	"os"

	"github.com/kyselabs/kyse-sdk-go/auth"
	"github.com/kyselabs/kyse-sdk-go/kyse"
	"github.com/kyselabs/kyse-sdk-go/rest"
)

var (
	authAPI = "https://auth.kyse.io"
	kyseAPI = "https://api.kyse.io"
)

func init() {
	if value := os.Getenv("AUTH_API_ADDRESS"); value != "" {
		authAPI = value
	}

	if value := os.Getenv("KYSE_API_ADDRESS"); value != "" {
		kyseAPI = value
	}
}

func NewKyse(authMethod auth.AuthMethod, callbackToken func(token *auth.Token)) *kyse.Kyse {
	return &kyse.Kyse{
		AuthClient: auth.NewAuthClient(authAPI, authMethod, callbackToken),
		RestClient: rest.NewRestClient(kyseAPI),
	}
}
