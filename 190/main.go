package main

import "strconv"

func main() {

	result := reverseBits(43261596)

	println(strconv.FormatInt(int64(43261596), 2))
	println(strconv.FormatInt(int64(result), 2))
	println(strconv.FormatInt(int64(964176192), 2))
	println(result == 964176192)
}

func reverseBits(num uint32) (result uint32) {

	for i := 0; i < 32; i++ {

		var set uint32 = (1 << i) & num

		var add uint32 = 0
		if i < 16 {
			add = set << (31 - (i * 2))

		} else {
			add = set >> ((i-16)*2 + 1)
		}
		result |= add

	}
	return result
}
