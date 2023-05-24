package main

func main() {
	//println(isPalindrome(" "))
	//println(isPalindrome("ab"))
	println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	clean := ""
	for i := range s {

		if s[i] >= 'a' && s[i] <= 'z' {
			clean = clean + string(s[i])
		}

		if s[i] >= 'A' && s[i] <= 'Z' {
			clean = clean + string(s[i]+32)
		}
	}

	length := len(clean) - 1
	for i := range clean {

		if clean[i] != clean[length-i] {
			return false
		}
	}

	return true

}
