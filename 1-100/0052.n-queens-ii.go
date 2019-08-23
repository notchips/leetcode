/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
// 每一行有且仅有一个皇后，且每行皇后的列号不同，
// 再令皇后们不在同一对角线，就可得出一种正解
func totalNQueens(n int) int {
	cnt := 0
	vis := make([]bool, n)     // vis记录已选列号
	cols := make([]int, n)     // cols[i]：表示第i行皇后的列号为cols[i]
	dfs(&cnt, cols, vis, 0, n) // 从第0行开始，为每一行皇后确定列号
	return cnt
}

func dfs(cnt *int, cols []int, vis []bool, row, n int) {
	if row == n {
		*cnt++
		return
	}
	for col := 0; col < n; col++ {
		if !vis[col] {
			valid := true
			// 遍历之前已选的皇后，检查是否和当前皇后在一条对角线上
			for preRow := 0; preRow < row; preRow++ {
				if abs(col-cols[preRow]) == abs(row-preRow) {
					valid = false
					break
				}
			}
			if valid {
				vis[col] = true
				cols[row] = col
				dfs(cnt, cols, vis, row+1, n)
				vis[col] = false
			}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
