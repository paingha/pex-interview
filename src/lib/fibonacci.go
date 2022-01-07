package lib

// FibonacciSequence returns the fibonacci sequence number
func FibonacciSequence(number int) int {
	if number <= 1 {
		return number
	}
	return FibonacciSequence(number-1) + FibonacciSequence(number-2)
}
