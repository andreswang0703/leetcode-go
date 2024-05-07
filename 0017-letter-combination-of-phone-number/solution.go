package _017_letter_combination_of_phone_number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numToStr := map[rune]string{
		'0': "",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	curString := []rune{}
	res := []string{}
	dfs(numToStr, &res, &curString, digits, 0)
	return res
}

func dfs(numToStr map[rune]string, res *[]string, curString *[]rune, digits string, idx int) {
	if idx == len(digits) {
		*res = append(*res, string(*curString))
		return
	}
	num := rune(digits[idx])
	numStr := numToStr[num]
	for _, c := range numStr {
		*curString = append(*curString, c)
		dfs(numToStr, res, curString, digits, idx+1)
		*curString = (*curString)[:len(*curString)-1]
	}
}
