package bitflyer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Market struct {
	ProductCode string `json:"product_code"`
	Alias       string `json:"alias"`
	MarketType  string `json:"market_type"`
}

const getMarketsEndpoint = "/v1/markets"

// GetMarkets https://lightning.bitflyer.com/docs?lang=ja#%E3%83%9E%E3%83%BC%E3%82%B1%E3%83%83%E3%83%88%E3%81%AE%E4%B8%80%E8%A6%A7
func (c *Client) GetMarkets() ([]Market, error) {
	res, err := c.client.Get(c.baseURL.String() + getMarketsEndpoint)
	if err != nil {
		return nil, fmt.Errorf("get markets: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errMessage, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("read error response body: %w", err)
		}
		return nil, fmt.Errorf("failed to get markets, got status: %d, response: %s", res.StatusCode, string(errMessage))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, fmt.Errorf("unmarshal response body: %w", err)
	}
	return markets, nil
}
