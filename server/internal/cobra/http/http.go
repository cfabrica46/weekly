package http

import (
	"server/internal/http"

	"github.com/spf13/cobra"
)

var HTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http server for Weekly",
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}
