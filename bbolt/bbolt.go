package bbolt

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

func ImportBbolt() error {
	db, err := bolt.Open("mydb.bolt", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	return nil
}

func UpdateBolt() error {
	db, err := bolt.Open("mydb.bolt", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("MyBucket"))
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
	return nil
}

func GetBolt() error {
	db, err := bolt.Open("mydb.bolt", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
	return nil
}
