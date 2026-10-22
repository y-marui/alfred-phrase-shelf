// Command example-alfred is the binary the packaged Alfred Workflow invokes
// (see workflow/info.plist). Alfred's Script Filter node runs it with the
// query following the "example" keyword as $1.
//
// Replace this with your own binary: rename the cmd/ directory, swap
// internal/example + internal/examplecmd for your own domain logic, and
// keep internal/scriptfilter — it's Alfred-independent and reusable as-is.
package main

import (
	"fmt"
	"os"

	"github.com/y-marui/alfred-workflow-template/internal/examplecmd"
	"github.com/y-marui/alfred-workflow-template/internal/scriptfilter"
)

func main() {
	query := ""
	if len(os.Args) > 1 {
		query = os.Args[1]
	}
	writeResponse(dispatch(query))
}

// dispatch recovers from any panic in examplecmd: an unhandled failure must
// still produce a visible Script Filter error item rather than
// empty/invalid output.
func dispatch(query string) (resp scriptfilter.Response) {
	defer func() {
		if r := recover(); r != nil {
			resp = errorResponse(fmt.Sprintf("%v", r))
		}
	}()
	return examplecmd.List(query)
}

func writeResponse(resp scriptfilter.Response) {
	if err := resp.Write(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "example-alfred: writing response:", err)
		os.Exit(1)
	}
}

func errorResponse(message string) scriptfilter.Response {
	return scriptfilter.Response{
		Items: []scriptfilter.Item{
			{
				Title:    "Workflow Error",
				Subtitle: message,
				Arg:      message,
				Valid:    scriptfilter.BoolPtr(false),
			},
		},
	}
}
