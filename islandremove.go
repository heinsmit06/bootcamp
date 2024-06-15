package bootcamp

func dfs(matrix [][]int, x, y int) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] == 0 {
		return
	}

	matrix[x][y] = 0
	dfs(matrix, x+1, y)
	dfs(matrix, x-1, y)
	dfs(matrix, x, y+1)
	dfs(matrix, x, y-1)
}

func IslandRemove(matrix [][]int, x, y int) {
	if matrix == nil {
		return
	}

	dfs(matrix, y, x)
}
