package themis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	restHost   string
	httpClient *http.Client
}

func NewClient(rest string) (*Client, error) {
	parsed, err := url.Parse(rest)
	if err != nil {
		return nil, fmt.Errorf("invalid rest server base url: %s", rest)
	}
	if parsed.Path != "/" {
		parsed.Path = ""
	}
	return &Client{
		restHost:   parsed.String(),
		httpClient: &http.Client{},
	}, nil
}

// ResponseWithHeight defines a response object type that wraps an original
// response with a height
type ResponseWithHeight struct {
	Height int64           `json:"height,string"`
	Result json.RawMessage `json:"result"`
}

// ClientError defines the attributes of a JSON error response
type ClientError struct {
	Path     string `json:"-"`
	HttpCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Err      string `json:"error"`
}

func (e ClientError) Error() string {
	return e.Err
}

func (e ClientError) NotFound() bool {
	return e.HttpCode == http.StatusNotFound
}

func (c *Client) Get(ctx context.Context, path string, result any) error {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, c.restHost+path, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data ResponseWithHeight
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		if err := json.Unmarshal(data.Result, result); err != nil {
			return err
		}
		return nil
	}

	var cerror ClientError
	if err := json.NewDecoder(resp.Body).Decode(&cerror); err != nil {
		return err
	}

	cerror.Path = path
	cerror.HttpCode = resp.StatusCode
	return cerror
}

func (c *Client) Post(ctx context.Context, path string, req, result any) error {
	reqdata, err := json.Marshal(req)
	if err != nil {
		return err
	}

	reqbody, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.restHost+path, bytes.NewReader(reqdata))
	if err != nil {
		return err
	}
	reqbody.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(reqbody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return err
		}
		return nil
	}

	var cerror ClientError
	if err := json.NewDecoder(resp.Body).Decode(&cerror); err != nil {
		return err
	}
	cerror.Path = path
	cerror.HttpCode = resp.StatusCode
	return cerror
}
