package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(p *pgxpool.Pool) *Repo {
	return &Repo{
		pool: p,
	}
}

func (r *Repo) CreateTable(ctx context.Context) error {

	query := `
	CREATE TABLE IF NOT EXISTS books(
		id bigint generated always as identity primary key,
		title text,
		author text,
		num_pages integer,
		rating real
		);
	`
	_, err := r.pool.Exec(ctx, query)
	return err
}

func (r *Repo) DropTable(ctx context.Context) error {
	_, err := r.pool.Exec(ctx,
		`
	DROP TABLE IF EXISTS books;
	`)
	return err
}

func (r *Repo) InsertData(ctx context.Context) error {
	query := `
	INSERT INTO books(title,author,num_pages, rating)
	VALUES($1,$2,$3,$4)
	`
	data := [][]any{
		{"book1", "author1", 150, 3.27},
		{"book2", "author2", 250, 4.44},
		{"book3", "author3", 350, 1.44},
		{"book4", "author4", 444, 2.44},
		{nil, "author5", nil, nil},
	}

	for _, vals := range data {
		_, err := r.pool.Exec(ctx, query, vals...)
		if err != nil {
			return err
		}

	}
	return nil
}
