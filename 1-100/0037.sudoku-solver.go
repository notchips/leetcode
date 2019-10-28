/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
package leetcode

// @lc code=start
func solveSudoku(board [][]byte) {
	var rows, columns, squares [9][9]bool
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				id := int(board[row][col] - '0' - 1)
				rows[row][id], columns[col][id], squares[row/3*3+col/3][id] = true, true, true
			}
		}
	}
	filled := false
	fillBoard(board, &rows, &columns, &squares, &filled)
}

func fillBoard(board [][]byte, rows, columns, squares *[9][9]bool, filled *bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for id := 0; id < 9; id++ {
					if !rows[row][id] && !columns[col][id] && !squares[row/3*3+col/3][id] {
						rows[row][id], columns[col][id], squares[row/3*3+col/3][id] = true, true, true
						board[row][col] = byte(id + 1 + '0')
						fillBoard(board, rows, columns, squares, filled)
						if *filled == true { return } // 填写正确，终止所有递归栈
						board[row][col] = '.'
						rows[row][id], columns[col][id], squares[row/3*3+col/3][id] = false, false, false
					}
				}
				return // 填写错误
			}
		}
	}
	*filled = true // 填写正确
}

// @lc code=end
