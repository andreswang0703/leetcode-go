package _001_two_sum

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	indexMap := make(map[int]int)
	for idx, n := range nums {
		indexMap[n] = idx
	}
	for idx, n := range nums {
		remain := target - n
		v, exists := indexMap[remain]
		if exists && v != idx {
			res[0] = idx
			res[1] = v
			return res
		}
	}
	return make([]int, 0)
}
