package arry

import (
	"reflect"
)

func Arr_mul_pixel(arr1 interface{}, arr2 interface{}) interface{} {
	value1 := reflect.ValueOf(arr1)
	switch arr1.(type) {
	case []int:
		switch arr2.(type) {
		case []int:
			arr1 := Interface2Int(arr1)
			arr2 := Interface2Int(arr2)
			return Arr_mul_pixel_int(arr1, arr2)
		case []float32:
			arr1 := Interface2Int(arr1)
			arr1_float_32 := Int2Float32(arr1)
			arr2 := Interface2Float32(arr2)
			return Arr_mul_pixel_float32(arr1_float_32, arr2)
		case []float64:
			arr1 := Interface2Int(arr1)
			arr1_float_32 := Int2Float64(arr1)
			arr2 := Interface2Float64(arr2)
			return Arr_mul_pixel_float64(arr1_float_32, arr2)
		case int:
			arr1 := Interface2Int(arr1)
			arr_tmp := make([]int, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = arr2.(int)
			}
			return Arr_mul_pixel_int(arr1, arr_tmp)
		case float32:
			arr1 := Interface2Int(arr1)
			arr1_float_32 := Int2Float32(arr1)
			arr_tmp := make([]float32, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = arr2.(float32)
			}
			return Arr_mul_pixel_float32(arr1_float_32, arr_tmp)
		case float64:
			arr1 := Interface2Int(arr1)
			arr1_float_64 := Int2Float64(arr1)
			arr_tmp := make([]float64, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = arr2.(float64)
			}
			return Arr_mul_pixel_float64(arr1_float_64, arr_tmp)
		}
	case []float32:
		switch arr2.(type) {
		case []int:
			arr1 := Interface2Float32(arr1)
			arr2 := Interface2Int(arr2)
			arr2_float_32 := Int2Float32(arr2)
			return Arr_mul_pixel_float32(arr1, arr2_float_32)
		case []float32:
			arr1 := Interface2Float32(arr1)
			arr2 := Interface2Float32(arr2)
			return Arr_mul_pixel_float32(arr1, arr2)
		case []float64:
			arr1 := Interface2Float32(arr1)
			arr1_float_64 := Float322Float64(arr1)
			arr2 := Interface2Float64(arr2)
			return Arr_mul_pixel_float64(arr1_float_64, arr2)
		case int:
			arr1 := Interface2Float32(arr1)
			arr_tmp := make([]float32, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = float32(arr2.(int))
			}
			return Arr_mul_pixel_float32(arr1, arr_tmp)
		case float32:
			arr1 := Interface2Float32(arr1)
			arr_tmp := make([]float32, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = arr2.(float32)
			}
			return Arr_mul_pixel_float32(arr1, arr_tmp)
		case float64:
			arr1 := Interface2Float64(arr1)
			arr_tmp := make([]float64, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = float64(arr2.(float32))
			}
			return Arr_mul_pixel_float64(arr1, arr_tmp)
		}
	case []float64:
		switch arr2.(type) {
		case []int:
			arr1 := Interface2Float64(arr1)
			arr2 := Interface2Int(arr2)
			arr2_float_64 := Int2Float64(arr2)
			return Arr_mul_pixel_float64(arr1, arr2_float_64)
		case []float32:
			arr1 := Interface2Float64(arr1)
			arr2 := Interface2Float32(arr2)
			arr2_float_64 := Float322Float64(arr2)
			return Arr_mul_pixel_float64(arr1, arr2_float_64)
		case []float64:
			arr1 := Interface2Float64(arr1)
			arr2 := Interface2Float64(arr2)
			return Arr_mul_pixel_float64(arr1, arr2)
		case int:
			arr1 := Interface2Float64(arr1)
			arr_tmp := make([]float64, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = float64(arr2.(int))
			}
			return Arr_mul_pixel_float64(arr1, arr_tmp)
		case float32:
			arr1 := Interface2Float64(arr1)
			arr_tmp := make([]float64, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = float64(arr2.(float32))
			}
			return Arr_mul_pixel_float64(arr1, arr_tmp)
		case float64:
			arr1 := Interface2Float64(arr1)
			arr_tmp := make([]float64, value1.Len())
			for i := 0; i < value1.Len(); i++ {
				arr_tmp[i] = arr2.(float64)
			}
			return Arr_mul_pixel_float64(arr1, arr_tmp)
		}
	}
	panic("type error")
}
func Arr_mul_pixel_float32(arr1 []float32, arr2 []float32) []float32 {
	res := make([]float32, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i] * arr2[i]
	}
	return res
}

func Arr_mul_pixel_float64(arr1 []float64, arr2 []float64) []float64 {
	res := make([]float64, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i] * arr2[i]
	}
	return res
}

func Arr_mul_pixel_int(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		res[i] = arr1[i] * arr2[i]
	}
	return res
}
