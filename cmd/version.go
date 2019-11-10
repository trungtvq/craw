package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

// SetRevision inject version from git
func SetRevision(r string) {
	if len(r) > 0 {
		version = fmt.Sprintf("%v-%v", version, r)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Transfer Promotion version -- %v\n", version)
	},
}
