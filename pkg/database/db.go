package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	db *bun.DB
}

func checkConnection(db *bun.DB) error {
	_, err := db.ExecContext(context.TODO(), "SELECT 1")
	return err
}

func Connect() (DB, error) {

	pgConnString := fmt.Sprintf("postgres://%s:%s@feed-my-health-db-1:%s/%s?sslmode=disable", os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pgConnString)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := checkConnection(db); err != nil {
		return DB{}, err
	}

	return DB{db}, nil
}

func (d DB) Close() {
	d.db.Close()
}
