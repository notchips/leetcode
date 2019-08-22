/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
func solveSudoku(board [][]byte) {
	var columns, rows, squares [9][9]bool // record filled nums
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := board[row][col] - '0' - 1
				columns[col][num], rows[row][num], squares[row/3*3+col/3][num] = true, true, true
			}
		}
	}
	finish := false
	fill(board, &columns, &rows, &squares, &finish)
}

func fill(board [][]byte, columns, rows, squares *[9][9]bool, finish *bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := 0; num < 9; num++ {
					if !columns[col][num] && !rows[row][num] && !squares[row/3*3+col/3][num] {
						board[row][col] = byte(num) + 1 + '0'
						columns[col][num], rows[row][num], squares[row/3*3+col/3][num] = true, true, true
						fill(board, columns, rows, squares, finish)
						if *finish { return }
						board[row][col] = '.'
						columns[col][num], rows[row][num], squares[row/3*3+col/3][num] = false, false, false
					}
				}
				return // fill wrong
			}
		}
	}
	*finish = true // fill right
}
