package bitflyer

import (
	"net/url"
)

type Bid struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Ask struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Bid   `json:"bids"`
	Asks     []Ask   `json:"asks"`
}

type State string

// https://lightning.bitflyer.com/docs?lang=ja#%E6%9D%BF%E3%81%AE%E7%8A%B6%E6%85%8B
const (
	Running      State = "RUNNING"
	Closed       State = "CLOSED"
	Starting     State = "STARTING"
	Preopen      State = "PREOPEN"
	CircuitBreak State = "CIRCUIT BREAK"
	AwaitingSQ   State = "AWAITING SQ"
	Matured      State = "MATURED"
)

type BoardState struct {
	Health string `json:"health"`
	State  string `json:"state"`
	Data   struct {
		SpecialQuotation int `json:"special_quotation"`
	} `json:"data"`
}

const getBoardEndpoint = "/v1/board"

// GetBoard https://lightning.bitflyer.com/docs?lang=ja#%E6%9D%BF%E6%83%85%E5%A0%B1
func (c *Client) GetBoard(productCode string) (Board, error) {
	query := url.Values{}
	query.Set("product_code", productCode)

	board := Board{}
	err := c.getPublicEndpoint(getBoardEndpoint, query, &board)
	if err != nil {
		return Board{}, err
	}
	return board, nil
}

const getBoardStateEndpoint = "/v1/getboardstate"

// GetBoardState https://lightning.bitflyer.com/docs?lang=ja#%E6%9D%BF%E3%81%AE%E7%8A%B6%E6%85%8B
func (c *Client) GetBoardState(productCode string) (BoardState, error) {
	query := url.Values{}
	query.Set("product_code", productCode)

	state := BoardState{}
	err := c.getPublicEndpoint(getBoardStateEndpoint, query, &state)
	if err != nil {
		return BoardState{}, err
	}
	return state, nil
}
