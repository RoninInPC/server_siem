package sender

import (
	"net/http"
	"net/url"
	"strings"
)

type CommandJWT interface {
	Command(address, message string) (*http.Response, error)
}

type CommandJWTPostForm struct {
}

func (c CommandJWTPostForm) Command(address, message string) (*http.Response, error) {
	return http.PostForm(address, url.Values{"json": {message}})
}

type CommandJWTUpdate struct {
}

func (c CommandJWTUpdate) Command(address, message string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPatch, address, strings.NewReader(url.Values{"json": {message}}.Encode()))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

type CommandJWTDelete struct {
	Address string
}

func (c CommandJWTDelete) Command(address, message string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, address, strings.NewReader(url.Values{"json": {message}}.Encode()))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
