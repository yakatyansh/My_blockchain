package main

import "github.com/boltdb/bolt"

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
