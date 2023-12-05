package print_practice

import "fmt"

func Print_int(a int) {
	fmt.Printf("输入的整型数是%d\n", a)
}
func Print_float(a float32) {
	fmt.Printf("输入的浮点数是%f\n", a)
	fmt.Printf("输入的浮点数保留两位小数是%.2f\n", a)
}
func Print_string(a string) {
	fmt.Printf("输入的字符串是%s\n", a)
}
func Print_int_float_string(a int, b float32, c string) {
	fmt.Printf("输入的整型数是%d,输入的浮点数是%f,输入的字符串是%s\n", a, b, c)
}
