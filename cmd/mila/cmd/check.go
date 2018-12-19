package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/julienbreux/mila/internal/mila/checker"
	"github.com/julienbreux/mila/internal/mila/reader"
)

const (
	defaultPathLongFlag  = "path"
	defaultPathShortFlag = "p"
	defaultPathValue     = "."
)

var (
	redString   = color.New(color.FgRed)
	greenString = color.New(color.FgGreen)
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check command is used to check environment variables.",
	Long:  "Check command is used to check environment variables.",
	Run: func(cmd *cobra.Command, args []string) {
		// Getting flag
		path, err := cmd.Flags().GetString(defaultPathLongFlag)
		if err != nil {
			failed("Unable to get schema filename flag\n\t%s", err)
			return
		}

		// Reading schema
		schema, err := reader.Reader(path)
		if err != nil {
			failed("%s", err)
			return
		}

		// Checking schema
		if err := checker.Checker(schema); err != nil {
			failed("%s", err)
			return
		}

		success("Environment is correct")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP(defaultPathLongFlag, defaultPathShortFlag, defaultPathValue, "Path of Mila schema")
}

func success(format string, a ...interface{}) {
	color.New(color.FgGreen).Printf("✓ %s\n", fmt.Sprintf(format, a...))
	os.Exit(0)
}

func failed(format string, a ...interface{}) {
	color.New(color.FgRed).Printf("✗ %s\n", fmt.Sprintf(format, a...))
	os.Exit(1)
}
