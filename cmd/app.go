package cmd

import (
	"log"

	"github.com/researchlab/confc/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.Flags().StringVarP(&tmpl, "tmpl", "t", "", "configuration template file(s) path")
	RootCmd.Flags().StringVarP(&env, "env", "e", "dev", "configuration env, multi-env split by ','; eg. dev, prod, test")
	RootCmd.Flags().StringVarP(&dist, "dist", "d", "config/", "generate dist path")
	RootCmd.Flags().StringVarP(&ptype, "ptype", "p", "json", "template parser type; support-list: json, yaml, yml, toml")
}

func generate(cmd *cobra.Command, args []string) {
	log.Println("generate comming")
	lib.Show()
	// Do Stuff Here
}
