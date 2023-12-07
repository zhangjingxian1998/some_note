package arry

import (
	"math"
)

func Arr_pow_pixel(arr interface{}, pow int) interface{} {
	switch arr.(type) {
	case []int:
		arrs := Interface2Int(arr)
		return Arr_pow_pixel_int(arrs, pow)
	case []float32:
		arrs := Interface2Float32(arr)
		return Arr_pow_pixel_float32(arrs, pow)
	case []float64:
		arrs := Interface2Float64(arr)
		return Arr_pow_pixel_float64(arrs, pow)
	default:
		panic("Only accept []int []float32 []float64 at first position and int at second position")
	}
}

func Arr_pow_pixel_int(arr []int, pow int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = int(math.Pow(float64(arr[i]), float64(pow)))
	}
	return arr
}

func Arr_pow_pixel_float32(arr []float32, pow int) []float32 {
	for i := 0; i < len(arr); i++ {
		arr[i] = float32(math.Pow(float64(arr[i]), float64(pow)))
	}
	return arr
}

func Arr_pow_pixel_float64(arr []float64, pow int) []float64 {
	for i := 0; i < len(arr); i++ {
		arr[i] = float64(math.Pow(float64(arr[i]), float64(pow)))
	}
	return arr
}
