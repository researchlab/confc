package main

import (
	"log"
	"runtime"

	"github.com/researchlab/confc/cmd"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Printf("confc %s (Git SHA: %s, Go Version: %s)\n", Version, GitSHA, runtime.Version())
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

}
