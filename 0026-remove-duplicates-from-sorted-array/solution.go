package _026_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	s := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[s] = nums[i]
			s++
		}
	}
	return s
}
