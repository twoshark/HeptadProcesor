package cmd

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "web",
		Short: "Run the Heptad Processor Utility UI",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infoln("Starting Heptad Processing Utility Application version: 0.1.0")
			serve()
		},
	})
}

func serve() {
	//TODO: Serve an app
}
