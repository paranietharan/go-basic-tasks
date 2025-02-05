package main

import (
	"bufio"
	"fmt"
	"os"
	"todolist/tasks"
)

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	// infinite loop
	i := 1
	for i > 0 {
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Complete")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		reader.ReadString('\n')

		switch choice {
		case 1:
			//var taskName string
			// fmt.Scanf("%s\n", &taskName)
			// addTask(taskName)
			fmt.Print("Enter the task name : ")
			taskName, _ := reader.ReadString('\n') // Waits until Enter is pressed
			taskName = taskName[:len(taskName)-1]  // Remove newline character
			tasks.AddTask(taskName)
		case 2:
			tasks.ViewTask()
		case 3:
			fmt.Print("Enter task number to mark as complete: ")
			var index int
			fmt.Scan(&index)
			tasks.MarkTaskComplete(index - 1)
		case 4:
			fmt.Print("Enter task number to mark as complete: ")
			var index int
			fmt.Scan(&index)
			tasks.DeleteTask(index - 1)
		case 5:
			fmt.Println("Exiting...")
			i = 0 // to exit the prgram initialize 0 to the i
		}
	}
}
