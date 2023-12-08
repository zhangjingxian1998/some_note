package main

import (
	"fmt"
	"go_project/arry"
	"go_project/fibonacci"
	"go_project/print_practice"
	"math/rand"
)

func main() {
	print_practice.Whole_show()
	fibonacci.Whole_show()
	arry.Whole_show()
	Linear_calculate()
}

func Linear_calculate() {
	// 生成数据集
	size := 100
	m := 2
	b := 5
	x_data := arry.Int2Float32(arry.Arr_arange(100))
	y_data := arry.Arr_mul_pixel(x_data, m)
	y_data = arry.Arr_add_pixel(y_data, b)
	noise_tmp := arry.Arr_random_noise(100, 0, 1.5)
	noise := arry.Float642Float32(noise_tmp.([]float64))

	y_data = arry.Arr_add_pixel(y_data, noise)

	// 设置超参数
	var epochs, batch_size int
	var lr float32
	epochs = 1000000
	lr = 0.0005
	batch_size = 100

	// 建立模型
	w := rand.Float32()
	bias := rand.Float32()

	var y, grad_w, grad_b, grad_w_sum, grad_b_sum float32
	var batch_count int
	for epoch := 0; epoch < epochs; epoch++ {
		batch_count = 0
		grad_w_sum = 0
		grad_b_sum = 0
		for i := 0; i < size; i++ {
			y = w*x_data[i] + bias
			grad_w = (y_data.([]float32)[i] - y) * (-1) * x_data[i]
			grad_b = (y_data.([]float32)[i] - y) * (-1)
			batch_count += 1
			if batch_count >= batch_size {
				grad_w_sum += grad_w
				grad_b_sum += grad_b_sum
				w = w - lr*grad_w_sum/float32(batch_size)
				bias = bias - lr*grad_b_sum/float32(batch_size)
				grad_w_sum = 0
				grad_b_sum = 0
				batch_count = 0
			} else {
				grad_w_sum += grad_w
				grad_b_sum += grad_b
			}
			if epoch%10000 == 0 {
				fmt.Println(w, bias)
			}

		}
	}
}
