package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Starts Cli management weekly",
	Run:   cliStart,
}

func cliStart(cmd *cobra.Command, args []string) {
	fmt.Println("Management Cli Weekly")
}
