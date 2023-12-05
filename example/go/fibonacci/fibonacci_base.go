package fibonacci

func Fibonacci_base(i int) int {
	if i == 1 {
		return 1
	}
	if i == 2 {
		return 1
	}
	return Fibonacci_base(i-1) + Fibonacci_base((i - 2))
}
