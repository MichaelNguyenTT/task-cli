package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "get the list of current tasks",
	Long: `Example command to get all the list of tasks:

		$ tasks list`,
	Run: RunListCmd,
}

func RunListCmd(cmd *cobra.Command, args []string) {

	file, err := loadFile("db.csv")
	if err != nil {
		fmt.Printf("Something went wrong during loading file: %s", err)
	}

	defer closeFile(file)
	// read the file content
	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal("failure to read the file", err)
	}

	tb := createTableData()
	parsedData := parser(data)
	tableData := appendData(tb, parsedData)

	pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(tableData).Render()

}

func appendData(tb [][]string, data []Task) [][]string {
	for _, content := range data {
		tb = append(tb, []string{strconv.Itoa(content.taskID), content.description, content.date, strconv.FormatBool(content.isCompleted)})
	}
	return tb
}

func createTableData() pterm.TableData {
	tableData := pterm.TableData{
		{"TaskID", "Description", "CreatedAt", "Completed"},
	}

	return tableData
}

func init() {
	rootCmd.AddCommand(listCmd)
}
