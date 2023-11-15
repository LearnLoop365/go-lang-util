package iter

type Iterable[T any] interface {
	GetIterator() Iterator[T]
}
