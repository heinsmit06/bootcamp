package bootcamp

func dfsCost(matrix [][]int, x, y int, cost *int) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] == 0 {
		return
	}

	*cost += matrix[x][y]
	matrix[x][y] = 0
	dfsCost(matrix, x+1, y, cost)
	dfsCost(matrix, x-1, y, cost)
	dfsCost(matrix, x, y+1, cost)
	dfsCost(matrix, x, y-1, cost)
}

func IslandCost(matrix [][]int, x, y int) int {
	if matrix == nil {
		return 0
	}

	cost := 0
	dfsCost(matrix, y, x, &cost)
	return cost
}
