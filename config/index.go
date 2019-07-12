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
	Driver *string
	Dsn *string
	ConnLife *time.Duration
	ConnIdle *int
	ConnOpen *int
}
type PathConfig struct{
	Static *string
}

var DB DatabaseConfig
var PATH PathConfig

func init(){
	cwd, err := os.Getwd()
	log.Printf("config.init CWD: %s\n", cwd)
	err = dotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to read dotenv: %v", err)
	}

	DB.Driver = flag.String("db_driver", os.Getenv("db_driver"), "database driver name")
	//dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	DB.Dsn = flag.String("db_source_name", os.Getenv("db_source_name"), "connection data source name")
	timeout, err := time.ParseDuration(os.Getenv("db_conn_life"))
	if err != nil {
		timeout = 0
	}
	DB.ConnLife = flag.Duration("db_conn_life", timeout, "connection max lifetime")
	var val int64
	val, err = strconv.ParseInt(os.Getenv("db_conn_idle"), 10, 0)
	if err != nil {
		val = 2
	}
	DB.ConnIdle = flag.Int("db_conn_idle", int(val), "max idle connections")
	val, err = strconv.ParseInt(os.Getenv("db_conn_open"), 10, 0)
	if err != nil {
		val = 2
	}
	DB.ConnOpen = flag.Int("db_conn_open", int(val), "max open connections")

	PATH.Static = flag.String("PATH_STATIC", os.Getenv("PATH_STATIC"), "static file path")

	flag.Parse()
}
