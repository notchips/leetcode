/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
package leetcode

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string)
	var key [26]int
	for _, str := range strs {
		// reset key
		for i := range key {
			key[i] = 0
		}

		// set key
		for i := 0; i < len(str); i++ {
			key[int(str[i]-'a')]++
		}

		// add to map
		strMap[key] = append(strMap[key], str)
	}

	ans := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
