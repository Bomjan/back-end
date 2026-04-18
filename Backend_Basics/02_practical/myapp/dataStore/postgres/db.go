package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "localhost"
	postgres_port     = "5432"
	postgres_user     = "sundrabomjan"
	postgres_password = "whoam100"
	postgres_dbname   = "enroll"
)

var Db *sql.DB

func init() {
	db_info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname,
	)
	fmt.Println(db_info)

	var err error
	Db, err = sql.Open("postgres", db_info)
	if err != nil {
		panic(err)
	} else {
		log.Println("Database created successfully")
	}
}
