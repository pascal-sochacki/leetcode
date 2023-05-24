package main

func main() {
	println(convertToTitle(1))
	println(convertToTitle(26))
	println(convertToTitle(27))
}

func convertToTitle(columnNumber int) string {

	temp := columnNumber

	result := ""

	var base = 26

	for temp != 0 {
		temp--
		remain := temp % base
		temp = temp / base

		s := string(rune('A' + remain))
		result = s + result

	}

	return result
}
