package cli

import (
	"fmt"
	"strings"
	"time"

	mydb "server/internal/db"
	"server/internal/util"

	"github.com/spf13/cobra"
)

var GetTasksTodayCmd = &cobra.Command{
	Use:   "tt",
	Short: "Get one-day tasks",
	Run:   cliGetTasksToday,
}

func cliGetTasksToday(_ *cobra.Command, _ []string) {
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

	tasks, err := db.GetOneDayTask(mydb.DayOfWeek(strings.ToLower(time.Now().Weekday().String())))
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
