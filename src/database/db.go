package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "user=postgres password=tu_password dbname=devpulse sslmode=disable"
	}

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to Postgres")

	createTables()
}

func createTables() {
	schema := `
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT UNIQUE NOT NULL,
		enabled BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS checks (
		id SERIAL PRIMARY KEY,
		project_id INTEGER REFERENCES projects(id) ON DELETE CASCADE,
		status_code INTEGER,
		latency_ms INTEGER,
		is_up BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_checks_project_at ON checks(project_id, created_at);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}
	fmt.Println("🚀 Tables verified/created successfully")
}