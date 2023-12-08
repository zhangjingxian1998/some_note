package arry

import (
	"math/rand"
)

func Arr_random_int(lenght int, range_ int) []int {
	arr := make([]int, lenght)
	for i := 0; i < lenght; i++ {
		arr[i] = rand.Intn(range_)
	}
	return arr
}

func Arr_random_float32(lenght int) []float32 {
	arr := make([]float32, lenght)
	for i := 0; i < lenght; i++ {
		arr[i] = rand.Float32()
	}
	return arr
}

func Arr_random_float64(lenght int) []float64 {
	arr := make([]float64, lenght)
	for i := 0; i < lenght; i++ {
		arr[i] = rand.Float64()
	}
	return arr
}

// func Arr_random_normalize_float32(length int, scale float64) float32 {
// 	arr := Arr_random_float64(length)
// 	mean_arr := Arr_mean(arr)
// 	var_arr := Arr_var(arr)
// 	tmp := Arr_sub_pixel(arr, mean_arr)
// 	tmp = Arr_dev_pixel(tmp, var_arr)
// 	tmp = Arr_mul_pixel(tmp, scale)
// 	return tmp
// }

func Arr_random_noise(length int, mean float64, sigma float64) interface{} {
	arr := Arr_random_float64(length)
	arr1 := Arr_sub_pixel(arr, 0.5)
	arr2 := Arr_mul_pixel(arr1, 2)

	mean_arr := Arr_mean(arr2)
	std_var_arr := Arr_std_var(arr2)
	tmp := Arr_sub_pixel(arr2, mean_arr)
	tmp = Arr_dev_pixel(tmp, std_var_arr)
	tmp = Arr_mul_pixel(tmp, sigma)
	tmp = Arr_add_pixel(tmp, mean)
	return tmp
}
