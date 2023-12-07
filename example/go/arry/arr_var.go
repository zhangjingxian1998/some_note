package arry

import "reflect"

func Arr_var(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)
	switch arr.(type) {
	case []int:
		arrs := Interface2Int(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, value.Len())
		return Arr_sum(tmp)
	case []float32:
		arrs := Interface2Float32(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, value.Len())
		return Arr_sum(tmp)
	case []float64:
		arrs := Interface2Float64(arr)
		mean := Arr_mean(arrs)
		tmp := Arr_sub_pixel(arrs, mean)
		tmp = Arr_pow_pixel(tmp, 2)
		tmp = Arr_dev_pixel(tmp, value.Len())
		return Arr_sum(tmp)
	default:
		panic("Only accept []int []floate32 []floate64 as input")
	}

}
