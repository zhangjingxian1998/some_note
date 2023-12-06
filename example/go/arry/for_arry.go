package arry

import "fmt"

func For_arry() {
	var arr1 [5]int
	arr2 := [...]int{0, 0, 0, 0, 0}
	arr3 := []int{0, 0, 0, 0, 0}
	arr4 := make([]int, 5)
	arr5 := new([5]int)
	for i := 0; i < 5; i++ {
		arr1[i] = i
		arr2[i] = i
		arr3[i] = i
		arr4[i] = i
		arr5[i] = i
	}
	fmt.Println("定长声明", arr1)
	fmt.Println("初始化声明1", arr2)
	fmt.Println("初始化声明2", arr3)
	fmt.Println("make声明", arr4)
	fmt.Println("new声明", arr5)
}
