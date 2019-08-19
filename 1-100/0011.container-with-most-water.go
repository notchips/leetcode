/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, area)
		// 证明遍历到了最大值：
		// 1.令v[i, j] 表示以i，j范围的面积；
		// 2.当 height[i] < height[j]时，i++，遍历v[i+1, j]。
		//   也就是说忽略了v[i, j-1], v[i, j-2]...的情况，
		//   而被忽略的情况由于底是小于(j-i)的，最大高度被限制为height[i]。
		//   因此被忽略的情况不包含最大值，遍历过程中肯定遍历到了最大值
		// 3.同理可推 height[i] >= height[j] 的情形
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
