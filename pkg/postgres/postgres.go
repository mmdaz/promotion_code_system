package postgres

import (
	"github.com/jackc/pgx"
	"github.com/labstack/gommon/log"
	"os"
)

type PGXDatabase struct {
	*pgx.ConnPool
}

type Option struct {
	Host string
	Port int
	User string
	Pass string
	Db   string
}

func NewPGXPostgres(option Option, maxConnections int) *PGXDatabase {
	var err error

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     option.Host,
			Port:     uint16(option.Port),
			User:     option.User,
			Password: option.Pass,
			Database: option.Db,
		},
		MaxConnections: maxConnections,
	}
	log.Debugf("Creating pgx connection pool. host: %v, port: %v", option.Host, option.Port)
	postgresPool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		log.Errorf("Unable to create connection pool. host: %v, error: %v", option.Host, err)
		os.Exit(1)
	}
	log.Debugf("Pgx connection pool created successfully.")

	return &PGXDatabase{postgresPool}
}
