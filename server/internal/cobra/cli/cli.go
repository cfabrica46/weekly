package cli

import (
	"fmt"

	mydb "server/internal/db"

	"github.com/spf13/cobra"
)

var (
	ID int

	Title       string
	Description string

	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool

	DisplayerJSON bool
	DisplayerXML  bool
)

var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Starts Cli management weekly",
}

func getDayOfWeek() (dayOfWeek mydb.DayOfWeek, err error) {
	switch {
	case Monday:
		dayOfWeek = mydb.Monday
	case Tuesday:
		dayOfWeek = mydb.Tuesday
	case Wednesday:
		dayOfWeek = mydb.Wednesday
	case Thursday:
		dayOfWeek = mydb.Thursday
	case Friday:
		dayOfWeek = mydb.Friday
	case Saturday:
		dayOfWeek = mydb.Saturday
	case Sunday:
		dayOfWeek = mydb.Sunday
	default:
		return "", fmt.Errorf("error needs at least one flag")
	}

	return dayOfWeek, nil
}
