package main

import (
	"errors"
	"fmt"
	"strings"
)
//represents todo item
type Task struct {
    id int
    Name string
    Status bool
}

var tasks []Task
var lastId int 


// GetTask retrieves a task by its ID.
func GetTask(id int) (Task, error) {
	for _, task := range tasks {   // this is basically go way of using for each 
		if task.id == id {
			return task, nil  //if found returns task and nil error which means no error 
		}
	}
	return Task{}, errors.New("Task not found") //return emty struct and error, not task  found 
}
// AddTask adds a new task to the list.
func AddTask(title string) {
	lastId++
	task := Task{
		id: lastId,
		Name: title,
		Status: false,
	}
	tasks = append(tasks, task)
}

// UpdateTask updates an existing task by its ID.
func UpdateTask(id int, updatedTask Task) error {   //we can see that only error is returned 
	for i, task := range tasks {
		if task.id == id {
			tasks[i] = updatedTask  //set the task to the the new one 
			return nil
		}
	}
	return errors.New("Task not found")
}

// DeleteTask deletes a task by its ID.
func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.id == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // this takes the tasks before i and tasks after i and it appends them together so task at i is getting removed 
			return nil
		}
	}
	return errors.New("Task not found")
}

// ListTasks lists all tasks.
func ListTasks() {
	fmt.Println("TODO List:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Status{
			status = "Complete"
		}
		// Mark completed tasks with a line through the text.
		if task.Status {
			taskNameWithLine := addLineThroughText(task.Name)
			fmt.Printf("Task %d: %s (%s)\n", task.id, taskNameWithLine, status)
		} else {
			fmt.Printf("Task %d: %s (%s)\n", task.id, task.Name, status)
		}
	}
}
// addLineThroughText adds a line through the text.
func addLineThroughText(text string) string {
	var lines []string
	for _, char := range text {
		lines = append(lines, string(char)+"\u0336")
	}
	return strings.Join(lines, "")
}

func main() {
	var choice int

	for {
		fmt.Println("TODO List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Done")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter task title: ")
			_, _ = fmt.Scanln(&title)
			AddTask(title)
			fmt.Println("Task added.")
		case 2:
			var taskID int
			fmt.Print("Enter the ID of the task you want to mark as done: ")
			_, _ = fmt.Scanln(&taskID)
			taskToMark, err := GetTask(taskID)
			if err == nil {
				taskToMark.Status = true
				UpdateTask(taskToMark.id, taskToMark)
				fmt.Println("Task marked as done.")
			} else {
				fmt.Println("Task not found.")
			}
		case 3:
			ListTasks()
		case 4:
			fmt.Println("Exiting the TODO list application.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}