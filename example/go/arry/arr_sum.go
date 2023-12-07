package arry

import (
	"fmt"
	"reflect"
)

func Type_test(args interface{}) {
	switch args.(type) {
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}

}

func Arr_sum(arrs interface{}) interface{} {
	value := reflect.ValueOf(arrs)
	res := value.Index(0).Interface()
	switch arrs.(type) {
	case []int:
		res = res.(int)
		arrs := Interface2Int(arrs)
		res = Arr_sum_int(arrs)
	case []float32:
		res = res.(float32)
		arrs := Interface2Float32(arrs)
		res = Arr_sum_float32(arrs)
	case []float64:
		res = res.(float64)
		arrs := Interface2Float64(arrs)
		res = Arr_sum_float64(arrs)
	}
	return res
}

func Arr_sum_int(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Arr_sum_float32(arr []float32) float32 {
	var sum float32
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Arr_sum_float64(arr []float64) float64 {
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
