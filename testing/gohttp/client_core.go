package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *httpClient) getRequestBody(contentType string, b interface{}) ([]byte, error) {
	if b == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(b)
	case "application/xml":
		return xml.Marshal(b)
	default:
		return json.Marshal(b)
	}
}

func (c *httpClient) do(m string, u string, h http.Header, b interface{}) (*http.Response, error) {

	fullHeaders := c.getRequestHeaders(h)
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), b)
	if err != nil {
		return nil, errors.New("unable to create new request")
	}
	request, err := http.NewRequest(m, u, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}
	request.Header = fullHeaders
	return c.client.Do(request)

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
