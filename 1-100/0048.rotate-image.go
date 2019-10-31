/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
package leetcode

// @lc code=start
func rotate(matrix [][]int) {
	for start, end := 0, len(matrix)-1; start < end; start, end = start+1, end-1 {
		for i := start; i < end; i++ {
			matrix[i][start], matrix[end][i] = matrix[end][i], matrix[i][start]
			j := start + end - i
			matrix[end][i], matrix[j][end] = matrix[j][end], matrix[end][i]
			matrix[j][end], matrix[start][j] = matrix[start][j], matrix[j][end]
		}
	}
}

// @lc code=end
