package tetora

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello, Tetora!")
}

func Map(arr []int, callback func(e int) int) []int {
	var res_arr []int
	for _, val := range arr {
		res := callback(val)
		res_arr = append(res_arr, res)
	}
	return res_arr
}