package arry

func Arr_normalize(arr interface{}) interface{} {
	arr_norm := Arr_norm(arr)
	return Arr_dev(arr, arr_norm)
}
