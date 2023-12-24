// Code generated by goa v3.14.1, DO NOT EDIT.
//
// transactions HTTP client types
//
// Command:
// $ goa gen djjuhasz/expenses/api/design --output api

package client

import (
	transactions "djjuhasz/expenses/api/gen/transactions"

	goa "goa.design/goa/v3/pkg"
)

// TransactionCollection is the type of the "transactions" service
// "transactions" endpoint HTTP response body.
type TransactionCollection []*Transaction

// Transaction is used to define fields on response body types.
type Transaction struct {
	// Identifier.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Account name.
	Account *string `form:"account,omitempty" json:"account,omitempty" xml:"account,omitempty"`
	// Date posted.
	DatePosted *string `form:"date_posted,omitempty" json:"date_posted,omitempty" xml:"date_posted,omitempty"`
	// Cheque number.
	ChequeNo *int `form:"cheque_no,omitempty" json:"cheque_no,omitempty" xml:"cheque_no,omitempty"`
	// Payee name.
	Payee *string `form:"payee,omitempty" json:"payee,omitempty" xml:"payee,omitempty"`
	// Category.
	Category *string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	// Notes.
	Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	// Dollar amount.
	Amount *float32 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// NewTransactionsResultOK builds a "transactions" service "transactions"
// endpoint result from a HTTP "OK" response.
func NewTransactionsResultOK(body TransactionCollection, marker *string) *transactions.TransactionsResult {
	v := make([]*transactions.Transaction, len(body))
	for i, val := range body {
		v[i] = unmarshalTransactionToTransactionsTransaction(val)
	}
	res := &transactions.TransactionsResult{
		Transactions: v,
	}
	res.Marker = marker

	return res
}

// ValidateTransactionCollection runs the validations defined on
// TransactionCollection
func ValidateTransactionCollection(body TransactionCollection) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateTransaction(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateTransaction runs the validations defined on Transaction
func ValidateTransaction(body *Transaction) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.DatePosted == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date_posted", "body"))
	}
	if body.Category == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("category", "body"))
	}
	if body.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("amount", "body"))
	}
	return
}
