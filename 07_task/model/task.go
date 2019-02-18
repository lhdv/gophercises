package model

import (
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

// NewTaskService TODO
func NewTaskService(db *bolt.DB) TaskStorage {
	return &taskService{
		db: Database{DB: db},
	}
}

// Add TODO
func (t *taskService) Add(desc string, date time.Time) (*Task, error) {
	return t.db.Add(desc, date)
}

// Get TODO
func (t *taskService) Get(id int) (*Task, error) {
	return t.db.Get(id)
}

// Complete TODO
func (t *taskService) Complete(ids []int, date time.Time) ([]Task, error) {
	return t.db.Complete(ids, date)
}

// Delete TODO
func (t *taskService) Delete(id int) error {
	return t.db.Delete(id)
}

// List TODO
func (t *taskService) List() ([]Task, error) {
	return t.db.List()
}
