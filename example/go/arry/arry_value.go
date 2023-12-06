package arry

import "fmt"

func Arr_value() {
	var arr1 [5]int
	arr2 := &arr1
	arr2[0] = 1
	arr1[1] = 1
	fmt.Println(arr1)
	fmt.Println(arr2)
}
