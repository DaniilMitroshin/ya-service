package domain

type CreateBook struct {
	Title    *string
	Author   *string
	NumPages *int
	Rating   *float32
}
