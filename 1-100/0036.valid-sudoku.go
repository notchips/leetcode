/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
package leetcode

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var rows, columns, squares [9][9]bool // record filled nums
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				id := int(board[row][col] - '0' - 1)
				if rows[row][id] || columns[col][id] || squares[row/3*3+col/3][id] {
					return false
				}
				rows[row][id], columns[col][id], squares[row/3*3+col/3][id] = true, true, true
			}
		}
	}
	return true
}

// @lc code=end
