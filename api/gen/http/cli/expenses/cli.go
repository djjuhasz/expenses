// Code generated by goa v3.14.1, DO NOT EDIT.
//
// expenses HTTP client CLI support package
//
// Command:
// $ goa gen djjuhasz/expenses/api/design --output api

package cli

import (
	transactionsc "djjuhasz/expenses/api/gen/http/transactions/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `transactions transactions
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` transactions transactions` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		transactionsFlags = flag.NewFlagSet("transactions", flag.ContinueOnError)

		transactionsTransactionsFlags = flag.NewFlagSet("transactions", flag.ExitOnError)
	)
	transactionsFlags.Usage = transactionsUsage
	transactionsTransactionsFlags.Usage = transactionsTransactionsUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "transactions":
			svcf = transactionsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "transactions":
			switch epn {
			case "transactions":
				epf = transactionsTransactionsFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "transactions":
			c := transactionsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "transactions":
				endpoint = c.Transactions()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// transactionsUsage displays the usage of the transactions command and its
// subcommands.
func transactionsUsage() {
	fmt.Fprintf(os.Stderr, `The transactions service performs operations on transactions.
Usage:
    %[1]s [globalflags] transactions COMMAND [flags]

COMMAND:
    transactions: List all transactions.

Additional help:
    %[1]s transactions COMMAND --help
`, os.Args[0])
}
func transactionsTransactionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] transactions transactions

List all transactions.

Example:
    %[1]s transactions transactions
`, os.Args[0])
}