// Code generated by goa v3.14.1, DO NOT EDIT.
//
// transactions client
//
// Command:
// $ goa gen djjuhasz/expenses/api/design --output api

package transactions

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "transactions" service client.
type Client struct {
	TransactionsEndpoint goa.Endpoint
}

// NewClient initializes a "transactions" service client given the endpoints.
func NewClient(transactions goa.Endpoint) *Client {
	return &Client{
		TransactionsEndpoint: transactions,
	}
}

// Transactions calls the "transactions" endpoint of the "transactions" service.
func (c *Client) Transactions(ctx context.Context) (res *TransactionsResult, err error) {
	var ires any
	ires, err = c.TransactionsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TransactionsResult), nil
}
