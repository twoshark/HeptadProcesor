package cmd

import (
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "manual",
		Short: "Run the Full Heptad Processor Utility",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infoln("Starting Heptad Processing Utility Application version: 0.1.0")
		},
	})
}
