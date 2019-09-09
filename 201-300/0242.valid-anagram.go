/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */
func isAnagram(s string, t string) bool {
	hash := make(map[rune]int)
	for _, c := range s {
		hash[c]++
	}
	for _, c := range t {
		hash[c]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
