package bootcamp

func ValidSudoku(n [][]int) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		if !isValidGroup(n[i]) {
			return false
		}
	}

	// Check columns
	for i := 0; i < 9; i++ {
		column := make([]int, 9)
		for j := 0; j < 9; j++ {
			column[j] = n[j][i]
		}
		if !isValidGroup(column) {
			return false
		}
	}

	// Check 3x3 sub-grids
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidSubGrid(n, i, j) {
				return false
			}
		}
	}

	return true
}

func isValidGroup(group []int) bool {
	seen := make(map[int]bool)
	for _, num := range group {
		if num != 0 {
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}
	return true
}

func isValidSubGrid(board [][]int, startRow, startCol int) bool {
	seen := make(map[int]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := board[startRow+i][startCol+j]
			if num != 0 {
				if seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}
	return true
}
