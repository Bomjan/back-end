package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	defaultPostgresHost   = "localhost"
	defaultPostgresPort   = "5432"
	defaultPostgresDBName = "enroll"
)

var Db *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	postgresHost := getEnv("POSTGRES_HOST", defaultPostgresHost)
	postgresPort := getEnv("POSTGRES_PORT", defaultPostgresPort)
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDBName := getEnv("POSTGRES_DBNAME", defaultPostgresDBName)

	if postgresUser == "" || postgresPassword == "" {
		log.Fatal("POSTGRES_USER and POSTGRES_PASSWORD must be set")
	}

	db_info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDBName,
	)

	var err error
	Db, err = sql.Open("postgres", db_info)
	if err != nil {
		panic(err)
	} else {
		log.Println("Database created successfully")
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
