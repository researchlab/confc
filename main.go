package main

import (
	"log"

	"github.com/researchlab/confc/cmd"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

}
