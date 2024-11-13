package main

import (
	"click-counter/internal/config"
	"click-counter/internal/utils"
	"click-counter/pkg/logger"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"path/filepath"
)

const (
	dir = "cmd/migration/migrations"
)

var (
	up   bool
	down bool
)

func init() {
	logger.InitZeroLogger()
	config.Config.MustInitializeConfig()
}

func main() {
	flag.BoolVar(&up, "up", false, "involves creating new tables, columns, or other database structures")
	flag.BoolVar(&down, "down", false, "involves dropping tables, columns, or other structures")
	flag.Parse()

	conn, err := utils.SetupDefaultPostgreConnection()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if up {
		err = migrate(conn, "up")
		if err != nil {
			logger.ZeroLogger.Error().Msg(fmt.Errorf("migration.up - migrate: %w", err).Error())
			return
		}
	}

	if down {
		err = migrate(conn, "down")
		if err != nil {
			logger.ZeroLogger.Error().Msg(fmt.Errorf("migration.down - migrate: %w", err).Error())
			return
		}
	}
}

func migrate(db *sqlx.DB, action string) (err error) {
	logger.ZeroLogger.Info().Msg(fmt.Sprintf("migrate - running %s migration", action))

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(cwd, dir, fmt.Sprintf("*.%s.sql", action)))
	if err != nil {
		return fmt.Errorf("migrate - filepath.Glob: %w", err)
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("migrate - os.ReadFile: %w", err)
		}

		_, err = db.Exec(string(data))
		if err != nil {
			return fmt.Errorf("migrate - db.Exec: %w", err)
		}
	}

	logger.ZeroLogger.Info().Msg("Migration finished successfully")
	return
}
