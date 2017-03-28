package cli

import (
	"volume/app"
	"volume/config"
	"volume/log"

	"github.com/spf13/cobra"
)

// Commands
var rootCmd = &cobra.Command{
	Use:   "volume",
	Short: "SOON_ FM 2.0 Volume Control",
	Long:  "Allows real time control of the playback volume",
	Run: func(*cobra.Command, []string) {
		volume := app.New()
		if err := volume.Run(); err != nil {
			log.WithError(err).Error("application run error")
		}
	},
}

// CLI Initialisation
func init() {
	cobra.OnInitialize(
		config.ReadInConfig, // Read configuration file
		log.Initialize,      // Initialise the global logger
	)
	// Flags
	rootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"Absolute path to configuration file")
	rootCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Log Level (debug,info,warn,error)")
	// Bind Flags
	config.BindConfigPathFlag(rootCmd.PersistentFlags().Lookup("config"))
	log.BindLogLevelFlag(rootCmd.PersistentFlags().Lookup("log-level"))
}

// Execute CLI
func Execute() error {
	return rootCmd.Execute()
}
