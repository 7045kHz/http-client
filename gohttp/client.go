package gohttp

import "net/http"

type httpClient struct {
	Headers http.Header
}

type HttpClient interface {
	Get(u string, h http.Header) (*http.Response, error)
	Post(u string, h http.Header, b interface{}) (*http.Response, error)
	Put(u string, h http.Header, b interface{}) (*http.Response, error)
	Patch(u string, h http.Header, b interface{}) (*http.Response, error)
	Delete(u string, h http.Header) (*http.Response, error)
	SetHeaders(h http.Header)
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

func (c *httpClient) SetHeaders(h http.Header) {
	c.Headers = h

}
func (c *httpClient) Get(u string, h http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, u, h, nil)
}
func (c *httpClient) Post(u string, h http.Header, b interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, u, h, b)
}
func (c *httpClient) Put(u string, h http.Header, b interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, u, h, b)
}
func (c *httpClient) Patch(u string, h http.Header, b interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, u, h, b)
}
func (c *httpClient) Delete(u string, h http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, u, h, nil)
}
