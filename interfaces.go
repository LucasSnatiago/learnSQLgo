package main

import "github.com/LucasSnatiago/learnSQLgo/internal/storage"

type Database interface {
	CreateDB(config storage.SqlConfig) (storage.DBConnection, error)
}
