package _921_minimum_add_to_make_parentheses_valid

func MinAddToMakeValid(s string) int {
	open := 0
	ans := 0
	for _, r := range s {
		if r == '(' {
			open++
		} else if r == ')' {
			if open == 0 {
				ans++
			} else {
				open--
			}
		}
	}
	return ans + open
}
