package main

import "fmt"

func main() {
	nums := []int{2, 1, 3, 4, 50, 78, 32, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(mergeSort(nums))
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left := mergeSort(nums[0 : len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	leftIdx := 0
	rightIdx := 0

	nIdx := 0
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			result[nIdx] = left[leftIdx]
			leftIdx++
		} else {
			result[nIdx] = right[rightIdx]
			rightIdx++
		}
		nIdx++
	}

	for leftIdx < len(left) {
		result[nIdx] = left[leftIdx]

		leftIdx++
		nIdx++
	}

	for rightIdx < len(right) {
		result[nIdx] = right[rightIdx]
		rightIdx++
		nIdx++
	}
	return result
}
