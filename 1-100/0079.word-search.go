/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
func exist(board [][]byte, word string) bool {
	if word == "" || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if check(board, word, 0, y, x) {
				return true
			}
		}
	}
	return false
}

func check(board [][]byte, word string, i, y, x int) bool {
	if i == len(word) {
		return true
	}
	if y < 0 || y == len(board) || x < 0 || x == len(board[0]) || board[y][x] != word[i] {
		return false
	}
	// 利用ascii码只使用了0-127的特性，令最高位为1，代表board[y][x]已选择（肯定不等于word中的字符）。
	board[y][x] ^= 1 << 7
	ret := check(board, word, i+1, y+1, x) || check(board, word, i+1, y-1, x) ||
		check(board, word, i+1, y, x+1) || check(board, word, i+1, y, x-1)
	board[y][x] ^= 1 << 7 // 还原
	return ret
}
