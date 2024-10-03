package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// store for functions to be reused in other files

type Task struct {
	taskID      int
	description string
	date        string
	isCompleted bool
}

// opening the file
func loadFile(fPath string) (*os.File, error) {
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, fPath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// exclusive lock obtained on the filre descriptor
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		_ = file.Close()
		return nil, err
	}

	return file, nil
}

func closeFile(file *os.File) error {
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	return file.Close()
}

// file reader
func processCSVReader(file *os.File) ([][]string, error) {
	r := csv.NewReader(file)
	content, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return content, nil
}

// parser will return an array of []Task
func parser(data [][]string) []Task {
	parsingData := make([]Task, len(data))

	for i, line := range data {
		taskid, _ := strconv.Atoi(line[0])
		isCompleted, _ := strconv.ParseBool(line[3])
		parsingData[i] = Task{
			taskID:      taskid,
			description: line[1],
			date:        line[2],
			isCompleted: isCompleted,
		}
	}

	return parsingData
}

func TaskMapper(data [][]string) map[int]Task {
	mapper := make(map[int]Task, len(data))

	for _, line := range data {
		taskid, _ := strconv.Atoi(line[0])
		isCompleted, _ := strconv.ParseBool(line[3])
		task := Task{
			taskID:      taskid,
			description: line[1],
			date:        line[2],
			isCompleted: isCompleted,
		}

		mapper[taskid] = task
	}

	return mapper
}

// writing to the file
func processCSVWriter(file *os.File, records []Task) {
	w := csv.NewWriter(file)
	for _, record := range records {
		row := []string{strconv.Itoa(record.taskID), record.description, record.date, strconv.FormatBool(record.isCompleted)}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing new task to file", err)
		}
	}
	defer w.Flush()
}

// func DeleteTask(tasks map[int]Task, taskID int) []Task {
//
// 	if taskID <= 0 || taskID > len(tasks)-1 {
// 		fmt.Printf("TaskID: %d is out of bounds")
// 		return nil
// 	}
//
//
//
// 	return append(t[:idx], t[idx+1:]...)
// }
