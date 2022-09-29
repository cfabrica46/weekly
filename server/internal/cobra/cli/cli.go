package cli

import (
	"fmt"

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
	Run:   cliCMD,
}

func cliCMD(_ *cobra.Command, _ []string) {
	var optionMenu int
	var exit bool

	for !exit {
		displayMenu()

		fmt.Print("> ")
		fmt.Scan(&optionMenu)
		fmt.Println()

		if optionMenu >= 0 && optionMenu <= 5 {
			exit = true
		}

		switch optionMenu {
		case 1:
			menuDisplayer()

			cliGetTasks(nil, nil)

		case 2:
			menuDaysOfWeek()

			menuDisplayer()

			cliGetOneDayTasks(nil, nil)

		case 3:
			menuDaysOfWeek()

			fmt.Print("Title: ")
			fmt.Scan(&Title)
			fmt.Print("Description: ")
			fmt.Scan(&Description)

			cliInsertTask(nil, nil)

		case 4:
			fmt.Print("ID: ")
			fmt.Scan(&ID)

			menuDaysOfWeek()

			cliToggleDay(nil, nil)

		case 5:
			fmt.Print("ID: ")
			fmt.Scan(&ID)

			cliDeleteTask(nil, nil)
		case 0:
			fmt.Println("Thaks! 😜")
			return

		default:
			fmt.Println("Please select a valid option!!! 😠")
			fmt.Println()
		}
	}
}
