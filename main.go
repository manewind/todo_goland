package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"time"
)

type TaskStatus string

const (
	inProgress TaskStatus = "in-progress"
	completed  TaskStatus = "completed"
)

func getTaskStatuses() []TaskStatus {
	return []TaskStatus{inProgress, completed}
}

type Task struct {
	ID          int
	Name        string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func isEmpty(slice []Task) bool {
	return len(slice) == 0
}

var tasks []Task
var nextID int = 1

func showMenu() (string, error) {
	options := []string{
		"1. Add task",
		"2. Remove task",
		"3. Edit task",
		"4. Show all tasks",
		"5. Show tasks by status",
		"6. Exit",
	}

	prompt := promptui.Select{
		Label: "Select option",
		Items: options,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return options[index], nil
}

func addTask() {
	prompt := promptui.Prompt{
		Label: "Add task",
	}

	taskName, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error reading task name  %v\n", err)
		return
	}

	descPrompt := promptui.Prompt{
		Label: "Enter task description",
	}

	taskDesc, err := descPrompt.Run()
	if err != nil {
		fmt.Printf("Error reading task description %v\n", err)
		return
	}

	task := Task{
		ID:          nextID,
		Name:        taskName,
		Description: taskDesc,
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	nextID++

	fmt.Println("Task added:", task.Name)
}

func removeTask() {
	if isEmpty(tasks) {
		fmt.Println("No tasks to remove")
		return
	}
	fmt.Println("Enter the name of task you want to remove")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Name)
	}

	promptRemove := promptui.Prompt{
		Label: "Task to remove",
	}

	taskToRemove, err := promptRemove.Run()
	if err != nil {
		fmt.Printf("Error reading task name %v\n", err)
		return
	}

	removed := false
	for i, task := range tasks {
		if task.Name == taskToRemove {
			tasks = append(tasks[:i], tasks[i+1:]...)
			removed = true
			fmt.Println("Task removed:", task.Name)
			break
		}
	}

	if !removed {
		fmt.Printf("No task found with the name '%s'.\n", taskToRemove)
	}
}

func showTasks() {
	if isEmpty(tasks) {
		fmt.Println("no tasks to show")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s (Status: %s)\n", i+1, task.Name, task.Status)
	}
}

func editTask() {
	if isEmpty(tasks) {
		fmt.Println("no tasks to edit")
		return
	}

	for _, task := range tasks {
		fmt.Printf("%d. %s (ID: %d, Status: %s)\n", task.ID, task.Name, task.ID, task.Status)
	}

	promptEdit := promptui.Prompt{
		Label: "Enter the task ID",
	}

	taskIDStr, err := promptEdit.Run()
	if err != nil {
		fmt.Printf("Error reading task ID %v\n", err)
		return
	}

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		fmt.Printf("Error reading task ID %v\n", err)
		return
	}
	var taskToEdit *Task
	for i := range tasks {
		if tasks[i].ID == taskID {
			taskToEdit = &tasks[i]
			break
		}
	}

	if taskToEdit == nil {
		fmt.Printf("No task found with ID %d\n", taskID)
		return
	}

	editPrompt := promptui.Prompt{
		Label: "Enter new task name (or keep current)",
	}

	newTaskName, err := editPrompt.Run()
	if err != nil {
		fmt.Printf("Error reading task name %v\n", err)
	}
	if newTaskName != "" {
		taskToEdit.Name = newTaskName
	}

	statusPrompt := promptui.Select{
		Label: "Enter new task status (or keep current)",
		Items: []TaskStatus{inProgress, completed},
	}

	_, newStatus, err := statusPrompt.Run()
	if err != nil {
		fmt.Printf("Error selecting status %v\n", err)
		return
	}

	taskToEdit.Status = TaskStatus(newStatus)
	taskToEdit.UpdatedAt = time.Now()

	fmt.Printf("Task '%s' updated successfully.\n", taskToEdit.Name)
}

func showTaskByStatus() {
	if isEmpty(tasks) {
		fmt.Println("no tasks to show")
		return
	}

	statuses := getTaskStatuses()

	statusStrings := make([]string, len(statuses))
	for i, status := range statuses {
		statusStrings[i] = string(status)
	}

	promptSelectStatus := promptui.Select{
		Label: "Enter task status to show",
		Items: statusStrings,
	}

	_, selectedStatus, err := promptSelectStatus.Run()
	if err != nil {
		fmt.Printf("Error reading status selection %v\n", err)
		return
	}

	var status TaskStatus
	for _, s := range statuses {
		if selectedStatus == string(s) {
			status = s
			break
		}
	}

	found := false
	for i, task := range tasks {
		if task.Status == status {
			fmt.Printf("%d. %s (Status: %s)\n", i+1, task.Name, task.Status)
			found = true
		}
	}
	if !found {
		fmt.Println("No tasks to show for selected status")
	}
}

func main() {
	for {
		selectedOption, err := showMenu()
		if err != nil {
			fmt.Println(err)
			return
		}

		if selectedOption == "6. Exit" {
			fmt.Println("Exiting program.")
			break
		}

		switch selectedOption {
		case "1. Add task":
			addTask()
		case "2. Remove task":
			removeTask()
		case "3. Edit task":
			editTask()
		case "4. Show all tasks":
			showTasks()
		case "5. Show tasks by status":
			showTaskByStatus()
		}
	}
}
