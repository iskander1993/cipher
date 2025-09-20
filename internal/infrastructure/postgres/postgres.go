package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// ConnectPostgres создаёт подключение к PostgreSQL и возвращает *sql.DB
func ConnectPostgres() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("POSTGRES_USER", "user"),
		getEnv("POSTGRES_PASSWORD", "password"),
		getEnv("POSTGRES_DB", "ave_db"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot reach PostgreSQL:", err)
	}

	// Создаём таблицу пользователей
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return db
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
