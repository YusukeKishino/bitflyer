package bitflyer

import "net/url"

type Health string

// https://lightning.bitflyer.com/docs?lang=ja#%E6%9D%BF%E3%81%AE%E7%8A%B6%E6%85%8B
const (
	Normal    Health = "NORMAL"
	Busy      Health = "BUSY"
	VeryBusy  Health = "VERY BUSY"
	SuperBusy Health = "SUPER BUSY"
	NoOrder   Health = "NO ORDER"
	Stop      Health = "STOP"
)

const getHealthEndpoint = "/v1/gethealth"

func (c *Client) GetHealth(productCode string) (Health, error) {
	query := url.Values{}
	query.Set("product_code", productCode)

	health := struct {
		Status Health `json:"status"`
	}{}
	err := c.getPublicEndpoint(getHealthEndpoint, query, &health)
	if err != nil {
		return "", err
	}
	return health.Status, nil
}
