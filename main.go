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

func getTasks() []Task {
	file, err := os.OpenFile(DB_FILE_PATH, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed to open or create file: %s", err)
	}
	defer file.Close()

	// Check if the file is empty and initialize it with an empty JSON array if it is
	info, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file stats: %s", err)
	}
	if info.Size() == 0 {
		_, err = file.WriteString("[]")
		if err != nil {
			log.Fatalf("Failed to initialize file: %s", err)
		}
	}

	// Read file
	data, err := os.ReadFile(DB_FILE_PATH)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Unmarshal the JSON data into the struct
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}
	return tasks
}

func saveTasks(tasks []Task) {
	// Marshal the struct back into JSON
	modifiedJSON, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	// Write the modified JSON back to the file
	if err := os.WriteFile(DB_FILE_PATH, modifiedJSON, 0644); err != nil {
		log.Fatalf("Error writing JSON to file: %s", err)
	}
}

func add(tasks *[]Task, title string, detail string) {
	maxID := 0
	for _, task := range *tasks {
		if task.Id > maxID {
			maxID = task.Id
		}
	}
	newID := maxID + 1

	newTask := Task{
		Id:     newID,
		Title:  title,
		Detail: detail,
	}
	*tasks = append(*tasks, newTask)
}

func remove(tasks *[]Task, taskId int) {
	for i, task := range *tasks {
		if task.Id == taskId {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return
		}
	}
}

func show(tasks []Task, verbose bool) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show.")
		return
	}

	for _, task := range tasks {
		if verbose {
			fmt.Printf("ID: %d, Title: %s, Detail: %s\n", task.Id, task.Title, task.Detail)
		} else {
			fmt.Printf("ID: %d, Title: %s\n", task.Id, task.Title)
		}
	}
}

func flush(tasks *[]Task) {
	*tasks = make([]Task, 0)
}

func main() {
	parser := argparse.NewParser("To-Do CLI App", "A CLI for a To-Do App")

	addMode := parser.NewCommand("add", "Add new task")
	addModeTaskTitle := addMode.String("t", "title", &argparse.Options{Required: true, Help: "Title for new task"})
	addModeTaskDetail := addMode.String("d", "details", &argparse.Options{Required: false, Help: "Details for new task", Default: ""})

	removeMode := parser.NewCommand("remove", "Remove a task using its id")
	removeModeTaskId := removeMode.Int("i", "id", &argparse.Options{Required: true, Help: "Id of Task to remove"})

	showMode := parser.NewCommand("show", "Show all tasks")
	showModeVerbose := showMode.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Enable verbose output",
		Default:  false,
	})

	flushMode := parser.NewCommand("flush", "Flush the database")

	// Parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	tasks := getTasks()

	switch {
	case addMode.Happened():
		add(&tasks, *addModeTaskTitle, *addModeTaskDetail)
		saveTasks(tasks)

	case removeMode.Happened():
		remove(&tasks, *removeModeTaskId)
		saveTasks(tasks)

	case showMode.Happened():
		show(tasks, *showModeVerbose)

	case flushMode.Happened():
		flush(&tasks)
		saveTasks(tasks)

	default:
		fmt.Println("No operation selected")
	}
}
