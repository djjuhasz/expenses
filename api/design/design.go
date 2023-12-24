package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("expenses", func() {
	Title("Home expenses service")
	Description("Service for home expenses")

	Server("expenses", func() {
		// List the services hosted by this server.
		Services("transactions")

		Host("localhost", func() {
			URI("http://localhost:8000/")
		})
	})
})

var Transaction = ResultType("application/vnd.goa.djjuhasz.transaction", "Transaction", func() {
	Attribute("id", String, "Identifier.")
	Attribute("account", String, "Account name.")
	Attribute("date_posted", String, "Date posted.")
	Attribute("cheque_no", Int, "Cheque number.")
	Attribute("payee", String, "Payee name.")
	Attribute("category", String, "Category.")
	Attribute("notes", String, "Notes.")
	Attribute("amount", Float32, "Dollar amount.")

	Required("id", "date_posted", "category", "amount")
})

var Transactions = CollectionOf(Transaction)

var _ = Service("transactions", func() {
	Description("The transactions service performs operations on transactions.")

	Method("transactions", func() {
		Description("List all transactions.")
		Payload(func() {})

		Result(func() {
			Attribute("marker", String, "Pagination marker")
			Attribute("transactions", Transactions, "list of transactions")
		})

		HTTP(func() {
			GET("/transactions")
			Response(StatusOK, func() {
				Header("marker")
				Body("transactions")
			})
		})

	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
