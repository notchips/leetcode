/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	tempArr := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += (numRows - 1) * 2 {
			tempArr = append(tempArr, s[j])
			extra := j + (numRows-1-i)*2
			if i != 0 && i != numRows-1 && extra < len(s) {
				tempArr = append(tempArr, s[extra])
			}
		}
	}
	return string(tempArr)
}
