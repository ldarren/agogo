package models

import (
	"os"
	"log"
	"time"
	"context"
	"os/signal"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ldarren/agogo/config"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

var pool *sql.DB // Database connection pool.

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 60 * time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func Query(ctx context.Context, username string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var password string
	err := pool.QueryRowContext(ctx, "select u.password from user as u where u.username = :username;", sql.Named("username", username)).Scan(&password)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}
	log.Println("password=", password)
}

func init(){
	var err error
	// Opening a driver typically will not attempt to connect to the database.
	pool, err = sql.Open(*config.DB.Ddriver, *config.DB.Dsn)

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(*config.DB.Dclt)
	pool.SetMaxIdleConns(*config.DB.Didle)
	pool.SetMaxOpenConns(*config.DB.Dopen)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	Ping(ctx)

	Query(ctx, "hello")
}
