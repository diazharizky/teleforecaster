package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"time"

	urlutils "net/url"
)

type Client struct {
	baseURL string
	headers map[string]string
	agent   *http.Client
	apiName string
}

type ClientConfig struct {
	BaseURL string
	Headers map[string]string
	Timeout time.Duration
	APIName string
}

func New(cfg ClientConfig) Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 5 * time.Second
	}

	client := Client{
		baseURL: cfg.BaseURL,
		headers: cfg.Headers,
		agent: &http.Client{
			Timeout: cfg.Timeout,
		},
		apiName: cfg.APIName,
	}

	return client
}

func (h Client) Get(path string, params map[string]string) (*Response, error) {
	return h.sendRequest(http.MethodGet, path, params, nil)
}

func (c Client) sendRequest(method, path string, params map[string]string, body io.Reader) (*Response, error) {
	url := c.baseURL + path

	if len(params) > 0 {
		qs := urlutils.Values{}
		for key, val := range params {
			qs.Add(key, val)
		}

		url += fmt.Sprintf("?%s", qs.Encode())
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if len(c.headers) > 0 {
		for key, val := range c.headers {
			req.Header.Add(key, val)
		}
	}

	resp, err := c.agent.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       resp.Body,
	}, nil
}
