package dbmodel

import "github.com/jackc/pgx/v5/pgtype"

type Book struct {
	ID       int
	Title    pgtype.Text
	Author   pgtype.Text
	NumPages pgtype.Int4
	Rating   pgtype.Float4
}
