package main

import (
	"math"
)

func maxAverageSubArray(nums []int, k int) float64 {
	var windowSum int
	var start int
	max := math.Inf(-1)

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]
		if (end - start + 1) == k {
			max = math.Max(max, float64(windowSum)/float64(k))
			windowSum -= nums[start]
			start += 1
		}
	}
	return max
}