package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func connectDb() {
	var err error

	var connectionString = os.Getenv("DB_URL")
	db, err = pgx.Connect(context.Background(), connectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Database connected successfully")
}
