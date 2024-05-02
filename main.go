package main

import (
	"fmt"
	"log"

	"github.com/LucasSnatiago/learnSQLgo/internal/storage"
)

func main() {
	_, err := storage.CreateDB(storage.SqlConfig{
		User:         "user",
		Password:     "password",
		DatabaseName: "db",
		IP:           "localhost",
		Port:         3306,
	})
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}

	fmt.Println("Connected to the database!")
}
