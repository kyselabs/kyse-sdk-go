package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

type Response struct {
	http.Response
}

func NewRestClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL, HttpClient: &http.Client{}}
}

func (r *Response) Unmarshal(target interface{}) error {
	content, err := io.ReadAll(r.Response.Body)
	json.Unmarshal(content, target)
	defer r.Response.Body.Close()
	return err
}

func (c *Client) Request(method string, path string, headers map[string]string, body interface{}) (*Response, error) {
	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	doResponse, err := c.HttpClient.Do(req)
	return &Response{Response: *doResponse}, err

}
