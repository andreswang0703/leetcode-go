package _053_maximum_subarray

// This implements Kadane's algorithm with one dimension DP.
// At each index, local_maximum[i] = max(A[i], A[i] + local_maximum[i-1])
func maxSubArray(nums []int) int {
	maximum := make([]int, len(nums))
	maximum[0] = nums[0]
	maxSub := nums[0]

	for i := 1; i < len(nums); i++ {
		maximum[i] = max(nums[i], maximum[i-1]+nums[i])
		maxSub = max(maxSub, maximum[i])
	}
	return maxSub
}
