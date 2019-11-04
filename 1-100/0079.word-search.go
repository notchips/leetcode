/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
package leetcode

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	m, n := len(board), len(board[0])
	if i < 0 || i >= m ||
		j < 0 || j >= n ||
		word[0] != board[i][j] {
		return false
	}
	// 利用ascii码只使用了0-127的特性，令最高位为1，代表board[y][x]已选择
	board[i][j] ^= 1 << 7
	ret := search(board, word[1:], i+1, j) ||
		search(board, word[1:], i-1, j) ||
		search(board, word[1:], i, j+1) ||
		search(board, word[1:], i, j-1)
	board[i][j] ^= 1 << 7
	return ret
}

// @lc code=end
