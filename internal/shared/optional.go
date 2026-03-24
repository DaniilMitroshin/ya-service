package shared

type Optional[T any] struct {
	Set   bool
	Value *T
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{
		Set:   true,
		Value: &v,
	}
}

func Null[T any]() Optional[T] {
	return Optional[T]{
		Set:   true,
		Value: nil,
	}
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) IsSet() bool {
	return o.Set
}

func (o Optional[T]) Ptr() *T {
	return o.Value
}
