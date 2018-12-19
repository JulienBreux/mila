package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	defaultSchemaFlag     = "schema"
	defaultSchemaFilename = ".mila.yaml"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check command is used to check environment variables.",
	Long:  "Check command is used to check environment variables.",
	Run: func(cmd *cobra.Command, args []string) {
		// Getting flag
		filename, err := cmd.Flags().GetString(defaultSchemaFlag)
		if err != nil {
			color.New(color.FgRed).
				Printf("✗ Unable to get schema filename flag\n")
			os.Exit(1)
			return
		}

		// Checking existing file
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			color.New(color.FgRed).
				Printf("✗ Schema filename \"%s\" not found\n", filename)
			os.Exit(1)
			return
		}

		fmt.Printf("check command called with filename: %s\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP(defaultSchemaFlag, "s", defaultSchemaFilename, "Filename of schema used to check environment variables")
}
