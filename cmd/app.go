package cmd

import (
	"github.com/researchlab/confc/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.Flags().StringVarP(&tmpl, "tmpl", "t", "", "configuration template file(s) path")
	RootCmd.Flags().StringVarP(&env, "env", "e", "dev", "configuration env, multi-env split by ','; eg. dev, prod, test")
	RootCmd.Flags().StringVarP(&dist, "dist", "d", "config/", "generate dist path")
	RootCmd.Flags().StringVarP(&ptype, "ptype", "p", "json", "template parser type; support-list: json, yaml, yml, toml")
	RootCmd.Flags().StringVarP(&cache, "cache", "c", "on", "cache cmdline parameters into the system environment variable; on: cache, off: no cache, flush: flush cache")
}

func genConf(cmd *cobra.Command, args []string) error {
	var genConfCtx = &lib.GenConfCtx{
		Tmpl:  tmpl,
		Env:   env,
		Dist:  dist,
		Ptype: ptype,
		Cache: cache,
	}
	err := lib.Assembly(genConfCtx)
	if err != nil {
		return err
	}
	return nil
	// Do Stuff Here
}
