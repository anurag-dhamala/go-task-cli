package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

const DbName string = "task.db"
const BucketName string = "task-cli-bucket"

func getBolt() *bolt.DB {
	database, err := bolt.Open(DbName, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	return database
}

func Insert(k string, v string) {
	dbInstance := getBolt()
	err := dbInstance.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return err
		}
		putErr := bucket.Put([]byte(k), []byte(v))
		return putErr
	})
	if err != nil {
		fmt.Println("error occurred while insert", err)
		return
	}
	defer dbInstance.Close()
}

func Get(k string) {
	dbInstance := getBolt()
	err := dbInstance.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("task-cli-bucket"))
		val := b.Get([]byte(k))
		fmt.Println(string(val))
		return nil
	})
	if err != nil {
		fmt.Println("here", err)
		return
	}
	defer dbInstance.Close()
}

func GetAll() {
	dbInstance := getBolt()
	err := dbInstance.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		cursor := b.Cursor()
		for n, d := cursor.First(); n != nil; n, d = cursor.Next() {
			fmt.Printf("name=%s, desc=%s\n", n, d)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	defer dbInstance.Close()
}
