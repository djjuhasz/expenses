// Code generated by goa v3.14.1, DO NOT EDIT.
//
// transactions HTTP client encoders and decoders
//
// Command:
// $ goa gen djjuhasz/expenses/api/design --output api

package client

import (
	"bytes"
	"context"
	transactions "djjuhasz/expenses/api/gen/transactions"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildTransactionsRequest instantiates a HTTP request object with method and
// path set to call the "transactions" service "transactions" endpoint
func (c *Client) BuildTransactionsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: TransactionsTransactionsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("transactions", "transactions", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeTransactionsResponse returns a decoder for responses returned by the
// transactions transactions endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeTransactionsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body TransactionCollection
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("transactions", "transactions", err)
			}
			err = ValidateTransactionCollection(body)
			if err != nil {
				return nil, goahttp.ErrValidationError("transactions", "transactions", err)
			}
			var (
				marker *string
			)
			markerRaw := resp.Header.Get("Marker")
			if markerRaw != "" {
				marker = &markerRaw
			}
			res := NewTransactionsResultOK(body, marker)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("transactions", "transactions", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTransactionToTransactionsTransaction builds a value of type
// *transactions.Transaction from a value of type *Transaction.
func unmarshalTransactionToTransactionsTransaction(v *Transaction) *transactions.Transaction {
	res := &transactions.Transaction{
		ID:         *v.ID,
		Account:    v.Account,
		DatePosted: *v.DatePosted,
		ChequeNo:   v.ChequeNo,
		Payee:      v.Payee,
		Category:   *v.Category,
		Notes:      v.Notes,
		Amount:     *v.Amount,
	}

	return res
}
