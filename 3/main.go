package main

func main() {

	print("\n\n")
	println(lengthOfLongestSubstring("abcabcbb"), lengthOfLongestSubstring("abcabcbb") == 3)
	println(lengthOfLongestSubstring("bbbbb") == 1)
	println(lengthOfLongestSubstring("pwwkew") == 3)
	println(lengthOfLongestSubstring("tmmzuxt") == 5)
	print("\n\n")

}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int16, len(s))

	var begin int16 = 0
	var max int16 = 0

	for index, value := range s {

		pos, ok := m[value]

		if ok {
			if begin < pos+1 {
				begin = pos + 1
			}
		}

		length := int16(index) - begin + 1
		if length > max {
			max = length
		}

		m[value] = int16(index)

	}
	return int(max)
}
