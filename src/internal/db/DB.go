package DB

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type DB struct {
	Conn *pgx.Conn
}

func InitDB(url string) *DB {
	var db DB

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.Conn = conn
	return &db
}
