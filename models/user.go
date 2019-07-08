package models

import (
	"log"
	"time"
	"context"
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

func Insert(ctx context.Context, username string, password string) {
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	stmtIns, err := pool.Prepare("INSERT INTO user VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		log.Fatalf("unable to create prepared statement: %v", err)
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(username, password)
	if err != nil {
		log.Fatalf("unable to exec prepared statement: %v", err)
	}
}

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
func Query(ctx context.Context, username string) {
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	// Prepare statement for reading data
	stmtOut, err := pool.Prepare("SELECT password FROM user WHERE username = ?")
	if err != nil {
		log.Fatalf("unable to create prepared statement: %v", err)
	}
	defer stmtOut.Close()

	var password string

	err = stmtOut.QueryRow(username).Scan(&password) // WHERE number = 13
	if err != nil {
		log.Fatalf("unable to exec prepared statement: %v", err)
	}
	log.Printf("the password of %s is: %s", username, password)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Ping(ctx)

	Query(ctx, "foo")
}
