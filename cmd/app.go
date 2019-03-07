package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.Flags().StringVarP(&tmpl, "tmpl", "t", "", "configuration template file(s) path")
	RootCmd.Flags().StringVarP(&env, "env", "e", "", "configuration env, multi-env split by ','; default:dev")
	RootCmd.Flags().StringVarP(&dist, "dist", "d", "", "generate dist path; default:config/")
	RootCmd.Flags().StringVarP(&ptype, "ptype", "p", "json", "template parser type, support json,yaml,yml,toml; default:json")
}

func generate(cmd *cobra.Command, args []string) {
	log.Println("generate comming")
	// Do Stuff Here
}
