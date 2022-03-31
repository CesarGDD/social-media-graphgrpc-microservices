package pg

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	hostname      = os.Getenv("POSTGRES_HOST")
	host_port     = 5432
	username      = os.Getenv("POSTGRES_USER")
	password      = os.Getenv("POSTGRES_PASSWORD")
	database_name = os.Getenv("POSTGRES_DB")
)

func Connect() *sql.DB {
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		fmt.Printf("Error conecting to the DB", err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS post_likes (id SERIAL PRIMARY KEY, created_at BIGINT NOT NULL, user_id INTEGER NOT NULL, post_id INTEGER NOT NULL);")

	db.Exec("CREATE TABLE IF NOT EXISTS comment_likes (id SERIAL PRIMARY KEY, created_at BIGINT NOT NULL, user_id INTEGER NOT NULL, comment_id INTEGER NOT NULL);")

	return db
}
