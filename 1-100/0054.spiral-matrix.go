/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)
	travelEdge(&ans, matrix)
	return ans
}

func travelEdge(ans *[]int, matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 从左往右，遍历上行
	for i := 0; i < n; i++ {
		*ans = append(*ans, matrix[0][i])
	}
	// 从上往下，遍历右列
	for j := 1; j < m-1; j++ {
		*ans = append(*ans, matrix[j][n-1])
	}
	// 从右往左，遍历下行
	if m > 1 {
		for i := n - 1; i >= 0; i-- {
			*ans = append(*ans, matrix[m-1][i])
		}
	}
	// 从下往上，遍历左列
	if n > 1 {
		for j := m - 2; j >= 1; j-- {
			*ans = append(*ans, matrix[j][0])
		}
	}
	// 内部还存在数字
	if m > 2 && n > 2 {
		matrix = matrix[1 : m-1]
		for i := range matrix {
			matrix[i] = matrix[i][1 : n-1]
		}
		travelEdge(ans, matrix)
	}
}
