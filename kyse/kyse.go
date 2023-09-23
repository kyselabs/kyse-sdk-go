package kyse

import (
	"fmt"

	"github.com/kyselabs/kyse-sdk-go/auth"
	"github.com/kyselabs/kyse-sdk-go/rest"
)

type Kyse struct {
	AuthClient *auth.Client
	RestClient *rest.Client
}

func (k *Kyse) Request(method string, path string, headers map[string]string, body interface{}) (*rest.Response, error) {
	headersWithAuth := make(map[string]string)

	if token, err := k.AuthClient.Authenticate(); err == nil {
		headersWithAuth["Authorization"] = fmt.Sprintf("Bearer %s", token.Access)
	}

	for header, value := range headers {
		headersWithAuth[header] = value
	}

	return k.RestClient.Request(method, path, headersWithAuth, body)
}
