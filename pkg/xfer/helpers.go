package xfer

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type ErrorHandler struct {
	Debug *bool
}

func FlagVal(val *string, flag string, cmd cobra.Command) bool {
	*val = cmd.Flag(flag).Value.String()
	return *val != ""
}

// Wraps the error in a stack trace if the error is not = nil
// Return true if an error occurred and false otherwise
func (eh *ErrorHandler) HandleErr(err *error, args ...interface{}) bool {
	if len(args) == 0 {
		args = append(args, "")
	}
	e := errors.Wrap(*err, fmt.Sprintf(args[0].(string), args...))

	if *(eh.Debug) && e != nil {
		fmt.Println(e.Error())
		*err = e
		return true
	}

	return false
}
