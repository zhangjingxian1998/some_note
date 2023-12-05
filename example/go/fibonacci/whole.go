package fibonacci

import (
	"fmt"
	"time"
)

func Whole_show() {
	calculate_times_low := 10
	calculate_times_high := 100
	start := time.Now()
	result := Fibonacci_base(calculate_times_low)
	end := time.Now()
	fmt.Printf("递归斐波那契%d,结果%d,耗时:", calculate_times_low, result)
	fmt.Println(end.Sub(start))
	start = time.Now()
	result = Fibonacci_memory_start(calculate_times_high)
	end = time.Now()
	fmt.Printf("空间换时间斐波那契%d,结果%d,耗时:", calculate_times_high, result)
	fmt.Println(end.Sub(start))
	start = time.Now()
	result = Fibonacci_arr(calculate_times_high)
	end = time.Now()
	fmt.Printf("数组斐波那契%d,结果%d,耗时:", calculate_times_high, result)
	fmt.Println(end.Sub(start))
}
