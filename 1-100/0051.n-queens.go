/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */
package leetcode

// @lc code=start
// 每一行有且仅有一个皇后，且每行皇后的列号不同，
// 再令皇后们不在同一对角线，就可得出一种正解
func solveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}
	var ans [][]string
	cols := make([]int, n) // cols[i] 表示第i行皇后在cols[i]列
	vis := make([]bool, n) // vis记录已选列号
	dfs(&ans, cols, vis, 0)
	return ans
}

func dfs(ans *[][]string, cols []int, vis []bool, curRow int) {
	n := len(vis)
	if curRow == n {
		*ans = append(*ans, newBoard(cols))
		return
	}
	for curCol := 0; curCol < n; curCol++ {
		if !vis[curCol] {
			valid := true
			// 遍历之前已选的皇后，检查是否和当前皇后在一条对角线上
			for oldRow := 0; oldRow < curRow; oldRow++ {
				oldCol := cols[oldRow]
				if curRow-oldRow == abs(curCol-oldCol) {
					valid = false
					break
				}
			}
			if valid {
				cols[curRow] = curCol
				vis[curCol] = true
				dfs(ans, cols, vis, curRow+1)
				vis[curCol] = false
			}
		}
	}
}

// 根据cols生成棋图
func newBoard(cols []int) []string {
	n := len(cols)
	board := make([]string, 0, n)
	for i := 0; i < n; i++ {
		buf := make([]byte, n)
		for j := 0; j < n; j++ {
			if cols[i] == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		board = append(board, string(buf))
	}
	return board
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// @lc code=end
