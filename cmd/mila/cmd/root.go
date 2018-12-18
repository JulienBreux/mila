package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mila",
	Short: "Mila helps you manage your environment variables",
	Run:   runRootCommand,
}

// Execute command system
func Execute(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runRootCommand is default command
func runRootCommand(cmd *cobra.Command, args []string) {
	cmd.Help()
}
