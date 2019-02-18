package model

import "time"

// Task represents a task in a TODO list
type Task struct {
	ID          int
	Desc        string
	CreateAt    time.Time
	CompletedAt time.Time
}

// Add(desc string, date time.Time) (*Task, error)
// Get(id int) (*Task, error)
// Complete(ids []int, date time.Time) ([]Task, error)
// Delete(id int) error
// List() ([]Task, error)

// CreateTask put a task in a TODO list
func CreateTask(task string) (*Task, error) {
	return nil, nil
}

// ListTasks get all tasks and list them
func ListTasks() ([]Task, error) {
	return nil, nil
}

// GetTask get a specific task by its list order
func GetTask(id int) (*Task, error) {
	return nil, nil
}

// DeleteTask get a specific task by its list order
func DeleteTask(id int) error {
	return nil
}

// CompleteTasks mark tasks as done on an  given time
func CompleteTasks(ids []int, when time.Time) ([]Task, error) {
	return nil, nil
}
