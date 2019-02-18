package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
)

// StorageConfig TODO
type StorageConfig func(*StorageService) error

// StorageService TODO
type StorageService struct {
	Task TaskStorage
	db   *bolt.DB
}

// NewStorageService TODO
func NewStorageService(cfgs ...StorageConfig) (*StorageService, error) {
	var storageSvc StorageService

	for _, cfg := range cfgs {
		if err := cfg(&storageSvc); err != nil {
			log.Println("[STORAGE][ERROR] - NewStorageService: ", err)
			return nil, err
		}
	}

	return &storageSvc, nil
}

// Close TODO
func (ss *StorageService) Close() error {
	return ss.db.Close()
}

// WithBoltDB TODO
func WithBoltDB(dbfile string) StorageConfig {
	return func(ss *StorageService) error {
		if strings.TrimSpace(dbfile) == "" {
			dbfile = "tasks.db"
		}

		db, err := bolt.Open(dbfile, 0600, nil)
		if err != nil {
			log.Println("[STORAGE][ERROR] - OpenDatabase:", err)
			return err
		}

		ss.db = db

		return nil
	}
}

// WithBucket TODO
func WithBucket(bucket string) StorageConfig {
	return func(ss *StorageService) error {

		bucket = strings.TrimSpace(bucket)
		bucketName := []byte(bucket)

		if bucket == "" {
			bucketName = taskBucket
		}

		err := ss.db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(bucketName)
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})

		return err
	}
}

// WithTask TODO
func WithTask() StorageConfig {
	return func(ss *StorageService) error {
		ss.Task = NewTaskService(ss.db)
		return nil
	}
}
