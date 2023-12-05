package fibonacci

func Fibonacci_memory(i int, arr *[]int) int {
	if i == 1 {
		(*arr)[i-1] = 1
		return 1
	}
	if i == 2 {
		(*arr)[i-1] = 1
		return 1
	}
	if (*arr)[i] != 0 {
		return (*arr)[i-1]
	}

	var result1 int
	var result2 int
	result1 = Fibonacci_memory(i-1, arr)
	result2 = Fibonacci_memory(i-2, arr)

	(*arr)[i-1] = result1 + result2
	return result1 + result2
}
func Fibonacci_memory_start(i int) int {
	arr := make([]int, i)
	if i == 1 {
		arr[i-1] = 1
		return 1
	}
	if i == 2 {
		arr[i-1] = 1
		return 1
	}
	var result1 int
	var result2 int
	result1 = Fibonacci_memory(i-1, &arr)
	result2 = Fibonacci_memory(i-2, &arr)
	return result1 + result2
}
