package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Starts Cli management weekly",
	Run:   cliStart,
}

func cliStart(cmd *cobra.Command, args []string) {
	fmt.Println("cli")
}
