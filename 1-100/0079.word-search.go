/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if search(board, word, row, col, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, row, col, k int) bool {
	if k == len(word) {
		return true
	}
	if row < 0 || row >= len(board) ||
		col < 0 || col >= len(board[0]) ||
		board[row][col] != word[k] {
		return false
	}
	// 利用ascii码只使用了0-127的特性，令最高位为1，代表board[y][x]已选择
	board[row][col] ^= 1 << 7
	ret := search(board, word, row+1, col, k+1) ||
		search(board, word, row-1, col, k+1) ||
		search(board, word, row, col+1, k+1) ||
		search(board, word, row, col-1, k+1)
	board[row][col] ^= 1 << 7 // 还原
	return ret
}
