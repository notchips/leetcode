/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package leetcode

// 证明遍历到了最大值：
// 1.令v[l, r] 表示以l，r范围的面积；
// 2.当 height[r] < height[r]时，l++，遍历v[l+1, j]。
//   也就是说忽略了v[l, r-1], v[l, r-2]...的情况，
//   而被忽略的情况由于底是小于(r-l)的，最大高度被限制为height[l]。
//   而height[l] < height[r]，因此遍历过程中肯定遍历到了最大值

// @lc code=start
func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for l < r {
		h := min(height[l], height[r])
		area := h * (r - l)
		if area > maxArea {
			maxArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
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

// @lc code=end
