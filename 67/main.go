package main

func main() {
	println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {

	var i int8 = 0

	aLen := int8(len(a))
	bLen := int8(len(b))
	result := ""
	var carry int8 = 0
	var sum int8 = 0

	for {

		if aLen-i <= 0 && bLen-i <= 0 && carry == 0 {
			return result
		}

		var a_1 int8 = 0
		if aLen-i-1 >= 0 && a[aLen-i-1] == '1' {
			a_1 = 1
		}

		var b_1 int8 = 0
		if bLen-i-1 >= 0 && b[bLen-i-1] == '1' {
			b_1 = 1
		}

		sum, carry = add(a_1, b_1, carry)

		if sum == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}

		i++

	}

}

func add(a1 int8, b1 int8, c int8) (int8, int8) {

	sum := a1 + b1 + c
	if sum == 3 {
		return 1, 1
	} else if sum == 2 {
		return 0, 1
	} else if sum == 1 {
		return 1, 0
	} else {
		return 0, 0
	}

}

func (receiver) name() {

}
