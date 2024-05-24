default:
    just --list

db_url := "postgres://carbonable:carbonable@localhost:5432/carbonable_leaderboard?sslmode=disable"

# start docker database
start_db:
    docker compose up -d

# stop docker database
stop_db:
    docker compose stop

# run api
api:
    DATABASE_URL={{db_url}} go run cmd/api/main.go

# run handlers
handlers:
    DATABASE_URL={{db_url}} go run cmd/handler/main.go

