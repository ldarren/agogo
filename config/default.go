package config

import (
	"os"
	"log"
	"flag"
	"time"
	"strconv"
	dotenv "github.com/joho/godotenv"
)

type DatabaseConfig struct{
	Ddriver *string
	Dsn *string
	Dclt *time.Duration
	Didle *int
	Dopen *int
}
var DB DatabaseConfig

func init(){
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("unable to read dotenv: %v", err)
	}

	DB.Ddriver = flag.String("db_driver", os.Getenv("db_driver"), "database driver name")
	//dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	DB.Dsn = flag.String("db_source_name", os.Getenv("db_source_name"), "connection data source name")
	timeout, err := time.ParseDuration(os.Getenv("db_conn_life"))
	if err != nil {
		timeout = 0
	}
	DB.Dclt = flag.Duration("db_conn_life", timeout, "connection max lifetime")
	var val int64
	val, err = strconv.ParseInt(os.Getenv("db_conn_idle"), 10, 0)
	if err != nil {
		val = 2
	}
	DB.Didle = flag.Int("db_conn_idle", int(val), "max idle connections")
	val, err = strconv.ParseInt(os.Getenv("db_conn_open"), 10, 0)
	if err != nil {
		val = 2
	}
	DB.Dopen = flag.Int("db_conn_open", int(val), "max open connections")
	flag.Parse()
}
