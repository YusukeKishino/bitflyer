package bitflyer

import (
	"fmt"
	"net/url"
)

type Execution struct {
	Id                         int     `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"`
	BuyChildOrderAcceptanceId  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceId string  `json:"sell_child_order_acceptance_id"`
}

const getExecutionsEndpoint = "/v1/executions"

// GetExecutions https://lightning.bitflyer.com/docs?lang=ja#%E7%B4%84%E5%AE%9A%E5%B1%A5%E6%AD%B4
func (c *Client) GetExecutions(productCode string, pagination ...PaginationParam) ([]Execution, error) {
	query := url.Values{}
	query.Set("product_code", productCode)
	if len(pagination) > 0 {
		param := pagination[0]
		if param.Count != 0 {
			query.Set("count", fmt.Sprintf("%d", param.Count))
		}
		if param.Before != 0 {
			query.Set("before", fmt.Sprintf("%d", param.Before))
		}
		if param.After != 0 {
			query.Set("after", fmt.Sprintf("%d", param.After))
		}
	}

	var executions []Execution
	err := c.getPublicEndpoint(getExecutionsEndpoint, query, &executions)
	if err != nil {
		return nil, err
	}
	return executions, nil
}
