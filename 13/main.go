package main

func main() {
	input := []string{"III", "LVIII", "MCMXCIV"}
	output := []int{3, 58, 1994}

	for i, s := range input {
		println(input[i], romanToInt(s), output[i] == romanToInt(s))
	}

}

func romanToInt(s string) int {

	result := 0

	for index := 0; index < len(s); index++ {

		current := s[index]

		var next uint8 = 0
		if index+1 < len(s) {
			next = s[index+1]
		}

		if current == 'I' {
			if next == 'V' {
				result += 4
				index++
			} else if next == 'X' {
				result += 9
				index++
			} else {
				result += 1
			}
		}
		if current == 'V' {
			result += 5
		}

		if current == 'X' {
			if next == 'L' {
				result += 40
				index++
			} else if next == 'C' {
				result += 90
				index++
			} else {
				result += 10
			}
		}

		if current == 'L' {
			result += 50
		}

		if current == 'C' {
			if next == 'D' {
				result += 400
				index++
			} else if next == 'M' {
				result += 900
				index++
			} else {
				result += 100
			}
		}

		if current == 'D' {
			result += 500
		}

		if current == 'M' {
			result += 1000
		}
	}

	return result

}
