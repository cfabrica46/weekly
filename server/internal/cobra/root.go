package cobra

import (
	"fmt"

	"server/internal/cobra/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "weekly"}

func Execute() (err error) {
	{
		cli.GetTasksTodayCmd.Flags().BoolVarP(&cli.DisplayerJSON, "json", "j", false, "format output to JSON")
		cli.GetTasksTodayCmd.Flags().BoolVarP(&cli.DisplayerXML, "xml", "x", false, "format output to XML")

		cli.GetTasksTodayCmd.MarkFlagsMutuallyExclusive(
			"json",
			"xml",
		)

		cli.CliCmd.AddCommand(cli.GetTasksTodayCmd)
	}

	{
		cli.GetTasksCmd.Flags().BoolVarP(&cli.DisplayerJSON, "json", "j", false, "format output to JSON")
		cli.GetTasksCmd.Flags().BoolVarP(&cli.DisplayerXML, "xml", "x", false, "format output to XML")

		cli.GetTasksCmd.MarkFlagsMutuallyExclusive(
			"json",
			"xml",
		)

		cli.CliCmd.AddCommand(cli.GetTasksCmd)
	}

	{
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Monday, "monday", "1", false, "set monday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Tuesday, "tuesday", "2", false, "set tuesday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Wednesday, "wednesday", "3", false, "set wednesday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Thursday, "thursday", "4", false, "set thursday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Friday, "friday", "5", false, "set friday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Saturday, "saturday", "6", false, "set saturday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.Sunday, "sunday", "7", false, "set sunday as an available day for the task")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.DisplayerJSON, "json", "j", false, "format output to JSON")
		cli.GetOneDayTasksCmd.Flags().BoolVarP(&cli.DisplayerXML, "xml", "x", false, "format output to XML")

		cli.GetTasksCmd.MarkFlagsMutuallyExclusive(
			"json",
			"xml",
		)

		cli.GetOneDayTasksCmd.MarkFlagsMutuallyExclusive(
			"monday",
			"tuesday",
			"wednesday",
			"thursday",
			"friday",
			"saturday",
			"sunday",
		)

		cli.CliCmd.AddCommand(cli.GetOneDayTasksCmd)
	}

	{
		cli.InsertTaskCmd.Flags().StringVarP(&cli.Title, "title", "t", "", "title of new task")
		cli.InsertTaskCmd.Flags().StringVarP(&cli.Description, "description", "d", "", "description of new task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Monday, "monday", "1", false, "set monday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Tuesday, "tuesday", "2", false, "set tuesday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Wednesday, "wednesday", "3", false, "set wednesday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Thursday, "thursday", "4", false, "set thursday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Friday, "friday", "5", false, "set friday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Saturday, "saturday", "6", false, "set saturday as an available day for the task")
		cli.InsertTaskCmd.Flags().BoolVarP(&cli.Sunday, "sunday", "7", false, "set sunday as an available day for the task")

		_ = cli.InsertTaskCmd.MarkFlagRequired("title")
		_ = cli.InsertTaskCmd.MarkFlagRequired("description")

		cli.CliCmd.AddCommand(cli.InsertTaskCmd)
	}

	//---
	{
		cli.ToggleDayCmd.Flags().IntVarP(&cli.ID, "id", "i", 0, "id to day of toggle")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Monday, "monday", "1", false, "set monday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Tuesday, "tuesday", "2", false, "set tuesday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Wednesday, "wednesday", "3", false, "set wednesday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Thursday, "thursday", "4", false, "set thursday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Friday, "friday", "5", false, "set friday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Saturday, "saturday", "6", false, "set saturday as an available day for the task")
		cli.ToggleDayCmd.Flags().BoolVarP(&cli.Sunday, "sunday", "7", false, "set sunday as an available day for the task")

		_ = cli.InsertTaskCmd.MarkFlagRequired("id")
		cli.GetOneDayTasksCmd.MarkFlagsMutuallyExclusive(
			"monday",
			"tuesday",
			"wednesday",
			"thursday",
			"friday",
			"saturday",
			"sunday",
		)

		cli.CliCmd.AddCommand(cli.ToggleDayCmd)
	}

	//---
	{
		cli.DeleteTaskCmd.Flags().IntVarP(&cli.ID, "id", "i", 0, "id to delete task")

		_ = cli.DeleteTaskCmd.MarkFlagRequired("id")

		cli.CliCmd.AddCommand(cli.DeleteTaskCmd)
	}

	rootCmd.AddCommand(cli.CliCmd)

	err = rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("error to execute cobra : %w", rootCmd.Execute())
	}

	return nil
}
