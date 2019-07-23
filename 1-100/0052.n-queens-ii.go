/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
// 只要每一行皇后的列号不一样，就能保证皇后们不在一条线上
// 最后只需要满足皇后们不在同一对角线，就可得出一种正解
func totalNQueens(n int) int {
	cnt := 0
	vis := make([]bool, n)    // vis记录已选列号
	ans := make([]int, n)     // ans记录正解，ans[i]=j：表示第i行皇后的列号为j
	dfs(&cnt, ans, vis, n, 0) // 从第0行开始，为每一行皇后确定列号
	return cnt
}

func dfs(cnt *int, ans []int, vis []bool, n, row int) {
	if row >= n {
		*cnt++
		return
	}

	for col := 0; col < n; col++ { // 选择一个未选列号
		if !vis[col] {
			valid := true

			// 遍历之前已选的皇后，检查是否和当前皇后在一条对角线上
			for r := 0; r < row; r++ {
				if abs(col-ans[r]) == abs(row-r) {
					valid = false
					break
				}
			}

			if valid {
				vis[col] = true
				ans[row] = col
				dfs(cnt, ans, vis, n, row+1)
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

