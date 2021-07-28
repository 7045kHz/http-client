package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(m string, u string, h http.Header, b interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(m, u, nil)
	if err != nil {
		return nil, errors.New("unable to create new request")
	}
	fullHeaders := c.getRequestHeaders(h)
	request.Header = fullHeaders
	return client.Do(request)

}
func (c *httpClient) getRequestHeaders(h http.Header) http.Header {
	result := make(http.Header)
	// Adding common headers
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	// Adding custom headers
	for header, value := range h {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
