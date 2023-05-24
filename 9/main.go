package main

import "strconv"

func main() {

	println(1321)
	println(isPalindrome(1321))
	println(-121)
	println(isPalindrome(-121))
	println(10)
	println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	s := strconv.FormatInt(int64(x), 10)
	for i, _ := range s {
		if len(s)-i-1 == 0 {
			break
		}

		if s[len(s)-i-1] != s[i] {
			return false
		}
	}
	return true
}
