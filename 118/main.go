package main

func main() {
	i := generate(4)
	println(i)
}

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	i := generate(numRows - 1)

	last := i[numRows-2]

	ints := make([]int, numRows)

	for i := 0; i < numRows; i++ {

		if i == 0 || i == numRows-1 {
			ints[i] = 1
		} else {
			ints[i] = last[i-1] + last[i]
		}

	}

	return append(i, ints)

}
