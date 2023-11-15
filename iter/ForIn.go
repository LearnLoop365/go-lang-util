package iter

func ForIn[E any](iterable Iterable[E], closure func(item E)) {
	var iter Iterator[E] = iterable.GetIterator()
	for iter.HasNext() {
		closure(iter.Next())
	}
}
