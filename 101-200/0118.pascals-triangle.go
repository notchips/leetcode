/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
	}
	for row := 0; row < numRows; row++ {
		for col := 0; col < len(triangle[row]); col++ {
			if col == 0 || col == len(triangle[row]) -1 {
				triangle[row][col] = 1
			} else {
				triangle[row][col] = triangle[row-1][col-1] + triangle[row-1][col]
			}
		}
	}
	return triangle
}
