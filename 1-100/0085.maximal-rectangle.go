/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	maxArea := 0
	heights := make([]int, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == '1' {
				heights[col]++
			} else {
				heights[col] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	leftFirstLessPos := make([]int, n)
	leftFirstLessPos[0] = -1
	for i := 1; i < n; i++ {
		left := i - 1
		for left >= 0 && heights[left] >= heights[i] {
			left = leftFirstLessPos[left]
		}
		leftFirstLessPos[i] = left
	}
	rightFirstLessPos := make([]int, n)
	rightFirstLessPos[n-1] = n
	for i := n - 2; i >= 0; i-- {
		right := i + 1
		for right < n && heights[right] >= heights[i] {
			right = rightFirstLessPos[right]
		}
		rightFirstLessPos[i] = right
	}
	maxArea := 0
	for i := range heights {
		area := heights[i] * (rightFirstLessPos[i] - leftFirstLessPos[i] - 1)
		maxArea = max(maxArea, area)
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
