package iter

// Iterator interface like java.util.Iterator<E>
type Iterator[E any] interface {
	HasNext() bool
	Next() E
}
