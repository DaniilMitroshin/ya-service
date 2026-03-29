package domain

import "fmt"

type Book struct {
	Id       int
	Title    *string
	Author   *string
	NumPages *int
	Rating   *float32
}

/*
func (b *Book) String() string {
	return fmt.Sprintf("%+v\n", *b)
}
*/

func val[T any](v *T) any {
	if v == nil {
		return nil
	}
	return *v
}

func (b *Book) String() string {
	return fmt.Sprintf("{Id: %v, Title:%v, Author: %v, NumPages:%v, Rating: %v}\n",
		b.Id,
		val(b.Title),
		val(b.Author),
		val(b.NumPages),
		val(b.Rating))
}
