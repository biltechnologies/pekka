package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func init() {
	version = "0.0.2"
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pekka %s\n", version)
	},
}
