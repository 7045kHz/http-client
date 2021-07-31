package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/7045kHz/http-client/gohttp"
	"github.com/7045kHz/http-client/gomime"
)

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

var (
	httpClient = getHttpClient()
)

func main() {

	GetEndpoints()

}
func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	fmt.Println("Starting new client")
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Fedes-Computer").
		SetNtlm(true).
		Build()
	return client
}
func GetEndpoints() (*Endpoints, error) {
	// Make the request and wait for the response:
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		// Deal with the error as you need:
		return nil, err
	}

	// Interacting with the response:
	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status))
	//fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	// Processing JSON responses:
	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with the unmarshal error as you need:
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))
	return &endpoints, nil
}
