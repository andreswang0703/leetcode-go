package _009_palindrome_number

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	res := 0
	for x > 0 {
		remainder := x % 10
		x /= 10
		res = res*10 + remainder
	}
	return res
}
