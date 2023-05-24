package main

func main() {
	one := plusOne([]int{})
	one = plusOne([]int{9})
	one = plusOne([]int{9, 9})
	one = plusOne([]int{8})
	println(one)
	plusOne([]int{9})
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	if digits[len(digits)-1] == 9 {
		return append(plusOne(digits[:len(digits)-1]), 0)
	} else {
		digits[len(digits)-1]++
		return digits
	}

}
