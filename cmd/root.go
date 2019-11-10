package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cfgFile string

func initService() {
}

func init() {
	cobra.OnInitialize(initService)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file as toml (default is config.dev.toml)")
}

// RootCmd Process message from Kafka, gRPC, RestAPI
var RootCmd = &cobra.Command{
	Use:   "GO UTILS",
	Short: "Many of function can be used in compile with shell script way",
}

// Execute ...
func Execute() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(doCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Start program failed", zap.Error(err))
		os.Exit(-1)
	}
}
