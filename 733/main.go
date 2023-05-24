package main

func main() {

	result := floodFill(
		[][]int{
			{0, 0, 0},
			{0, 0, 0},
		},
		1,
		0,
		2,
	)
	println(result)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	origin := image[sr][sc]

	if origin == color {
		return image
	}

	image[sr][sc] = color

	if len(image) > sr+1 && image[sr+1][sc] == origin {
		image = floodFill(image, sr+1, sc, color)
	}

	if 0 <= sr-1 && image[sr-1][sc] == origin {
		image = floodFill(image, sr-1, sc, color)
	}

	if len(image[0]) > sc+1 && image[sr][sc+1] == origin {
		image = floodFill(image, sr, sc+1, color)
	}

	if 0 <= sc-1 && image[sr][sc-1] == origin {
		image = floodFill(image, sr, sc-1, color)
	}

	return image
}
