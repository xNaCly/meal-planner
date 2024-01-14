package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgx.Conn
}

var internal *Database = &Database{}

func Get() *Database {
	return internal
}

func (d *Database) Init(connStr string) error {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}
	d.conn = conn
	return nil
}

func (d *Database) Destroy() error {
	return d.conn.Close(context.Background())
}
