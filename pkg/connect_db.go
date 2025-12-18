package pkg

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/stdlib"
)

func ConnectDB() (*sql.DB, error) {

	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	ssl := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || name == "" || user == "" {
		return nil, errors.New("database env vars are not set")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
