package cli

import (
	"fmt"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var DeleteTaskCmd = &cobra.Command{
	Use:   "dt",
	Short: "Delete task",
	Run:   cliDeleteTask,
}

func cliDeleteTask(_ *cobra.Command, _ []string) {
	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	err = db.DeleteTask(ID)
	if err != nil {
		fmt.Println("error to delete task:", err)
	}

	fmt.Println("task deleted")
}
