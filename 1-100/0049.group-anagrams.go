/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, s := range strs {
		k := getKey(s)
		hash[*k] = append(hash[*k], s)
	}
	ans := make([][]string, 0, len(hash))
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

func getKey(str string) *[26]int {
	var k [26]int
	for _, c := range str {
		k[c-'a']++
	}
	return &k
}

