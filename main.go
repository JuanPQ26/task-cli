package main

import (
	"fmt"
	"os"
	"strconv"
	"task-cli/tasks"
)

func main() {
	tasksList, err := tasks.GetTasks()

	if err != nil {
		fmt.Println("an error has occurred ❌")
		return
	}

	args := os.Args[1:]

	if len(args) < 1 {
		return
	}

	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			return
		}

		taskName := args[1]

		tasksList = tasks.AddTask(tasksList, taskName)
	case "list":
		if len(tasksList) == 0 {
			fmt.Println("empty task list")
			return
		}

		for _, task := range tasksList {
			fmt.Printf("%d - %s - %s \n", task.Id, task.Name, task.Status)
		}

		return

	case "update":
		if len(args) < 3 {
			return
		}

		taskId, err := strconv.Atoi(args[1])
		taskName := args[2]

		if err != nil {
			fmt.Println("an error has occurred ❌")
			return
		}

		tasksList = tasks.UpdateTaskName(tasksList, taskId, taskName)
	case "mark-done":
		if len(args) < 2 {
			return
		}

		taskId, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("an error has occurred ❌")
			return
		}

		tasksList = tasks.MarkAsDone(tasksList, taskId)
	case "mark-in-progress":
		if len(args) < 2 {
			return
		}

		taskId, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("an error has occurred ❌")
			return
		}

		tasksList = tasks.MarkAsInProgress(tasksList, taskId)
	}

	err = tasks.SaveTasks(tasksList)

	if err != nil {
		fmt.Println("an error has occurred ❌")
		return
	}

	fmt.Println("tasks saved successfully ✅")
}
