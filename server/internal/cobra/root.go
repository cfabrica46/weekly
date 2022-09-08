package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "weekly"}

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(cliCmd)
}

func Execute() (err error) {
	return fmt.Errorf("error to execute cobra : %w", rootCmd.Execute())
}
