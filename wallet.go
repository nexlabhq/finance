package finance

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type WalletBalancesInput struct {
	UserID       string `json:"user_id,omitempty"`
	WalletTypeID string `json:"wallet_type_id,omitempty"`
	Promotion    *bool  `json:"promotion,omitempty"`
	Default      *bool  `json:"default,omitempty"`
}

type WalletBalance struct {
	UserID             string `graphql:"user_id"`
	WalletTypeID       string `graphql:"wallet_type_id"`
	Balance            int64  `graphql:"balance"`
	Amount             int64  `graphql:"amount"`
	HoldAmount         int64  `graphql:"hold_amount"`
	MinBalanceCapacity int64  `graphql:"min_balance_capacity"`
}

// GetWalletBalances get wallet balances of the user
func (c *Client) GetWalletBalances(input WalletBalancesInput) ([]WalletBalance, error) {
	var query struct {
		WalletBalances []WalletBalance `graphql:"walletBalances(data: $data)"`
	}

	variables := map[string]any{
		"data": input,
	}

	err := c.client.Query(context.Background(), &query, variables, graphql.OperationName("GetWalletBalances"))

	if err != nil {
		return nil, err
	}

	return query.WalletBalances, nil
}
