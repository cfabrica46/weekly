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

func cliToggleDay(_ *cobra.Command, _ []string) {
	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	err = db.ToggleDayOfTask(1, mydb.Monday)
	if err != nil {
		fmt.Println("error to get all tasks:", err)
	}

	fmt.Println("day toggled")
}
