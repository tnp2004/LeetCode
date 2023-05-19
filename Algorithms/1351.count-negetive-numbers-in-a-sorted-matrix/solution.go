package countnegetivenumbersinasortedmatrix

func countNegatives(grid [][]int) int {
	var c int

	for _, a := range grid {
		for _, v := range a {
			if v < 0 {
				c++
			}
		}
	}

	return c
}
