package cli

import (
	"fmt"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var GetOneDayTasksCmd = &cobra.Command{
	Use:   "tod",
	Short: "Get one-day tasks",
	Run:   cliGetOneDayTasks,
}

func cliGetOneDayTasks(cmd *cobra.Command, _ []string) {
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

	tasks, err := db.GetOneDayTask(dayOfWeek)
	if err != nil {
		fmt.Println("error to get all tasks:", err)

		return
	}

	fmt.Println(tasks)
}
