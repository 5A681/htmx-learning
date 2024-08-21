package database

import (
	"fmt"
	"htmx-learning/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDatabase(config config.Config) *sqlx.DB {
	conStr := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable`,
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT)
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
