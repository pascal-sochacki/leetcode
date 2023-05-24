package main

import "strconv"

func main() {
	res := countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"})
	println(res)

}

func countSeniors(details []string) (result int) {

	for _, v := range details {
		vInt, _ := strconv.Atoi(v[11:13])
		if vInt > 60 {
			result += 1
		}
	}
	return
}
