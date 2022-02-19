package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var SQLX *sqlx.DB

func ConnectDatabaseBySQLX() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "P@ssw0rd")
	dbName := getEnv("DB_NAME", "golang")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic("Failed to connect to database with sqlx")
	}

	SQLX = db
}

func InitDatabase() {
	schema := `
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				firstname text,
				lastname text,
				username text
			);
			`
	SQLX.MustExec(schema)
}
