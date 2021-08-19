package bitflyer

import (
	"fmt"
	"net/url"
	"time"
)

type Chat struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

const getChatsEndpoint = "getchats"

func (c *Client) GetChats(fromDate ...time.Time) ([]Chat, error) {
	query := url.Values{}
	if len(fromDate) > 0 {
		date := fromDate[0]
		query.Set("from_date", date.Format("2006-01-02T15:04:05.000"))
	}
	fmt.Println(query)

	var chats []Chat
	err := c.getPublicEndpoint(getChatsEndpoint, query, &chats)
	if err != nil {
		return nil, err
	}
	return chats, nil
}
