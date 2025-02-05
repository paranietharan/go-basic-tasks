package tasks

import (
	"fmt"
)

type Task struct {
	taskName   string
	isFinished bool
}

func NewTask(taskName string) *Task {
	newTask := Task{
		taskName:   taskName,
		isFinished: false,
	}

	return &newTask
}

// array 2 store the tasks
var taskList []Task

func AddTask(taskName string) {
	taskList = append(taskList, *NewTask(taskName))
	fmt.Println("Task added: " + taskName)
}

func DeleteTask(index int) {
	if index < 0 || index > len(taskList) {
		fmt.Println("You entered invalide Index")
	}

	taskList = append(taskList[:index], taskList[index+1:]...)
	fmt.Println("Task deleted successfully!")
}

func ViewTask() {
	if len(taskList) == 0 {
		fmt.Println("Task is empty!")
	}
	fmt.Println("Tasks : ")
	for i, task := range taskList {
		status := "[\u2718]"
		if task.isFinished {
			status = "[\u2714]"
		}
		fmt.Printf("%d: %s %s\n", i+1, status, task.taskName)
	}
}

func MarkTaskComplete(index int) {
	if index < 0 || index > len(taskList) {
		fmt.Println("You entered invalide Index")
	}
	taskList[index].isFinished = true
	fmt.Println("Task marked as complete!")
}
