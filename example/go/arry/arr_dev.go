package arry

import (
	"reflect"
)

func Arr_dev(arr1 interface{}, dev interface{}) interface{} {
	value1 := reflect.ValueOf(arr1)
	type_arr1 := reflect.TypeOf(value1.Index(0).Interface())
	type_arr2 := reflect.TypeOf(dev)
	if type_arr1 != type_arr2 {
		panic("两个数组类型必须相同")
	}

	res := value1.Interface()
	switch arr1.(type) {
	case []int:
		res = res.([]float32)
		arr1 := Interface2Int(arr1)
		res = Arr_dev_int(arr1, dev.(int))
	case []float32:
		res = res.([]float32)
		arr1 := Interface2Float32(arr1)
		res = Arr_dev_float32(arr1, dev.(float32))
	case []float64:
		res = res.([]float64)
		arr1 := Interface2Float64(arr1)
		res = Arr_dev_float64(arr1, dev.(float64))
	default:
		panic("只能接收int,float32,float64的数组类型")
	}
	return res
}
func Arr_dev_float32(arr1 []float32, dev float32) []float32 {
	res := make([]float32, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i] / dev
	}
	return res
}

func Arr_dev_float64(arr1 []float64, dev float64) []float64 {
	res := make([]float64, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i] / dev
	}
	return res
}

func Arr_dev_int(arr1 []int, dev int) []float32 {
	res := make([]float32, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = float32(arr1[i]) / float32(dev)
	}
	return res
}
