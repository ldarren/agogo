# agogo
A go api framwork

## setup
- copy config/dotenv to .env and replace your environment setting
- go run main.go

## features
- multiple route groups
- route parameters
- json rest api
- sql read write
- service can be configured from commoand line, environment variable or dotenv file
- secure configuration, prevented credential checkin into repository
- minimum dependencies, make use of official libraries wherever possible: net/http, html/template, flag, log, db/sql, encoding/json

## dependencies
- "github.com/joho/godotenv"
- "github.com/go-sql-driver/mysql"
- "github.com/julienschmidt/httprouter"
