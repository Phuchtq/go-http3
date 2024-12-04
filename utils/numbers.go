package utils

func IsNumberValid[T int64 | float64 | float32](number T) bool {
	return number > 0
}
