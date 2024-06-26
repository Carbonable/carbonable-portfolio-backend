package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/carbonable/carbonable-portfolio-backend/ent/gql"
	"github.com/carbonable/carbonable-portfolio-backend/ent/resolver"
	apputils "github.com/carbonable/carbonable-portfolio-backend/internal/utils"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := apputils.OpenDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed opening connection to database", err)
		return
	}

	rpcClient, err := rpc.NewProvider(os.Getenv("RPC_URL"), ethrpc.WithHeader("x-apikey", os.Getenv("RPC_API_KEY")))
	if err != nil {
		slog.Error("failed creating rpc provider", err)
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong !")
	})

	graphqlHandler := handler.NewDefaultServer(
		gql.NewExecutableSchema(
			gql.Config{Resolvers: &resolver.Resolver{Client: db, Rpc: rpcClient}},
		),
	)
	graphqlHandler.Use(extension.Introspection{})
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
