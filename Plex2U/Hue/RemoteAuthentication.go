package Hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Authentication struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresIn  string `json:"access_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
}

type httpClient struct {
	http.Client
	baseUrl string
	token   string
	secret  string
}

var client *httpClient

func (c *httpClient) callHue(method string, path string, content []byte) ([]byte, error) {
	req, err := http.NewRequest(method, c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}
	if content != nil {
		req, err = http.NewRequest(method, c.baseUrl+path, bytes.NewReader(content))
		if err != nil {
			return nil, err
		}
	}
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

func AuthGetFromCode(code string, token *string, secret *string) (*Authentication, error) {
	auth := Authentication{}
	content, err := getHttpClient(token, secret).callHue("GET", fmt.Sprintf("/oauth2/token?code=rp1xXeEu&grant_type=authorization_code"), nil)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(content, &auth); err != nil {
		return nil, err
	}

	return &auth, nil
}
func getHttpClient(token *string, secret *string) *httpClient {
	if client == nil {
		client = &httpClient{Client: http.Client{}, baseUrl: "https://api.meethue.com/", token: *token, secret: *secret}
	}
	return client
}

func CreateCredentials() string {
	return ""
}

func RenewCredentials() {

}
