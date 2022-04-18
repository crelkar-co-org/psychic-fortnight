package utils

var count = 0

func Trigger() int {
	count = count + 1
	return count
}
