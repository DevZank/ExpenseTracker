package config

import (
	"database/sql"
	f "fmt"
	"os"

	dotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := f.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)	
	}

	f.Println("Sucessfuly Conection with Database...")

	return dbConn
}