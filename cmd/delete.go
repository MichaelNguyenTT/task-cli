package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteAllFlag bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Run $ task delete <taskid> to delete your desired task`,
	Run:   RunDeleteCmd,
}

func RunDeleteCmd(cmd *cobra.Command, args []string) {
	// handle target command flags
	if deleteAllFlag {
		file, err := loadFile("db.csv")
		if err != nil {
			log.Fatal(err)
		}

		// truncate the clear contents
		err = file.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}
		defer closeFile(file)

		fmt.Println("All your tasks have been deleted :)")
		return
	}
	taskID, _ := strconv.Atoi(args[0])
	file, err := loadFile("db.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(file)

	// read the contents of the file
	content, err := processCSVReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// the keys are the taskID
	mappedData := TaskMapper(content)

	// if key doesnt exist
	if _, ok := mappedData[taskID]; ok {
		fmt.Printf("Removed TaskID: %d", taskID)
	} else {
		log.Fatalf("TaskID: %d not found or out of bounds...", taskID)
	}

	delete(mappedData, taskID)
	overwriteFileContents(mappedData)
}

func overwriteFileContents(d map[int]Task) {
	s := make([]Task, 0, len(d))

	for _, task := range d {
		s = append(s, task)
	}

	file, err := os.Create("db.csv")
	if err != nil {
		log.Fatalf("Something went wrong overwriting file contents")
	}

	defer closeFile(file)

	processCSVWriter(file, s)
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVarP(&deleteAllFlag, "all", "a", false, "Delete all current tasks")

}
