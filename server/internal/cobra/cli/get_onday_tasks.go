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

func cliGetOneDayTasks(_ *cobra.Command, _ []string) {
	var disp displayer

	db, err := mydb.Open(mydb.DriverDefault, util.GetDBInfo())
	if err != nil {
		fmt.Print("error to connect database:", err)

		return
	}
	defer db.Close()

	switch {
	case DisplayerJSON:
		disp = getJSONDisplayer()
	case DisplayerXML:
		disp = getXMLDisplayer()
	default:
		disp = getDefaultDisplayer()
	}

	dayOfWeek, err := getDayOfWeek()
	if err != nil {
		fmt.Println("error to get day of week:", err)

		return
	}

	tasks, err := db.GetOneDayTask(dayOfWeek)
	if err != nil {
		fmt.Println("error to get all tasks:", err)

		return
	}

	resp, err := disp.display(tasks)
	if err != nil {
		fmt.Print("error display tasks:", err)

		return
	}

	fmt.Println(string(resp))
}
