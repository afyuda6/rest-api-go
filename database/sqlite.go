package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() {
	var _ error
	DB, _ = sql.Open("sqlite3", "./rest_api_go.db")
	dropUsersTableSQL := `DROP TABLE IF EXISTS users;`
	_, _ = DB.Exec(dropUsersTableSQL)
	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT NOT NULL)`
	_, _ = DB.Exec(createUsersTableSQL)
}
