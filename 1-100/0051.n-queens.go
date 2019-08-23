/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */
// 每一行有且仅有一个皇后，且每行皇后的列号不同，
// 再令皇后们不在同一对角线，就可得出一种正解
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0, 32)
	vis := make([]bool, n)     // vis记录已选列号
	cols := make([]int, n)     // cols[i]：表示第i行皇后的列号为cols[i]
	dfs(&ans, cols, vis, 0, n) // 从第0行开始，为每一行皇后确定列号
	return ans
}

func dfs(ans *[][]string, cols []int, vis []bool, row, n int) {
	if row == n {
		*ans = append(*ans, getSolution(cols))
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
				dfs(ans, cols, vis, row+1, n)
				vis[col] = false
			}
		}
	}
}

// 根据cols生成棋图
func getSolution(cols []int) []string {
	answer := make([]string, len(cols))
	for i := 0; i < len(cols); i++ {
		buf := make([]byte, len(cols))
		for j := 0; j < len(buf); j++ {
			if j == cols[i] {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		answer[i] = string(buf)
	}
	return answer
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
