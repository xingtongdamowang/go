package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice, _ = deleteSlice1(slice, 3)
	slice, _ = deleteSlice2(slice, 3)
	fmt.Printf("删除元素后的数组为 %+v\n", slice)
}
