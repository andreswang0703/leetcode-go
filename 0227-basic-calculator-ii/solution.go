package _227_basic_calculator_ii

import "unicode"

func Calculate(s string) int {
	arr := []rune(s)
	cur := 0
	op := '+'
	prev := 0
	res := 0

	for i, c := range arr {
		if unicode.IsDigit(c) {
			cur = cur*10 + int(c-'0')
		}

		if !unicode.IsDigit(c) && c != ' ' || i == len(s)-1 { // pay attention to empty space
			if op == '+' || op == '-' {
				res += prev
				prev = cur
				if op == '-' {
					prev = -prev
				}
			} else if op == '*' || op == '/' {
				prev = operate(prev, cur, op)
			}

			op = c
			cur = 0
		}
	}
	return res + prev
}

func operate(a int, b int, op rune) int {
	if op == '+' {
		return a + b
	}
	if op == '-' {
		return a - b
	}
	if op == '*' {
		return a * b
	}
	if op == '/' {
		return a / b
	}
	return -1
}

// example to go through
// 3 - 5 * 6 + 13 * 7
//       ^ add prev (3) to res, also setting prev = -5
// 3 - 5 * 6 + 13 * 7
//           ^ (op == * or /) update prev
// 3 - 5 * 6 + 13 * 7
//                ^  (op == +/-) add 30 (prev) to res
