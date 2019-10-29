package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task It is a struct with two fields 
type Task struct {
	Key   int
	Value string
}

//Init not same as init(which is called automatically)
// just an exported function which we need to call.
func Init(dbPath string) error {
	// used var err error insted of db,err :=
	// The reason is if we use :=, then the db will be local function scoped variable
	// We want it to be a package level variable(mentioned above as var db *bolt.DB)
	// In order for it work, we can not be declaring new variables and instead created err as var err error

	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	// This is where we create the bucket if it doesn't exist
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask Creates bucket & task sequence
func CreateTask(task string) (int, error) {
	var id int
	// Update is used for read and write transactions in the DB.
	err := db.Update(func(tx *bolt.Tx) error {
		// Grab the bucket
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		// A closure is a function value that references variables from outside its body.
		// since the id occured inside this closure,and we might never have access to it.
		// so we initialize it above as var id int, in order to access it.
		id = int(id64)
		// in order to put our key and values in the we want it to []byte type
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	return id, err

}

// AllTasks Lists all the task
func AllTasks() ([]Task, error) {
	// We are declaring outside db connection, so that we can set this inside of closure.
	var tasks []Task
	// View is used for only reading from Bolt DB.
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)

		// Iterating over keys can be done via cursor
		// The cursor will start on the very first item in the bucket,
		// And we can continuosly call next,which grabs the next item in the bucket.
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// for every key value pair, we want it to add to your tasks slice.
			tasks = append(tasks, Task{
				Key: btoi(k),
				// The reason we are type casting value to string is that,
				// because we dont want the v(which is []byte) with data as they are going to be reused.
				// So we make sure to copy it to []byte, which is not going to be used here.
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}

// DeleteTasks It Deletes the task with the mentioned key.
func DeleteTasks(key int) error {
	// we are using delete here to mark the task as complete.
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	// binary.BigEndian is the encoding type
	// PutUint64 means we take uint64 and put it into byte slice
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// we want the reverse of ITOB as well
// Later when we want to read this task from db,
// Then we would have no way of taking this []byte back to int
// So that we can tell our cli what intergers to mark as completed(do,basically deleting the entry)
func btoi(b []byte) int {
	// we are casting whatever result we have
	return int(binary.BigEndian.Uint64(b))
}
