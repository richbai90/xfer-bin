package xfer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

type Inspect struct {
	CreatedAt  string
	Driver     string
	Labels     map[string]string
	Mountpoint string
	Name       string
	Options    map[string]string
	Scope      string
}

// TODO: This is so simple as to not merit its own executable. Refactor to include as part of the bundle
func Restore(src *string, dest *string, _debug *bool) func(cmd *cobra.Command, args []string) error {
	handler := ErrorHandler{Debug: _debug}
	return func(cmd *cobra.Command, args []string) error {
		debug := *_debug
		info, err := os.Stat(*src)
		if handler.HandleErr(&err, "Unable to stat source file %s ", *src) {
			return err
		}

		if debug {
			fmt.Fprintln(os.Stderr, "Source File Found: ")
			if j, err := json.MarshalIndent(&info, "", "\t"); err == nil {
				fmt.Fprintln(os.Stderr, string(j))
			}
		}

		oscmd := exec.Command("tar", "-zxvf", *src)
		oscmd.Dir = *dest
		oscmd.Run()

		return nil
	}
}
