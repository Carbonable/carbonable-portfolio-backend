package utils

import (
	"database/sql"

	"entgo.io/ent/dialect"
	"github.com/carbonable/carbonable-portfolio-backend/ent"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenDB(url string) (*ent.Client, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
