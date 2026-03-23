package modelDTO

//func Optional[T any](val *T, isValid bool){}

type BookUpdateDTO struct {
	Id       int
	Title    Optional[string]
	Author   Optional[string]
	NumPages Optional[int]
	Rating   Optional[float32]
}
