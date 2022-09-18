package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GetTasksCmd = &cobra.Command{
	Use:   "ts",
	Short: "Starts Cli management weekly",
	Run:   getTasks,
}

func getTasks(cmd *cobra.Command, args []string) {
	fmt.Println("get Tasks")
}
