package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func GetDBFilePath() string {
	dbFilePath := os.Getenv("TODO_DBFILE")
	if dbFilePath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("[ERROR] Can't return the current directory: %s\n", err)
		}
		dbFilePath = filepath.Join(currentDir, "scheduler.db")
	}
	return dbFilePath
}
func InitDB() {
	DB, err := sql.Open("sqlite", GetDBFilePath())
	if err != nil {
		log.Fatalf("[ERROR] Can't open database: %s\n", err)
	}
	querySQL := `CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		title TEXT NOT NULL,
		comment TEXT,
		repeat TEXT CHECK(length(repeat) <= 128)
	);`

	_, err = DB.Exec(querySQL)
	log.Println("[INFO] Creating new table")

	if err != nil {
		log.Fatalf("[ERROR] Executes error: %s\n", err)
	}
	_, err = DB.Exec(`CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);`)
	if err != nil {
		log.Fatal(err)
	}
}
