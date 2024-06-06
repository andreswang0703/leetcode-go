package _918_maximum_sum_circular_subarray

func maxSubarraySumCircular(nums []int) int {
	curMax := 0
	maxSum := nums[0]

	curMin := 0
	minSum := nums[0]

	totalSum := 0

	for _, num := range nums {
		curMax = max(0, curMax) + num
		maxSum = max(maxSum, curMax)

		curMin = min(0, curMin) + num
		minSum = min(minSum, curMin)

		totalSum += num
	}

	// edge case: when min subarray includes all elements (e.g. all negative numbers)
	// the reverse of min subarray which is what we return as a max subarray would be empty, and it's not allowed.
	if totalSum == minSum {
		return maxSum
	}

	return max(totalSum-minSum, maxSum)
}
