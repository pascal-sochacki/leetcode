package main

import "strings"

func main() {

	println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {

	split := strings.Split(strings.TrimSpace(s), " ")
	return len(split[len(split)-1])
}
