package cli

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print application version",
	Long:  "Prints application version, build time, os and architecture",
	Run: func(*cobra.Command, []string) {
	},
}
