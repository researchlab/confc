package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "confc",
	Short: "confc is is a lightweight configuration management tool.",
	Long: `A tool which keeping local configuration files up-to-date using data stored in each env config file itself;
	              and is really convenience for multi-env configuration files generate.
                `,
	RunE: genConf,
}
