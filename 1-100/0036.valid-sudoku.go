/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
func isValidSudoku(board [][]byte) bool {
	set := make(map[string]struct{}, 10)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c >= '1' && c <= '9' {
				rowKey := fmt.Sprintf("%cr%d", c, i)
				columnKey := fmt.Sprintf("%cc%d", c, j)
				squareKey := fmt.Sprintf("%cs%d%d", c, i/3, j/3)
				if !check(rowKey, set) || !check(columnKey, set) || !check(squareKey, set) {
					return false
				}
			}
		}
	}
	return true
}

func check(key string, set map[string]struct{}) bool {
	if _, ok := set[key]; ok {
		return false
	}
	set[key] = struct{}{}
	return true
}
