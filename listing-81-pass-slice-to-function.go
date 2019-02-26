package main

import "fmt"

func doubleInts(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	doubleInts(nums)
	fmt.Println(nums)
}
