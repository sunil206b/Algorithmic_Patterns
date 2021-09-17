package main

import "fmt"

func main() {
	nums := []int{1, 2, 7, 12, 43, 44, 54, 100, 124}
	target := 124
	idx := binarySearch(nums, target)
	fmt.Println(idx)
}

func binarySearch(nums []int, target int) int {
	start := 0
	l := len(nums)
	end := l - 1
	for start <= end {
		midIdx := (start + end) / 2
		if nums[midIdx] == target {
			return midIdx
		}
		if nums[midIdx] < target {
			start = midIdx + 1
		} else {
			end = midIdx - 1
		}
	}
	return -1
}
