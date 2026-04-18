package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database connection constants - change these to connect to a different database.
// Keeping credentials in code is fine for development, but should use env vars in production.
const (
	postgres_host     = "localhost"
	postgres_port     = "5432"
	postgres_user     = "sundrabomjan"
	postgres_password = "whoam100"
	postgres_dbname   = "enroll"
)

// Db is the shared database connection pool used across the application.
// It's safe for concurrent use by multiple goroutines.
var Db *sql.DB

// init runs automatically when the package is imported - before main() even starts.
// This ensures the database connection is ready before any request is handled.
func init() {
	db_info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname,
	)
	fmt.Println(db_info)

	var err error
	// sql.Open doesn't actually connect - it just prepares the connection pool.
	// The first real connection happens when a query is executed.
	Db, err = sql.Open("postgres", db_info)
	if err != nil {
		// Panic here because the app can't function without a database connection.
		// Better to crash early than fail silently on every request.
		panic(err)
	} else {
		log.Println("Database created successfully")
	}
}
