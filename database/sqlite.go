package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./rest_api_go.db")
	if err != nil {
		log.Fatalf("SQL error (open database): %v", err)
	}

	dropUsersTableSQL := `DROP TABLE IF EXISTS users;`
	_, err = DB.Exec(dropUsersTableSQL)
	if err != nil {
		log.Fatalf("SQL error (drop users table): %v", err)
	}

	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT NOT NULL)`
	_, err = DB.Exec(createUsersTableSQL)
	if err != nil {
		log.Fatalf("SQL error (create users table): %v", err)
	}

}
