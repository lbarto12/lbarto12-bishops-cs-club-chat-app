package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"time"
)

const (
	maxOpenConn = 75
	maxIdleTime = 3 * time.Minute
	maxIdleConn = 25
)

var Database *sqlx.DB

func InitDB() error {
	host, ok := os.LookupEnv("PG_HOST")
	if !ok {
		return errors.New("PG_HOST environment variable not set")
	}
	port, ok := os.LookupEnv("PG_PORT")
	if !ok {
		return errors.New("PG_PORT environment variable not set")
	}
	user, ok := os.LookupEnv("PG_USER")
	if !ok {
		return errors.New("PG_USER environment variable not set")
	}
	password, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		return errors.New("PG_PASSWORD environment variable not set")
	}
	db, ok := os.LookupEnv("PG_DB")
	if !ok {
		return errors.New("PG_DB environment variable not set")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db,
	)

	database, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	database.SetMaxOpenConns(maxOpenConn)
	database.SetConnMaxIdleTime(maxIdleTime)
	database.SetMaxIdleConns(maxIdleConn)

	Database = database
	return database.Ping()
}
