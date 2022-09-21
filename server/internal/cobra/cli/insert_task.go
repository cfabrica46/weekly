package cli

import (
	"fmt"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var InsertTaskCmd = &cobra.Command{
	Use:   "it",
	Short: "Insert task",
	Run:   cliInsertTask,
}

func cliInsertTask(_ *cobra.Command, _ []string) {
	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	daysOfWeek := mydb.GetDaysOfWeek(true, false, false, false, false, false, false)

	err = db.InsertTask(mydb.GetTask("weekly", "do weekly", daysOfWeek))
	if err != nil {
		fmt.Println("error to insert task:", err)
	}

	fmt.Println("day inserted")
}
