package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of confc",
	Long:  "All software has versions. This is confc's.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(ConfcVersion)
	},
}
var ConfcVersion string

func init() {
	RootCmd.AddCommand(versionCmd)
}
