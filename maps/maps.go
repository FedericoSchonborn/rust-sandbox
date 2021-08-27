package maps

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Insert(key K, value V) {
	m[key] = value
}
