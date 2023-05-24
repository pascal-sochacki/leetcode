package main

func main() {
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
}

func isValid(s string) bool {
	stack := make([]uint8, len(s))

	pointer := -1
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			pointer++
			stack[pointer] = s[i]
		} else {

			if pointer < 0 {
				return false
			}

			last := stack[pointer]
			pointer--

			if s[i] == ')' && last == '(' {
				continue
			}

			if s[i] == '}' && last == '{' {
				continue
			}

			if s[i] == ']' && last == '[' {
				continue
			}

			return false

		}
	}

	return pointer == -1

}
