package _249_minimum_remove_to_make_valid_parentheses

import "strings"

func MinRemoveToMakeValid(s string) string {
	arr := []rune(s)
	open := 0
	for idx, r := range s {
		if r == '(' {
			open++
		} else if r == ')' {
			if open == 0 {
				arr[idx] = '*'
			} else {
				open--
			}
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if open > 0 && arr[i] == '(' {
			arr[i] = '*'
			open--
		}
	}

	stringBuilder := strings.Builder{}
	for _, r := range arr {
		if r != '*' {
			stringBuilder.WriteRune(r)
		}
	}
	return stringBuilder.String()
}
