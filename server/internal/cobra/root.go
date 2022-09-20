package cobra

import (
	"fmt"

	"server/internal/cobra/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "weekly"}

func Execute() (err error) {
	cli.CliCmd.AddCommand(cli.GetTasksCmd)

	rootCmd.AddCommand(cli.CliCmd)

	err = rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("error to execute cobra : %w", rootCmd.Execute())
	}

	return nil
}
