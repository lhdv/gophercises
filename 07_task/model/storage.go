package model

import (
	"time"
)

// TaskStorage is an interface to define what a Task model can do
type TaskStorage interface {
	Add(desc string, date time.Time) (*Task, error)
	Get(id int) (*Task, error)
	Complete(ids []int, date time.Time) ([]Task, error)
	Delete(id int) error
	List() ([]Task, error)
}

// // StorageConfig TODO
// type StorageConfig func(*StorageService) error

// // StorageService TODO
// type StorageService struct {
// 	Task TaskStorage
// 	DB   *bolt.DB
// }

// // NewStorageService TODO
// func NewStorageService(cfgs ...StorageConfig) (*StorageService, error) {
// 	var storageSvc StorageService

//         for _, cfg := range cfgs {

//         }

// 	return &storageSvc, nil
// }
