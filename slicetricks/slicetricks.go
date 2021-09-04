package slicetricks

func AppendVector[T any](slice []T, elems ...T) []T {
	return append(slice, elems...)
}

func Copy[T any](slice []T) []T {
	new := make([]T, len(slice))
	copy(slice, new)

	return new
}

func Cut[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index:]...)
}

func Delete[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteUnordered[T any](slice []T, index int) []T {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Expand[T any](slice []T, index int, n int) []T {
	return append(slice[:index], append(make([]T, n), slice[index:]...)...)
}

func Extend[T any](slice []T, n int) []T {
	return append(slice, make([]T, n)...)
}

func Filter[T any](slice []T, keep func(T) bool) []T {
	index := 0
	for _, value := range slice {
		if keep(value) {
			slice[index] = value
			index++
		}
	}

	return slice[:index]
}

// TODO: Insert, InsertVector

func Push[T any](slice []T, elem T) []T {
	return append(slice, elem)
}

func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func PushFront[T any](slice []T, elem T) []T {
	return append([]T{elem}, slice...)
}

func PopFront[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

func Reverse[T any](slice []T) []T {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right+1 {
		slice[left], slice[right] = slice[right], slice[left]
	}

	return slice
}

// TODO: All the rest.
