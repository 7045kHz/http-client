package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	maxIdleConnsPerHost int
	connectionTimeout   time.Duration
	requestTimeout      time.Duration
	client              *http.Client
	Headers             http.Header
}

type HttpClient interface {
	Get(u string, h http.Header) (*http.Response, error)
	Post(u string, h http.Header, b interface{}) (*http.Response, error)
	Put(u string, h http.Header, b interface{}) (*http.Response, error)
	Patch(u string, h http.Header, b interface{}) (*http.Response, error)
	Delete(u string, h http.Header) (*http.Response, error)
	SetHeaders(h http.Header)
	SetConnectionTimeout(t time.Duration)
	SetRequestTimeout(t time.Duration)
	SetMaxIdleConnections(i int)
}

func New() HttpClient {
	client := http.Client{}
	httpClient := &httpClient{
		client: &client,
	}
	return httpClient
}

func (c *httpClient) SetConnectionTimeout(t time.Duration) {
	c.connectionTimeout = t
}

func (c *httpClient) SetRequestTimeout(t time.Duration) {
	c.requestTimeout = t
}
func (c *httpClient) SetMaxIdleConnections(i int) {
	c.maxIdleConnsPerHost = i
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
