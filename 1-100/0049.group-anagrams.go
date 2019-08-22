/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
 type key [26]int

 func groupAnagrams(strs []string) [][]string {
	 m := make(map[key][]string)
	 for _, s := range strs {
		 k := getKey(s)
		 m[*k] = append(m[*k], s)
	 }
	 ans := make([][]string, 0, len(m))
	 for _, v := range m {
		 ans = append(ans, v)
	 }
	 return ans
 }
 
 func getKey(s string) *key {
	 var k key
	 for _, c := range s {
		 k[c-'a']++
	 }
	 return &k
 }
 