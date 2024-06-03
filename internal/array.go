package internal

func ReverseArray[T any](arr []T) []T {
	length := len(arr)
	reversed := make([]T, length)
	for i, v := range arr {
		reversed[length-i-1] = v
	}
	return reversed
}
