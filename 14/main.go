package main

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
func longestCommonPrefix(strs []string) string {

	var i uint8 = 0

	result := ""
	for {

		for j := range strs {
			if uint8(len(strs[j])) == i || strs[0][i] != strs[j][i] {
				return result
			}
		}
		result += string(strs[0][i])
		i++
	}

}
