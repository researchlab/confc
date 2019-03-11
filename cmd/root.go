package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "confc-cli",
	Short: "confc-cli is is a lightweight configuration management tool.",
	Long: `A tool which keeping local configuration files up-to-date using data stored in each env config file itself;
	              is really convenience for multi-env config file generate.
                `,
	RunE: genConf,
}
