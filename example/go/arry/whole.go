package arry

import "fmt"

func Whole_show() {
	For_arry()
	Arr_value()
	var a int
	var b float32
	var c float64
	var d string
	Type_test(a)
	Type_test(b)
	Type_test(c)
	Type_test(d)
	arr := []float32{1.1, 2, 3, 4}
	arr1 := []float32{3.2, 5.3, 8, 100}
	arr2 := []int{1, 2, 3, 4}
	fmt.Println("结果是", Arr_sum(arr))
	fmt.Println(Arr_add_pixel(arr, arr1))
	fmt.Println(Arr_sub_pixel(arr, arr1))
	fmt.Println(Arr_mul_pixel(arr, arr1))
	fmt.Println(Arr_dev_pixel(arr, arr1))
	fmt.Println(Arr_mean(arr))
	fmt.Println(Arr_mean(arr2))
	b = 2
	fmt.Println(Arr_dev(arr, b))
	fmt.Println(Arr_pow_pixel(arr, 3))
	fmt.Println(Arr_norm(arr))
	fmt.Println(Arr_normalize(arr))
	fmt.Println(Arr_random_int(10, 50))
	fmt.Println(Arr_random_float32(10))
	fmt.Println(Arr_random_float64(10))
	fmt.Println(Arr_mul_pixel(arr2, b))
	fmt.Println(Arr_var(arr2))
	fmt.Println(Arr_random_normalize_float64(10, 1))
}
