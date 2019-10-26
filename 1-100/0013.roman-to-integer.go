/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package leetcode

// @lc code=start
func romanToInt(s string) int {
	hash := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += hash[s[i]]
		if i < len(s)-1 {
			switch s[i : i+2] {
			case "IV", "IX", "XL", "XC", "CD", "CM":
				sum -= 2 * hash[s[i]]
			}
		}
	}
	return sum
}

// @lc code=end
