package postgre

import "time"

type SqlxDriverOptions struct {
	DriverName         string
	DataSourceName     string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxLifetime        time.Duration
}

func NewSqlxDriverOptions(
	driverName string,
	dataSourceName string,
	maxOpenConnections int,
	maxIdleConnections int,
	maxLifetime time.Duration,
) SqlxDriverOptions {
	return SqlxDriverOptions{
		driverName,
		dataSourceName,
		maxOpenConnections,
		maxIdleConnections,
		maxLifetime,
	}
}
