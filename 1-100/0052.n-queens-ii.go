/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
package leetcode

// @lc code=start
// 每一行有且仅有一个皇后，且每行皇后的列号不同，
// 再令皇后们不在同一对角线，就可得出一种正解
func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	cnt := 0
	cols := make([]int, n) // cols[i] 表示第i行皇后在cols[i]列
	vis := make([]bool, n) // vis记录已选列号
	dfs(&cnt, cols, vis, 0)
	return cnt
}

func dfs(cnt *int, cols []int, vis []bool, curRow int) {
	n := len(vis)
	if curRow == n {
		*cnt++
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
				dfs(cnt, cols, vis, curRow+1)
				vis[curCol] = false
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// @lc code=end
