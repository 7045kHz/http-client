package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(m string, u string, h http.Header, b interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(m, u, nil)
	if err != nil {
		return nil, errors.New("Unable to create new request")
	}
	// Adding custom headers
	for header, value := range h {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}
	return client.Do(request)

}
