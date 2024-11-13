package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (s *SqlxDriverOptions) Connect() (*sqlx.DB, error) {
	conn, err := sqlx.Open(s.DriverName, s.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("postgre.Connection - sqlx.Open: %w", err)
	}

	conn.SetMaxOpenConns(s.MaxOpenConnections)
	conn.SetMaxIdleConns(s.MaxIdleConnections)
	conn.SetConnMaxLifetime(s.MaxLifetime)

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("postgre.Connection - sqlx.Ping: %w", err)
	}

	return conn, nil
}
