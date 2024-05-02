package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlConfig struct {
	User         string
	Password     string
	DatabaseName string
	IP           string
	Port         uint16
}

type DBConnection struct {
	db *sql.DB
}

func CreateDB(config SqlConfig) (DBConnection, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.IP, config.Port, config.DatabaseName)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return DBConnection{}, fmt.Errorf("cannot connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return DBConnection{}, fmt.Errorf("cannot ping database: %v", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return DBConnection{db: db}, nil
}
