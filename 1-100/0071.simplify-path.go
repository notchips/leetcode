/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */
func simplifyPath(path string) string {
	nodes := strings.Split(path, "/")
	stack := make([]string, 0, 10)
	for _, node := range nodes {
		if node == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if node != "." && node != ".." && len(node) > 0 {
			stack = append(stack, node)
		}
	}
	return "/" + strings.Join(stack, "/")
}
