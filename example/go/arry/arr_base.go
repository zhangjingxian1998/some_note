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

func Arr_sum(arrs interface{}) {
	switch arrs.(type) {
	case []int:
		arrs := Interface2Int(arrs)
		fmt.Println(Arr_sum_int(arrs))
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}
	fmt.Println(reflect.TypeOf(arrs))

}
func Interface2Int(arr interface{}) []int {
	var res []int
	value := reflect.ValueOf(arr)
	for i := 0; i < value.Len(); i++ {
		res = append(res, value.Index(i).Interface().(int))
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
