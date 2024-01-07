package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

var DB_FILE_PATH = "db.json"

// Define a struct that matches your JSON structure.
type Task struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func readData() []Task {
	file, err := os.OpenFile(DB_FILE_PATH, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed to open or create file: %s", err)
	}
	defer file.Close()

	// Read file
	data, err := os.ReadFile(DB_FILE_PATH)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Unmarshal the JSON data into the struct
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Fatal(err)
	}
	return tasks
}

func saveData(tasks []Task) {
	// // Marshal the struct back into JSON
	// modifiedJSON, err := json.Marshal(person)
	// if err != nil {
	// 	log.Fatalf("Error marshalling JSON: %s", err)
	// }

	// // Write the modified JSON back to the file
	// if err := os.WriteFile(DB_FILE_PATH, modifiedJSON, os.ModePerm); err != nil {
	// 	log.Fatalf("Error writing JSON to file: %s", err)
	// }

	// // Output success message
	// log.Println("JSON file updated successfully")
}

func add(tasks []Task, newTask Task) {
	// assign new id
}

func remove(tasks []Task, taskId int) {}

func show(tasks []Task, verbose bool) {}

func main() {
	parser := argparse.NewParser("To-Do CLI App", "A CLI for a To-Do APP")

	addMode := parser.NewCommand("add", "Add")
	addModeTaskTitle := addMode.String("t", "title", &argparse.Options{Required: true, Help: "Title for new task"})
	addModeTaskDetail := addMode.String("d", "details", &argparse.Options{Required: false, Help: "Details for new task", Default: ""})

	removeMode := parser.NewCommand("remove", "Remove")
	removeModeTaskId := removeMode.Int("i", "id", &argparse.Options{Required: true, Help: "Id of Task to remove"})

	showMode := parser.NewCommand("show", "Show")
	showModeVerbose := showMode.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Enable verbose output",
		Default:  false,
	})

	// Parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	tasks := readData()

	switch {
	case addMode.Happened():
		add(tasks, Task{Id: -1, Title: *addModeTaskTitle, Detail: *addModeTaskDetail})

	case removeMode.Happened():
		remove(tasks, *removeModeTaskId)

	case showMode.Happened():
		show(tasks, *showModeVerbose)

	default:
		fmt.Println("No operation selected")
	}
}
