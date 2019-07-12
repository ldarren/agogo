# agogo
A go api framwork

## setup
- `make go-get` or `go get ./...`
- copy config/dotenv to .env and replace the values with your environment setting
- copy config/dotenv to models/.env and replace the values with your test environment setting
- `make go-build` or `go build -o bin/agogo`
- `make start-server` run from binary
- `go run main.go` for quick run

## docker
- `docker build -t agogo .`
- `docker run -d -p 8800:8800 agogo` # access with http://localhost:8800

## test
- `go test -v ./...`

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
