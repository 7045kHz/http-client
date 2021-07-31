package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/7045kHz/http-client/gohttp"
	"github.com/7045kHz/http-client/gomime"
)

type ConfigBenchmark struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	EndpointFile string `json:"endpoint_file"`
}
type Event struct {
	Events []Endpoints `json:"Event"`
}
type Endpoints struct {
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	Url               string           `json:"url"`
	PayloadFile       string           `json:"payloadfile"`
	Method            string           `json:"method"`
	SummaryOutputFile string           `json:"summary_output_file"`
	BodyOutputFile    string           `json:"body_output_file"`
	Headers           []EndpointHeader `json:"headers"`
}

type EndpointHeader struct {
	Header_Key   string `json:"header_key"`
	Header_Value string `json:"header_value"`
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
		SetUserAgent("Test-Computer").
		SetNtlm(false).
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
