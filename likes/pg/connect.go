package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname      = "postgres"
	host_port     = 5432
	username      = "user"
	password      = "password"
	database_name = "grpcDB"
)

func Connect() *sql.DB {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		fmt.Printf("Error conecting to the DB", err)
	}
	return db
}
