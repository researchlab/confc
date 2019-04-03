package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/researchlab/confc/cmd"
)

func init() {
	cmd.ConfcVersion = fmt.Sprintf("confc version: %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
}
func main() {
	log.SetFlags(log.Ldate | log.Ltime) //| log.Llongfile)
	//log.Printf("confc version: %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

}
