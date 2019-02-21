package model

import (
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

// Task represents a task in a TODO list
type Task struct {
	ID          int
	Desc        string
	CreateAt    time.Time
	CompletedAt time.Time
}

// TaskStorage is an interface to define what a Task model can do
type TaskStorage interface {
	Add(desc string, date time.Time) (*Task, error)
	Get(id int) (*Task, error)
	Complete(ids []int, date time.Time) ([]Task, error)
	Delete(id int) error
	List() ([]Task, error)
}

type taskService struct {
	db Database
}

// NewTaskService is responsible to return a TaskStorage interface to enable
// all the CRUD operations on a Task struct
func NewTaskService(db *bolt.DB) TaskStorage {
	return &taskService{
		db: Database{DB: db},
	}
}

// Add a task by calling the the db layer
func (t *taskService) Add(desc string, date time.Time) (*Task, error) {
	return t.db.Add(desc, date)
}

// Get a task by calling the the db layer
func (t *taskService) Get(id int) (*Task, error) {
	return t.db.Get(id)
}

// Complete an array of tasks ids by calling the the db layer
func (t *taskService) Complete(ids []int, date time.Time) ([]Task, error) {
	return t.db.Complete(ids, date)
}

// Delete a task by calling the the db layer
func (t *taskService) Delete(id int) error {
	return t.db.Delete(id)
}

// List gets all tasks from db and sort it 1st by CreatedAt then by CompletedAt
func (t *taskService) List() ([]Task, error) {

	tasks, err := t.db.List()
	if err != nil {
		return nil, err
	}

	createdAtSort := func(i, j int) bool {
		return tasks[i].CreateAt.Before(tasks[j].CreateAt)
	}

	sort.SliceStable(tasks, createdAtSort)

	completedAtSort := func(i, j int) bool {
		return tasks[i].CompletedAt.Before(tasks[j].CompletedAt)
	}

	sort.SliceStable(tasks, completedAtSort)

	return tasks, nil
}
