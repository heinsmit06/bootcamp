package bootcamp

func IslandExtend(matrix [][]int, x, y int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// Check if (x, y) is within the bounds of the matrix
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return false
	}

	// Check if the cell at (x, y) is already 1
	if matrix[x][y] == 1 {
		return false
	}

	// Define the directions for adjacent cells (up, down, left, right)
	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	// Check if any adjacent cell is 1
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && matrix[newX][newY] == 1 {
			matrix[x][y] = 1
			return true
		}
	}

	// If no adjacent cell is 1, return false
	return false
}
