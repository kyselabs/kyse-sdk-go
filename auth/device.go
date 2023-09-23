package auth

import (
	"errors"
	"net/http"

	"github.com/kyselabs/kyse-sdk-go/rest"
)

type DeviceBased struct {
	rest *rest.Client

	DeviceCode              string `json:"device_code"`
	ExpiresIn               int    `json:"expires_in"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
}

type GetTokenPayload struct {
	GrantType  string `json:"grant_type"`
	DeviceCode string `json:"device_code"`
}

func (d *DeviceBased) SetRestClient(client *rest.Client) {
	d.rest = client
}

func (d *DeviceBased) Authenticate() (*Token, error) {

	if d.DeviceCode == "" {
		return nil, errors.New("trying to authenticate before getting device codes")
	}

	response, err := d.rest.Request(http.MethodPost, "/token", nil, GetTokenPayload{
		GrantType:  "urn:ietf:params:oauth:grant-type:device_code",
		DeviceCode: d.DeviceCode,
	})

	if err == nil {
		var token Token
		response.Unmarshal(&token)

		return &token, err
	}

	return nil, err
}
