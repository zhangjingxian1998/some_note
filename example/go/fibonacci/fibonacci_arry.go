package fibonacci

func Fibonacci_arr(i int) int {
	arr := make([]int, i)
	arr[0] = 1
	arr[1] = 1
	for p := 2; p < i; p++ {
		arr[p] = arr[p-1] + arr[p-2]
	}
	return arr[i-1]
}
