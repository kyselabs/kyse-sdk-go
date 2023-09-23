package auth

import (
	"errors"
	"net/http"

	"github.com/kyselabs/kyse-sdk-go/rest"
)

type Client struct {
	RestClient *rest.Client
	AuthMethod AuthMethod
	Callback   func(token *Token)
	TempToken  *Token
}

type AuthMethod interface {
	SetRestClient(*rest.Client)
	Authenticate() (*Token, error)
}

type RefreshTokenPayload struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthClient(url string, authMethod AuthMethod, callbackToken func(token *Token)) *Client {
	restClient := rest.NewRestClient(url)
	if authMethod != nil {
		authMethod.SetRestClient(restClient)
	}

	client := &Client{
		AuthMethod: authMethod,
		RestClient: restClient,
		Callback:   callbackToken,
	}

	switch authMethod := authMethod.(type) {
	case *TokenBased:
		client.TempToken = authMethod.ToToken()
	}

	return client
}

func (c *Client) GetDeviceCode() (*DeviceBased, error) {
	if _, ok := c.AuthMethod.(*DeviceBased); ok {
		response, err := c.RestClient.Request(http.MethodPost, "/oauth/device/code", nil, nil)

		if err == nil {
			response.Unmarshal(&c.AuthMethod)
		}

		return c.AuthMethod.(*DeviceBased), nil
	}

	return nil, errors.New("invalid auth method")
}

func (c *Client) RefreshToken(refreshToken string) (*Token, error) {
	response, err := c.RestClient.Request(http.MethodPost, "/token", nil, RefreshTokenPayload{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	})

	if err == nil {
		var token Token
		response.Unmarshal(&token)

		return &token, err
	}

	return nil, err
}

func (c *Client) Authenticate() (*Token, error) {

	if c.TempToken != nil {
		if !c.TempToken.ExpiredAccess() {
			return c.TempToken, nil
		} else {
			if !c.TempToken.ExpiredRefresh() {
				if token, err := c.RefreshToken(c.TempToken.Refresh); err == nil {
					c.TempToken = token
					if c.Callback != nil {
						c.Callback(c.TempToken)
					}
					return token, nil
				}
			}
		}
	}

	token, err := c.AuthMethod.Authenticate()
	c.TempToken = token
	if c.Callback != nil {
		c.Callback(c.TempToken)
	}

	return c.TempToken, err
}
