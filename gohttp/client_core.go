package gohttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/7045kHz/http-client/core"
	"github.com/7045kHz/http-client/gohttp_mock"
	"github.com/7045kHz/http-client/gomime"
	"github.com/Azure/go-ntlmssp"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
	defaultNtlmSetting        = false
	defaultTlsInsecure        = true
	defaultProxyServer        = ""
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	response, err := c.getHttpClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := core.Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
	}
	return &finalResponse, nil
}

func (c *httpClient) getHttpClient() core.HttpClient {
	if gohttp_mock.MockupServer.IsEnabled() {
		return gohttp_mock.MockupServer.GetMockedClient()
	}

	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		proxyUrl, _ := url.Parse(c.getProxyServer())
		//func setProxyMode(proxyUrl *url.URL, c *httpClient) *url.URL {

		var trP *url.URL = setProxyMode(proxyUrl, c)

		if c.getNtlm() {
			log.Printf("Using NTLMv2 HTTP Client")
			c.client = &http.Client{
				Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
				Transport: ntlmssp.Negotiator{
					RoundTripper: &http.Transport{
						TLSClientConfig:       &tls.Config{InsecureSkipVerify: c.getTlsInsecureVerify()},
						Proxy:                 http.ProxyURL(trP),
						MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
						ResponseHeaderTimeout: c.getResponseTimeout(),
						DialContext: (&net.Dialer{
							Timeout: c.getConnectionTimeout(),
						}).DialContext,
					},
				},
				/*

					Transport: &http.Transport{
						RoundTripper:          &http.Transport{},
						MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
						ResponseHeaderTimeout: c.getResponseTimeout(),
						DialContext: (&net.Dialer{
							Timeout: c.getConnectionTimeout(),
						}).DialContext,
					},*/
			}
		} else {

			log.Printf("Using NON-NTLMv2 HTTP Client")
			c.client = &http.Client{
				Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
				Transport: &http.Transport{
					Proxy:                 http.ProxyURL(trP),
					TLSClientConfig:       &tls.Config{InsecureSkipVerify: c.getTlsInsecureVerify()},
					MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
					ResponseHeaderTimeout: c.getResponseTimeout(),
					DialContext: (&net.Dialer{
						Timeout: c.getConnectionTimeout(),
					}).DialContext,
				},
			}
		}
	})
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}
func (c *httpClient) getNtlm() bool {
	if c.builder.ntlm {
		return c.builder.ntlm
	}
	return defaultNtlmSetting
}

func (c *httpClient) getTlsInsecureVerify() bool {
	if c.builder.tls_insecure_verify {
		return c.builder.tls_insecure_verify
	}
	return defaultTlsInsecure
}

func (c *httpClient) getProxyServer() string {
	if c.builder.proxy != "" {
		return c.builder.proxy
	}
	return defaultProxyServer
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)

	case gomime.ContentTypeXml:
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}
func setProxyMode(proxyUrl *url.URL, c *httpClient) *url.URL {
	if c.getProxyServer() == "" {
		var v *url.URL = nil
		return v
	} else {
		return proxyUrl
	}
}
