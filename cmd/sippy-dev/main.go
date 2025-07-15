package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logLevel = "info"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sippy-dev",
	Short: "Developer utilities for Sippy",
	Long: `Developer utilities for Sippy, including tools for populating test data
and other development workflow helpers.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level, err := log.ParseLevel(logLevel)
		if err != nil {
			log.WithError(err).Fatal("cannot parse log-level")
		}
		log.SetLevel(level)
		log.Debug("debug logging enabled")
	},
}

func main() {
	// Add some millisecond precision to log timestamps, useful for debugging performance.
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = "2006-01-02T15:04:05.999Z07:00"
	formatter.FullTimestamp = true
	formatter.DisableColors = false
	log.SetFormatter(formatter)

	rootCmd.AddCommand(
		NewSeedDataCommand(),
	)

	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info",
		"Log level (trace,debug,info,warn,error) (default info)")

	err := rootCmd.Execute()
	if err != nil {
		log.WithError(err).Fatal("could not execute root command")
	}
}
