package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "CLI task app to add, delete, complete",
	Long: `A tasks application using the cobra command line interface. Heres the list of commands:

	add <description> -Add a new task followed by the description
	delete <taskid> - Delete the task from the data store
	complete <taskid> - mark a task as done
	list - get the list of current tasks`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
