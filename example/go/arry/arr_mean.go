package arry

import "reflect"

func Arr_mean(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)
	res := value.Index(0).Interface()
	switch arr.(type) {
	case []int:
		arrs := Interface2Int(arr)
		res = Arr_sum_int(arrs)
		res = float64(res.(int)) / float64(value.Len())
	case []float32:
		res = res.(float32)
		arrs := Interface2Float32(arr)
		res = Arr_sum_float32(arrs)
		res = res.(float32) / float32(value.Len())
	case []float64:
		res = res.(float64)
		arrs := Interface2Float64(arr)
		res = Arr_sum_float64(arrs)
		res = res.(float64) / float64(value.Len())
	}
	return res
}
