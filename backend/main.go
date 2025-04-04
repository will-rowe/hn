package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/will-rowe/hn/backend/server"
)

var rootCmd = &cobra.Command{
	Use:   "report-service",
	Short: "Launch the ReportService gRPC and REST server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
