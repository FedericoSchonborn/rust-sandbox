package encoding

type UnmarshalFunc func([]byte, any) error

func Unmarshal[T any](f UnmarshalFunc, data []byte) (value T, err error) {
	err = f(data, &value)
	return value, err
}

func UnmarshalSlice[S ~[]T, T any](f UnmarshalFunc, data []byte) (value S, err error) {
	err = f(data, &value)
	return value, err
}

func UnmarshalMap[M ~map[K]V, K comparable, V any](f UnmarshalFunc, data []byte) (value M, err error) {
	err = f(data, &value)
	return value, err
}
