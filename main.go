package main

import (
	"log"
	"os"

	"github.com/your-username/lifebar/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
