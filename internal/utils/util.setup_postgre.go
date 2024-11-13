package utils

import (
	"click-counter/internal/config"
	"click-counter/pkg/postgre"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

func SetupDefaultPostgreConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable timezone=Asia/Jakarta",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBName,
	)

	defaultDriverOptions := postgre.NewSqlxDriverOptions(
		"postgres",
		dsn,
		100,
		10,
		15*time.Minute,
	)

	conn, err := defaultDriverOptions.Connect()
	if err != nil {
		return nil, fmt.Errorf("SetupDefaultPostgreConnection - defaultDriverOptions.Connect: %w", err)
	}

	return conn, nil
}
