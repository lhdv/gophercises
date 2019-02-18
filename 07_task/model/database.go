package model

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("TaskBucket")

// Just to check if Database struct implements TaskStorage interface
var _ TaskStorage = &Database{}

// Database is the layer that talks direct to the BoltDB database
type Database struct {
	DB *bolt.DB
}

//
// TaskStorage Implementation
//

// Add a task into BoltDB
func (d *Database) Add(desc string, date time.Time) (*Task, error) {
	var task Task

	fn := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))

		id, _ := b.NextSequence()

		task.ID = int(id)
		task.Desc = strings.TrimSpace(desc)
		task.CreateAt = date

		// Prepare the data to be saved into DB
		buf, err := json.Marshal(task)
		if err != nil {
			log.Println("[DATABASE][ERROR] - Add: Can't marshal task into json")
			return err
		}

		err = b.Put(itob(task.ID), buf)

		return err
	}

	return &task, d.DB.Update(fn)
}

// Get an specific task based on its ID(which must be the same as the key)
func (d *Database) Get(id int) (*Task, error) {
	var task Task

	fn := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))

		v := b.Get(itob(id))

		err := json.Unmarshal(v, &task)
		if err != nil {
			log.Println("[DATABASE][ERROR] - Get: Can't unmarshal json into a task")
			return err
		}

		return nil
	}

	return &task, d.DB.View(fn)
}

// Complete set a completion date for a given list of task ids
func (d *Database) Complete(ids []int, date time.Time) ([]Task, error) {
	var tasks []Task

	for _, id := range ids {

		task, err := d.Get(id)
		if err != nil {
			return nil, err
		}

		fn := func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(taskBucket))

			task.CompletedAt = date

			// Prepare the data to be saved into DB
			buf, err := json.Marshal(*task)
			if err != nil {
				log.Println("[DATABASE][ERROR] - Complete: Can't marshal task into json")
				return err
			}

			err = b.Put(itob(task.ID), buf)

			return err
		}

		err = d.DB.Update(fn)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, *task)
	}

	return tasks, nil
}

// Delete will remove a given task from database
func (d *Database) Delete(id int) error {

	return d.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))
		err := b.Delete(itob(id))
		return err
	})
}

// List will retrieve all the tasks from the database
func (d *Database) List() ([]Task, error) {
	var tasks []Task
	var t Task

	fn := func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(taskBucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &t)
			if err != nil {
				log.Println("[DATABASE][ERROR] - List: Can't unmarshal json into a task")
				return err
			}

			tasks = append(tasks, t)
		}

		return nil
	}

	return tasks, d.DB.View(fn)
}

//
// Helpers
//

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
