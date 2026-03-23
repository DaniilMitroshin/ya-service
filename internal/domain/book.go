package domain

type Book struct {
	Id       int
	Title    *string
	Author   *string
	NumPages *int
	Rating   *float32
}
