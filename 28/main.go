package main

func main() {
	//print(strStr("sadbutsad", "sad"))
	//print(strStr("leetcode", "leeto"))
	print(strStr("a", "a"))
}

func strStr(haystack string, needle string) int {

	for i := 0; i <= (len(haystack) - len(needle)); i++ {

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1

}
