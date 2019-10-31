/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */
package leetcode

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	i := 0
	iota := func() int {
		i++
		return i
	}
	for start, end := 0, n-1; start <= end; start, end = start+1, end-1 {
		for col := start; col <= end; col++ {
			matrix[start][col] = iota()
		}
		for row := start + 1; row < end; row++ {
			matrix[row][end] = iota()
		}
		if start < end {
			for col := end; col >= start; col-- {
				matrix[end][col] = iota()
			}
			for row := end - 1; row > start; row-- {
				matrix[row][start] = iota()
			}
		}
	}
	return matrix
}

// @lc code=end
