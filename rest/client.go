// This package provides a simple http.Client wrapper with json Encoder/Decoder useful for REST API interaction.
// The main object is `Client`, it can be created using `New()` method and at minimum only need `baseurl` parameter.
// Request can be made via `Call()` method.
package rest

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type JSONHTMLEscaper interface {
	JSONEscapeHTML() bool
}

// Client is a wrapper for http.Client
type Client struct {
	httpClient *http.Client
	baseUrl    string
	headers    map[string]string
}

// New return Client with defaults of 2 minutes timeout and skipping TLS verification.
func New(baseUrl string) *Client {
	c := &Client{
		&http.Client{
			Timeout: 2 * time.Minute,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		strings.TrimRight(baseUrl, "/"),
		map[string]string{"Content-Type": "application/json"},
	}
	return c
}

// Call helps make an http request given the request method, path & body and response object
// the response will be on `response` object if there's no error
// This is useful only if you don't need to set custom request headers or parse custom response header.
func (c *Client) Call(ctx context.Context, method string, path string, body interface{}, response interface{}) (err error) {
	var b io.Reader
	if body != nil {
		bb := &bytes.Buffer{}
		enc := json.NewEncoder(bb)
		if x, ok := body.(JSONHTMLEscaper); ok {
			enc.SetEscapeHTML(x.JSONEscapeHTML())
		}
		err := enc.Encode(body)
		if err != nil {
			return err
		}
		bb.Truncate(bb.Len() - 1)
		b = bb
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseUrl+path, b)
	if err != nil {
		return err
	}
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	res, err := c.httpClient.Do(req)
	if res != nil && res.Body != nil {
		defer func(c io.Closer) {
			err = c.Close()
		}(res.Body)
	}
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(&response)
}

func (c *Client) CallWithoutBody(ctx context.Context, method string, path string, response interface{}) (err error) {
	var b io.Reader

	req, err := http.NewRequestWithContext(ctx, method, c.baseUrl+path, b)
	if err != nil {
		return err
	}
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	res, err := c.httpClient.Do(req)
	if res != nil && res.Body != nil {
		defer func(c io.Closer) {
			err = c.Close()
		}(res.Body)
	}
	if err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(&response)
}
