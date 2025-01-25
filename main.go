package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Id     int
	Name   string
	Status string
}

func main() {
	tasksList, err := getTasks()

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

		tasksList = addTask(tasksList, taskName)
	case "list":
		if len(tasksList) == 0 {
			fmt.Println("empty task list")
			return
		}

		for _, task := range tasksList {
			fmt.Printf("%d - %s \n", task.Id, task.Name)
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

		tasksList = updateTaskName(tasksList, taskId, taskName)
	case "mark-done":
		if len(args) < 2 {
			return
		}

		taskId, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("an error has occurred ❌")
			return
		}

		tasksList = markAsDone(tasksList, taskId)
	case "mark-in-progress":
		if len(args) < 2 {
			return
		}

		taskId, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("an error has occurred ❌")
			return
		}

		tasksList = markAsInProgress(tasksList, taskId)
	}

	err = saveTasks(tasksList)

	if err != nil {
		fmt.Println("an error has occurred ❌")
		return
	}

	fmt.Println("tasks saved successfully ✅")
}

func getTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")

	if err != nil {
		return nil, err
	}

	var tasks []Task

	err = json.Unmarshal(data, &tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func markAsInProgress(tasksList []Task, taskId int) []Task {
	taskFound := getTask(tasksList, taskId)

	taskFound.Status = "in-progress"

	tasksList = updateTask(tasksList, taskFound)

	return tasksList
}

func markAsDone(tasksList []Task, taskId int) []Task {
	taskFound := getTask(tasksList, taskId)

	taskFound.Status = "done"

	tasksList = updateTask(tasksList, taskFound)

	return tasksList
}

func updateTaskName(tasksList []Task, taskId int, taskName string) []Task {
	taskFound := getTask(tasksList, taskId)

	taskFound.Name = taskName

	tasksList = updateTask(tasksList, taskFound)

	return tasksList
}

func updateTask(tasksList []Task, updatedTask Task) []Task {
	var taskIndex int

	for i, task := range tasksList {
		if task.Id == updatedTask.Id {
			taskIndex = i
			break
		}
	}

	tasksList[taskIndex] = updatedTask

	return tasksList

}

func getTask(tasksList []Task, taskId int) Task {
	var taskFound Task

	for _, task := range tasksList {
		if task.Id == taskId {
			taskFound = task
			break
		}
	}

	return taskFound
}

func addTask(tasksList []Task, taskName string) []Task {
	var newTaskId int

	if len(tasksList) == 0 {
		newTaskId = 1
	} else {
		lastTask := tasksList[len(tasksList)-1]
		newTaskId = lastTask.Id + 1
	}

	newTask := Task{
		Id:     newTaskId,
		Name:   taskName,
		Status: "todo",
	}

	tasksList = append(tasksList, newTask)

	return tasksList
}

func saveTasks(tasksList []Task) error {
	data, err := json.Marshal(tasksList)

	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}
