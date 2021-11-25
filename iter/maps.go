package iter

type mapIterator[K comparable, V any] struct {
	m    map[K]V
	i    int
	keys []K
}

type MapItem[K comparable, V any] struct {
	Key   K
	Value V
}

func FromMap[M ~map[K]V, K comparable, V any](m M) Iterator[MapItem[K, V]] {
	kvs := make([]MapItem[K, V], len(m))

	var i int
	for key, value := range m {
		kvs[i] = MapItem[K, V]{
			Key:   key,
			Value: value,
		}
		i++
	}

	return FromSlice(kvs)
}

func FromMapKeys[M ~map[K]V, K comparable, V any](m M) Iterator[K] {
	keys := make([]K, len(m))

	var i int
	for key := range m {
		keys[i] = key
		i++
	}

	return FromSlice(keys)
}

func FromMapValues[M ~map[K]V, K comparable, V any](m M) Iterator[V] {
	values := make([]V, len(m))

	var i int
	for _, value := range m {
		values[i] = value
		i++
	}

	return FromSlice(values)
}
