package dbmodel

import "Vservice/internal/shared"

type BookUpdateParams struct {
	Id       int64
	Title    shared.Optional[string]
	Author   shared.Optional[string]
	NumPages shared.Optional[int]
	Rating   shared.Optional[float64]
}
