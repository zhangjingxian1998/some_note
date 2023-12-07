package arry

import "math/rand"

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

func Arr_random_normalize_float64(length int, scale float64) interface{} {
	arr := Arr_random_float64(length)
	tmp := Arr_mul_pixel(arr, scale)
	mean_arr := Arr_mean(tmp)
	var_arr := Arr_var(tmp)
	tmp = Arr_sub_pixel(tmp, mean_arr)
	tmp = Arr_dev_pixel(tmp, var_arr)
	tmp = Arr_mul_pixel(tmp, scale)
	return tmp
}
