package main

import "strings"

func main() {
	println(canBeTypedWords("hello world", "ad"))
}

func canBeTypedWords(text string, brokenLetters string) (result int) {

	set := make(map[uint8]bool, len(brokenLetters))

	for i := range brokenLetters {
		set[brokenLetters[i]] = true
	}

	split := strings.Split(text, " ")

	for _, word := range split {
		valid := true
		for j := range word {

			_, exists := set[word[j]]

			if exists {
				valid = false
				break
			}

		}

		if valid {
			result++
		}
	}

	return result

}
