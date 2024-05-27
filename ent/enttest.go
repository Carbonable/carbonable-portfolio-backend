package ent

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/dialect/sql"
)

func NewTestClient(t *testing.T) *Client {
	t.Helper()

	drv, err := sql.Open(
		"sqlite3",
		"file:db?mode=memory&_fk=1&_journal_mode=WAL",
	)
	if err != nil {
		t.Fatalf("opening ent client: %v", err)
	}

	c := NewClient(Driver(drv)).Debug()

	if err := c.Schema.Create(
		context.Background(),
	); err != nil {
		t.Fatalf("Running schema migration: %v", err)
	}

	return c
}
