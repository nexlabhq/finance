package finance

import (
	"github.com/hgiasac/graphql-utils/client"
)

// Client represents a graphql client with payment gateway actions
type Client struct {
	client client.Client
}

// NewClient creates the finance client instance
func NewClient(client client.Client) *Client {
	return &Client{
		client: client,
	}
}
