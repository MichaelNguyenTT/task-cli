package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add <description> command will create a new task",
	Long: `How to use:

	tasks add "example task"`,
	Run: AddTask,
}

func AddTask(cmd *cobra.Command, args []string) {
	description := make([]string, len(args))
	// store the description in a slice
	for i, arg := range args {
		description[i] = arg
	}
	desc := strings.Join(description, " ")

	// load file
	file, err := loadFile("db.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(file)

	content, err := processCSVReader(file)
	if err != nil {
		log.Fatalln("Error reading the CSV file", err)
	}

	mappedTasks := TaskMapper(content)

	newTaskID := len(mappedTasks) + 1

	// DATE
	// TODO: Take current time and create a diff timer for when its created
	createdAt := time.Now().Local().Format(time.UnixDate)

	defaultCompletion := false

	records := []Task{
		{newTaskID, desc, createdAt, defaultCompletion},
	}

	processCSVWriter(file, records)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
