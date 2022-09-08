package main

import (
	"log"
	"os"

	"server/cmd/config"
	"server/internal/cobra"
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
