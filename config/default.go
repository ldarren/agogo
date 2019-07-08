package config

import (
	"flag"
	"time"
	"log"
	dotenv "github.com/direnv/go-dotenv"
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
	env, err := dotenv.Parse(".env")
	if err != nil {
		log.Fatalf("unable to read dotenv: %v", err)
	}
	DB.Ddriver = flag.String("db_driver", env.db_driver, "database driver name")
	//dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	DB.Dsn = flag.String("db_source_name", env.db_source_name, "connection data source name")
	DB.Dclt = flag.Duration("db_conn_life",env.db_conn_life, "connection max lifetime")
	DB.Didle = flag.Int("db_conn_idle", env.db_conn_idle, "max idle connections")
	DB.Dopen = flag.Int("db_conn_open", env.db_conn_open, "max open connections")
	flag.Parse()
}
