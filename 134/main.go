package main

import (
	"fmt"
)

func main() {

	//Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
	//Output: 3
	println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	println("===============")
	println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	println("===============")

	println(canCompleteCircuit([]int{1, 2, 3, 4, 3, 2, 4, 1, 5, 3, 2, 4}, []int{1, 1, 1, 3, 2, 4, 3, 6, 7, 4, 3, 1}))
}

func canCompleteCircuit(gas []int, cost []int) int {

	start := 0
	i := 1
	tank := gas[0]
	for {

		if i > 100 {
			return -1
		}
		if i%len(cost) > 0 {
			tank = tank - cost[i%(len(cost))-1]
		} else {
			tank = tank - cost[len(cost)-1]
		}

		println(fmt.Sprintf("Tank at %d", tank))
		if start == i%len(gas) {
			if tank >= 0 {
				return start
			} else {
				return -1
			}
		}

		if tank < 0 {
			println(fmt.Sprintf("negativ tank at %d", i))
			start = i % (len(gas))
			tank = gas[i%len(gas)]
		} else {
			println(fmt.Sprintf("positiv tank at %d with start at %d", i, start))
			tank = gas[i%len(gas)] + tank
		}

		i += 1

	}
}
