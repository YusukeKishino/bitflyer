package bitflyer

import "net/url"

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	State           State   `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickId          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

const getTickerEndpoint = "/v1/ticker"

// GetTicker https://lightning.bitflyer.com/docs?lang=ja#ticker
func (c *Client) GetTicker(productCode string) (Ticker, error) {
	query := url.Values{}
	query.Set("product_code", productCode)

	ticker := Ticker{}
	err := c.getPublicEndpoint(getTickerEndpoint, query, &ticker)
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}
