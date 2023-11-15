package iter

// Iterable interface like java.lang.Iterable<T>
type Iterable[T any] interface {
	GetIterator() Iterator[T]
}
