package test_utils

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Scheme   string
	Host     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func (config *Config) DSN() string {
	dsn := url.URL{
		Scheme: config.Scheme,
		Host:   config.Host,
		User:   url.UserPassword(config.User, config.Password),
		Path:   config.DBName,
	}
	q := dsn.Query()
	q.Add("sslmode", config.SSLMode)
	dsn.RawQuery = q.Encode()
	return dsn.String()
}

func DefaultConfig() *Config {
	return &Config{
		Scheme:   "postgres",
		Host:     "127.0.0.1:5432",
		User:     "Henro",
		Password: "henro",
		DBName:   "rentsflow",
		SSLMode:  "disable",
		TimeZone: "Asia/Shanghai",
	}
}
func DatabaseConnection(config *Config) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DSN())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
