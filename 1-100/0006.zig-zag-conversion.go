/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
package leetcode

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	buf := make([]byte, 0, len(s))
	for row := 0; row < numRows; row++ {
		for pos := row; pos < len(s); pos += 2 * (numRows - 1) {
			buf = append(buf, s[pos])
			if 0 < row && row < numRows-1 {
				extra := pos + 2*(numRows-1-row)
				if extra < len(s) {
					buf = append(buf, s[extra])
				}
			}
		}
	}
	return string(buf)
}

// @lc code=end
