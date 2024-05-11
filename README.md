# To-Do CLI Tool in Go

This is a simple command line interface tool I made in Go in order to unserstand how CLI tools work. The data is stored in a local json file.

Install using
```
git clone https://github.com/loucaspapalazarou/todo-go-cli.git
cd todo-go-cli
go mod tidy
```

Compile and run using
```
go run main.go
```

The tool supports 4 commands
```
  add     Add new task
  remove  Remove a task using its id
  show    Show all tasks
  flush   Flush the database
```

Each command's specific flags can be seen by running the mode and help like so
```
go run main.go <command> -h
```

Example usage
```
go run main.go add -t "shower" -d "take a shower you smelly programmer"
go run main.go add -t "clean basement"
go run main.go show -v

ID: 1, Title: shower, Detail: take a shower you smelly programmer
ID: 2, Title: clean basement, Detail: None

go run main.go remove -i 2

ID: 1, Title: shower, Detail: take a shower you smelly programmer
```

Instead of using `go run` every time, we can compile the code using `go build main.go` and call the executable using `./main`.
