package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"entgo.io/ent/dialect"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-portfolio-backend/ent"
	"github.com/carbonable/carbonable-portfolio-backend/internal/model"

	entsql "entgo.io/ent/dialect/sql"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDb(url string) (*ent.Client, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	ctx := context.Background()
	db, err := openDb(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed opening connection to database", err)
		return
	}

	rpcClient, err := rpc.NewProvider(os.Getenv("RPC_URL"), ethrpc.WithHeader("x-apikey", os.Getenv("RPC_API_KEY")))
	if err != nil {
		slog.Error("failed dialing into rpc provider", err)
		return
	}

	// Update base project data
	err = model.Synchronize(ctx, db, rpcClient)
	if err != nil {
		slog.Error("failed to sync contracts", err)
	}

	// project, minter, offseter, yielder, migrator

	// Register handlers to react to events
	fmt.Println("Running handler")
}
