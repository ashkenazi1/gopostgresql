package gopostgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	DB *sql.DB
}

var dbInstance *PostgresqlRepository

func GetPostgresql(Host string, Port int, User string, Password string, Dbname string, SSLMode bool) (*PostgresqlRepository, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require", User, Password, Host, Port, Dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	dbInstance = &PostgresqlRepository{DB: db}
	return dbInstance, nil
}

func (p *PostgresqlRepository) Query(query string) (*sql.Rows, error) {
	return p.DB.Query(query)
}

func (p *PostgresqlRepository) Ping() error {
	return p.DB.Ping()
}

func (p *PostgresqlRepository) Close() {
	p.DB.Close()
}
