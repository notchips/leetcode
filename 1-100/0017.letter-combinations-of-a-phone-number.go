/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	hash := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	ans := make([]string, 0, 20)
	buf := make([]byte, 0, 4*len(digits))
	dfs(hash, digits, 0, &ans, &buf)
	return ans
}

func dfs(hash map[byte][]byte, digits string, i int, ans *[]string, buf *[]byte) {
	if i == len(digits) {
		newBuf := make([]byte, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, string(newBuf))
		return
	}
	chars := hash[digits[i]]
	for _, char := range chars {
		*buf = append(*buf, char)
		dfs(hash, digits, i+1, ans, buf)
		*buf = (*buf)[:len(*buf)-1]
	}
}
