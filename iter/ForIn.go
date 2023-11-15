package iter

// ForIn Function that mimics foreach-style iteration construct over an iterable.
// Each item provided by Next() method of the iterator is passed to the closure, so the closure body acts as the loop body.
// Limitation: No way to break the loop by code within the closure body other than the case HasNext() is false.
func ForIn[E any](iterable Iterable[E], closure func(item E)) {
	var iter Iterator[E] = iterable.GetIterator()
	for iter.HasNext() {
		closure(iter.Next())
	}
}
