package arry

import (
	"math"
	"reflect"
)

func Arr_norm(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)
	res := value.Index(0).Interface()
	switch arr.(type) {
	case []int:
		res = float32(res.(int))
		arrs := Arr_pow_pixel(arr, 2)
		res = Arr_sum(arrs)
		res = float32(math.Sqrt(float64(res.(int))))
	case []float32:
		res = res.(float32)
		arrs := Arr_pow_pixel(arr, 2)
		res = Arr_sum(arrs)
		res = float32(math.Sqrt(float64(res.(float32))))
	case []float64:
		res = res.(float64)
		arrs := Arr_pow_pixel(arr, 2)
		res = Arr_sum(arrs)
		res = math.Sqrt(float64(res.(float32)))
	}
	return res
}
