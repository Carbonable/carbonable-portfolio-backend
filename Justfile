default:
    just --list

db_url := "postgres://carbonable:carbonable@localhost:5432/carbonable_portfolio?sslmode=disable"

# start docker database
start_db:
    docker compose up -d

# stop docker database
stop_db:
    docker compose down

# run api
api:
    DATABASE_URL={{db_url}} go run cmd/api/main.go

# run handlers
handlers:
    DATABASE_URL={{db_url}} go run cmd/handler/main.go

# add migration
migrate_diff migration_name:
  atlas migrate diff {{migration_name}} \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres?search_path=public"

# migrate schema
migrate:
  atlas migrate apply \
    --dir "file://ent/migrate/migrations" \
    --url "postgres://carbonable:carbonable@localhost:5432/carbonable_portfolio?search_path=public&sslmode=disable"
