package main

func main() {
	println(minTimeToVisitAllPoints([][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}))
}
func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	for i := 0; i < len(points)-1; i++ {

		yDelta := points[i][0] - points[i+1][0]
		if yDelta < 0 {
			yDelta = -1 * yDelta
		}

		xDelta := points[i][1] - points[i+1][1]
		if xDelta < 0 {
			xDelta = -1 * xDelta
		}

		if xDelta > yDelta {
			result += xDelta
		} else {
			result += yDelta
		}
	}
	return result

}
