package bitflyer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultBaseURL = "https://api.bitflyer.com"
)

var defaultHttpClient = &http.Client{}

type PaginationParam struct {
	Count  uint
	Before uint
	After  uint
}

type Client struct {
	client       *http.Client
	baseURL      *url.URL
	accessKey    string
	accessSecret string
}

func DefaultClient() *Client {
	key, secret := readSecrets()
	return NewClient(key, secret)
}

// NewClient with bitFlyer API AccessKey and AccessSecret
// see details https://lightning.bitflyer.com/docs?lang=ja#%E8%AA%8D%E8%A8%BC
func NewClient(accessKey, accessSecret string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:       defaultHttpClient,
		baseURL:      baseURL,
		accessKey:    accessKey,
		accessSecret: accessSecret,
	}
}

func readSecrets() (accessKey, accessSecret string) {
	accessKey = os.Getenv("BITFLYER_ACCESS_KEY")
	accessSecret = os.Getenv("BITFLYER_ACCESS_SECRET")
	return
}

func (c *Client) String() string {
	s := struct {
		BaseURL string `json:"base_url"`
	}{
		BaseURL: c.baseURL.String(),
	}
	marshal, _ := json.Marshal(s)
	return string(marshal)
}

func (c *Client) getPublicEndpoint(endpoint string, query url.Values, dest interface{}) error {
	u, _ := url.Parse(c.baseURL.String() + endpoint)
	u.RawQuery = query.Encode()

	res, err := c.client.Get(u.String())
	if err != nil {
		return fmt.Errorf("GET %s: %w", endpoint, err)
	}
	if res.StatusCode != http.StatusOK {
		errMessage, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("read error response body: %w", err)
		}
		return fmt.Errorf("failed to GET %s, got status: %d, response: %s", endpoint, res.StatusCode, string(errMessage))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("unmarshal response body: %w", err)
	}
	return nil
}
