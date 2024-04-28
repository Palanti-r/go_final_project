package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Store struct {
	*sql.DB
}

func GetDBFilePath() string {
	dbFilePath := os.Getenv("TODO_DBFILE")
	if dbFilePath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("[ERROR] Can't return currnet the current directory: %s\n", err)
		}
		dbFilePath = filepath.Join(currentDir, "scheduler.db")
	}
	return dbFilePath
}
func InitlDB() {
	db, err := sql.Open("sqlite", GetDBFilePath())
	if err != nil {
		log.Fatalf("[ERROR] Can't open database: %s\n", err)
	}
	querySQL := `CREATE TABLE IF NOT EXISTS scheduler (
				id INTEGER PRIMARY KEY,
				date VARCHAR(8) NOT NULL,
  				title TEXT NOT NULL,
    			comment TEXT,
  				repeat VARCHAR(128));
				CREATE INDEX IF NOT EXISTS indexdate ON scheduler (date);`

	log.Println("[INFO] Creating new table")
	_, err = db.Exec(querySQL)

	if err != nil {
		log.Fatalf("[ERROR] Executes error: %s\n", err)
	}
}
