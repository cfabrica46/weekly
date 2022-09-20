package cli

import (
	"fmt"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var GetTasksCmd = &cobra.Command{
	Use:   "ts",
	Short: "Get all tasks",
	Run:   cliGetTasks,
}

func cliGetTasks(_ *cobra.Command, _ []string) {
	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	tasks, err := db.GetAllTasks()
	if err != nil {
		fmt.Println("error to get all tasks:", err)
	}

	fmt.Println(tasks)
}
