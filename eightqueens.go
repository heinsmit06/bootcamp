package bootcamp

func EightQueens() [][]int {
	var solutions [][]int
	board := make([]int, 8)

	isValid := func(row, col int) bool {
		for r := 0; r < row; r++ {
			c := board[r]
			if c == col || r-c == row-col || r+c == row+col {
				return false
			}
		}
		return true
	}

	convertToSolution := func(board []int) []int {
		solution := make([]int, 8)
		for i := 0; i < 8; i++ {
			solution[i] = board[i] + 1
		}
		return solution
	}

	var placeQueen func(int)
	placeQueen = func(row int) {
		if row == 8 {
			solutions = append(solutions, convertToSolution(board))
			return
		}

		for col := 0; col < 8; col++ {
			if isValid(row, col) {
				board[row] = col
				placeQueen(row + 1)
			}
		}
	}

	placeQueen(0)

	return solutions
}
