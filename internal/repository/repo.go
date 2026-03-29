package repository

import (
	"Vservice/internal/domain"
	"Vservice/internal/repository/dbmodel"
	"context"

	//"github.com/jackc/pgx/v5/pgtype"
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
		rating double precision 
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
	VALUES($1,$2,$3,$4);
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

func (r *Repo) SelectById(ctx context.Context, id int) (*domain.Book, error) {
	query := `
	SELECT * from books WHERE id = $1
	`
	var book domain.Book
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.NumPages,
		&book.Rating)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// update через case
// dinamic sql query builder - делать sql запрос конкатенацией и тп
// через 1 query - всегда update books set ....
func (r *Repo) UpdateData_tx(ctx context.Context, bookUpdateParams dbmodel.BookUpdateParams) error {

	querySelect := `
	SELECT id, title, author, num_pages, rating FROM books
	WHERE id = $1;
	`
	queryUpdate := `
	UPDATE books
	SET title = $1,
	author = $2,
	num_pages = $3,
	rating = $4
	WHERE id = $5;
	`

	var bookDB dbmodel.BookDB

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = tx.QueryRow(ctx, querySelect, bookUpdateParams.Id).Scan(
		&bookDB.ID,
		&bookDB.Title,
		&bookDB.Author,
		&bookDB.NumPages,
		&bookDB.Rating)
	if err != nil {
		return err
	}

	if bookUpdateParams.Title.IsSet() {
		if bookUpdateParams.Title.Value != nil {
			bookDB.Title.Valid = true
			bookDB.Title.String = *bookUpdateParams.Title.Ptr()
		} else {
			bookDB.Title.Valid = false
		}
	}

	if bookUpdateParams.Author.IsSet() {
		if bookUpdateParams.Author.Value != nil {
			bookDB.Author.Valid = true
			bookDB.Author.String = *bookUpdateParams.Author.Ptr()
		} else {
			bookDB.Author.Valid = false
		}
	}

	if bookUpdateParams.NumPages.IsSet() {
		if bookUpdateParams.NumPages.Value != nil {
			bookDB.NumPages.Valid = true
			bookDB.NumPages.Int64 = int64(*bookUpdateParams.NumPages.Ptr())
		} else {
			bookDB.NumPages.Valid = false
		}
	}

	if bookUpdateParams.Rating.IsSet() {
		if bookUpdateParams.Rating.Value != nil {
			bookDB.Rating.Valid = true
			bookDB.Rating.Float64 = float64(*bookUpdateParams.Rating.Ptr())
		} else {
			bookDB.Rating.Valid = false
		}
	}

	_, err = tx.Exec(ctx, queryUpdate, bookDB.Title, bookDB.Author, bookDB.NumPages, bookDB.Rating, bookDB.ID)
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
