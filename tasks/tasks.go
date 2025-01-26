package tasks

import (
	"encoding/json"
	"os"
)

type Task struct {
	Id     int
	Name   string
	Status string
}

func GetTasks() ([]Task, error) {
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

func MarkAsInProgress(tasksList []Task, taskId int) []Task {
	taskFound := GetTask(tasksList, taskId)

	taskFound.Status = "in-progress"

	tasksList = UpdateTask(tasksList, taskFound)

	return tasksList
}

func MarkAsDone(tasksList []Task, taskId int) []Task {
	taskFound := GetTask(tasksList, taskId)

	taskFound.Status = "done"

	tasksList = UpdateTask(tasksList, taskFound)

	return tasksList
}

func UpdateTaskName(tasksList []Task, taskId int, taskName string) []Task {
	taskFound := GetTask(tasksList, taskId)

	taskFound.Name = taskName

	tasksList = UpdateTask(tasksList, taskFound)

	return tasksList
}

func UpdateTask(tasksList []Task, updatedTask Task) []Task {
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

func GetTask(tasksList []Task, taskId int) Task {
	var taskFound Task

	for _, task := range tasksList {
		if task.Id == taskId {
			taskFound = task
			break
		}
	}

	return taskFound
}

func AddTask(tasksList []Task, taskName string) []Task {
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

func SaveTasks(tasksList []Task) error {
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
