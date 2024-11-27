package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("SQL error (open database): %v", err)
	}

	dropTableQuery := `DROP TABLE IF EXISTS users;`
	_, err = DB.Exec(dropTableQuery)
	if err != nil {
		log.Fatalf("SQL error (drop users table): %v", err)
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT NOT NULL)`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("SQL error (create users table): %v", err)
	}

}
