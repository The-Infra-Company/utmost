// Package cmd acts as the root of the command tree.
package cmd

import (
	"os"

	i "github.com/The-Infra-Company/utmost/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "utmost",
	Short:        "My Utmost for His Highest in the CLI",
	Long:         "A utility for bringing Oswald Chambers' My Utmost for His Highest to the CLI",
	RunE:         i.FetchDevotional,
	SilenceUsage: true,
}

func init() {
	// Disable auto-generated help and usage strings
	rootCmd.DisableAutoGenTag = true
}

func Execute() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
