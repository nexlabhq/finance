package finance

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

// CreatePayment creates a payment action request to external payment gateways
func (c *Client) CreatePayment(input CreatePaymentInput) (*CreatePaymentOutput, error) {
	var mutation struct {
		CreatePayment CreatePaymentOutput `graphql:"createPayment(data: $data)"`
	}

	variables := map[string]any{
		"data": input,
	}

	err := c.client.Mutate(context.TODO(), &mutation, variables, graphql.OperationName("CreatePayment"))
	if err != nil {
		return nil, err
	}
	return &mutation.CreatePayment, nil
}

// Refund creates a refund action request to external payment gateways
func (c *Client) Refund(input RefundPaymentInput) (*RefundPaymentOutput, error) {

	var mutation struct {
		RefundPayment RefundPaymentOutput `graphql:"refundPayment(data: $data)"`
	}

	variables := map[string]any{
		"data": input,
	}

	err := c.client.Mutate(context.TODO(), &mutation, variables, graphql.OperationName("CreatePayment"))
	if err != nil {
		return nil, err
	}
	return &mutation.RefundPayment, nil
}
