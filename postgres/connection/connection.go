package connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"log"
)

func CheckConnection() {
	ctx := context.Background()

	////
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "admin",
	})

	if err != nil {
		log.Fatal("Error DB connection: ", error.Error(err))
	}

	if err = conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("DB connection successes")
}
