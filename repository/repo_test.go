package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepo_CreateAndDropTable(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("set TEST_DATABASE_URL to run integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("create pool: %v", err)
	}
	defer pool.Close()

	repo := NewRepo(pool)

	if err := repo.DropTable(ctx); err != nil {
		t.Fatalf("pre-clean drop: %v", err)
	}

	if err := repo.CreateTable(ctx); err != nil {
		t.Fatalf("create table: %v", err)
	}

	var exists bool
	if err := pool.QueryRow(ctx, `SELECT to_regclass('public.books') IS NOT NULL`).Scan(&exists); err != nil {
		t.Fatalf("check table exists: %v", err)
	}
	if !exists {
		t.Fatal("books table was not created")
	}

	if err := repo.DropTable(ctx); err != nil {
		t.Fatalf("drop table: %v", err)
	}
}
