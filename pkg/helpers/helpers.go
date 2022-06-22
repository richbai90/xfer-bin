package helpers

import (
	"github.com/spf13/cobra"
)

func FlagVal(val *string, flag string, cmd cobra.Command) bool {
	*val = cmd.Flag(flag).Value.String()
	return *val != ""
}