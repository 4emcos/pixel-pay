package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var (
	db *pgxpool.Pool
)

type Pgx interface {
	Exec(context context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(context context.Context, sql string, args ...interface{}) (rows pgx.Rows, err error)
	QueryRow(context context.Context, sql string, args ...interface{}) pgx.Row
	Begin(context context.Context) (pgx.Tx, error)
	Close()
}

func Connect() Pgx {
	if db != nil {
		return db
	}

	hostDb := os.Getenv("DB_HOST")

	if hostDb == "" {
		hostDb = "localhost"
	}

	err := error(nil)
	db, err = pgxpool.New(context.Background(), fmt.Sprintf("user=admin password=oot123 dbname=pixel_pay host=%s port=5432", hostDb))
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	log.Print("Connected to database")
	return db
}
