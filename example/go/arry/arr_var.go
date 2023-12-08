package arry

import (
	"math"
)

// "math"
// "reflect"

func Arr_var(arr interface{}) interface{} {
	switch arr.(type) {
	case []int:
		arrs := Interface2Int(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, len(arrs))
		return Arr_sum(tmp)
	case []float32:
		arrs := Interface2Float32(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, len(arrs))
		return Arr_sum(tmp)
	case []float64:
		arrs := Interface2Float64(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, len(arrs))
		return Arr_sum(tmp)
	default:
		panic("Only accept []int []floate32 []floate64 as input")
	}
}

// func Arr_std_var(var_ interface{}, length interface{}) interface{} {
// 	switch var_.(type) {
// 	case int:
// 		return math.Sqrt(float64(var_.(int)))
// 	case float32:
// 		return float32(math.Sqrt(float64(var_.(float32))))
// 	case float64:
// 		return math.Sqrt(var_.(float64))
// 	default:
// 		panic("Only accept int floate32 floate64 as input")
// 	}
// }

func Arr_std_var(arr interface{}) interface{} {
	switch arr.(type) {
	case []int:
		var_ := Arr_var(arr)
		return math.Sqrt(var_.(float64))
	case []float32:
		var_ := Arr_var(arr)
		return math.Sqrt(float64(var_.(float32)))
	case []float64:
		var_ := Arr_var(arr)
		return math.Sqrt(var_.(float64))
	default:
		panic("Only accept []int []floate32 []floate64 as input")
	}
}
