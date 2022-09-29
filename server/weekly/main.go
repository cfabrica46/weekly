package main

import (
	"log"
	"os"

	"server/internal/cobra"
	"server/weekly/config"
)

func main() {
	if os.Getenv("IS_DOCKER") != "true" {
		if err := config.SetEnv(); err != nil {
			log.Fatal(err)
		}
	}

	if err := cobra.Execute(); err != nil {
		log.Fatal(err)
	}
}
