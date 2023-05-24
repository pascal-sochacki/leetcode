package main

import "strings"

func main() {

}

func mostWordsFound(sentences []string) int {

	max := 0
	for _, sentence := range sentences {
		curr := strings.Count(sentence, "")
		if curr != 0 {
			curr += 1
		}

		if curr > max {
			max = curr
		}
	}

	return max

}
