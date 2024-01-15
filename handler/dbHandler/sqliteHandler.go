package dbHandler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteHandler struct {
	Database *sql.DB
}

func CreateSqliteHandler(db_file_path string) DBHandler {
	/* Database File 생성 */
	database, err := sql.Open("sqlite3", db_file_path)
	if err != nil {
		panic(err)
	}

	/* Table이 없다면 생성 */
	statement, err := database.Prepare(
		`CREATE TABLE IF NOT EXISTS file_metadata (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			path TEXT NOT NULL,
			size INTEGER NOT NULL,
			hash TEXT NOT NULL,
			createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS file_hash ON file_metadata (
			hash ASC
		);`,
	)
	statement.Exec()

	return &SqliteHandler{Database: database}
}
