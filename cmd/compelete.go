package cmd

import (
	"github.com/spf13/cobra"
)

var compeleteCmd = &cobra.Command{
	Use:   "cmp",
	Short: "Mark a task as completed",
	Long:  `Target command to handle selecting a specific task and mark it as completed`,
	Run: cmpCommand,
}

//TODO: Still need to add logic for cmp command
func cmpCommand(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(compeleteCmd)
}
