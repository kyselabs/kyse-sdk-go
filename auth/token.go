package auth

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/kyselabs/kyse-sdk-go/rest"
)

type TokenBased struct {
	rest *rest.Client

	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (t *TokenBased) ToToken() *Token {
	return &Token{
		Access:  t.Access,
		Refresh: t.Refresh,
	}
}

func (t *TokenBased) SetRestClient(client *rest.Client) {
	t.rest = client
}

func (t *TokenBased) Authenticate() (*Token, error) {
	return &Token{
		Access:  t.Access,
		Refresh: t.Refresh,
	}, nil
}

type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (t *Token) ExpiredAccess() bool {
	expiredAccess, err := IsExpired(t.Access)
	if err != nil {
		return true
	}

	return expiredAccess
}

func (t *Token) ExpiredRefresh() bool {
	expiredRefresh, err := IsExpired(t.Refresh)
	if err != nil {
		return true
	}

	return expiredRefresh
}

func IsExpired(jwtToken string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtToken, jwt.MapClaims{})
	if err != nil {
		return true, errors.New("can't parse this token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true, errors.New("can't get token's claims")
	}

	if claims["exp"] != nil {
		var tm time.Time

		switch exp := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(exp), 0)

		case string:
			v, _ := strconv.ParseInt(exp, 10, 64)
			tm = time.Unix(v, 0)

		case json.Number:
			v, _ := exp.Int64()
			tm = time.Unix(v, 0)

		}

		return tm.Before(time.Now()), nil
	}

	return false, nil
}
