package arry

import "reflect"

func Interface2Int(arr interface{}) []int {
	value := reflect.ValueOf(arr)
	res := make([]int, value.Len())
	for i := 0; i < value.Len(); i++ {
		res[i] = value.Index(i).Interface().(int)
	}
	return res
}

func Interface2Float32(arr interface{}) []float32 {
	value := reflect.ValueOf(arr)
	res := make([]float32, value.Len())
	for i := 0; i < value.Len(); i++ {
		res[i] = value.Index(i).Interface().(float32)
	}
	return res
}

func Interface2Float64(arr interface{}) []float64 {
	value := reflect.ValueOf(arr)
	res := make([]float64, value.Len())
	for i := 0; i < value.Len(); i++ {
		res[i] = value.Index(i).Interface().(float64)
	}
	return res
}

func Int2Float32(arr []int) []float32 {
	arr_ := make([]float32, len(arr))
	for i := 0; i < len(arr); i++ {
		arr_[i] = float32(arr[i])
	}
	return arr_
}

func Int2Float64(arr []int) []float64 {
	arr_ := make([]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		arr_[i] = float64(arr[i])
	}
	return arr_
}

func Float322Float64(arr []float32) []float64 {
	arr_ := make([]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		arr_[i] = float64(arr[i])
	}
	return arr_
}
