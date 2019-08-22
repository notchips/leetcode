/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
func isValidSudoku(board [][]byte) bool {
	var columns, rows, squares [9][9]bool // record filled nums
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := board[row][col] - '0' - 1
				if columns[col][num] || rows[row][num] || squares[row/3*3+col/3][num] {
					return false
				}
				columns[col][num], rows[row][num], squares[row/3*3+col/3][num] = true, true, true
			}
		}
	}
	return true
}
