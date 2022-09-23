package cli

import (
	"fmt"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var ToggleDayCmd = &cobra.Command{
	Use:   "td",
	Short: "Toogle day of a Task",
	Run:   cliToggleDay,
}

func cliToggleDay(cmd *cobra.Command, _ []string) {
	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	dayOfWeek, err := getDayOfWeek()
	if err != nil {
		_ = cmd.Help()

		return
	}

	err = db.ToggleDayOfTask(ID, dayOfWeek)
	if err != nil {
		fmt.Println("error to toggle task:", err)
	}

	fmt.Println("day toggled")
}
