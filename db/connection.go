package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type databaseConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

func initDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
	}
}

func CreateConnection() (*sql.DB, error) {
	dbConfig := initDatabaseConfig()
	connectionConfigString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.DB_USERNAME, dbConfig.DB_PASSWORD, dbConfig.DB_HOST, dbConfig.DB_PORT, dbConfig.DB_DATABASE)
	db, err := sql.Open("mysql", connectionConfigString)
	if err != nil {
		return nil, err
	}
	return db, err
}
