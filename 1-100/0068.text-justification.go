/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */
package leetcode

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	var text []string
	for len(words) > 0 {
		picked := pickOneLineWords(words, maxWidth)
		words = words[len(picked):]
		spaceCnt := getSpaceCntAfterEachWord(picked, maxWidth, len(words) == 0)
		sentence := buildSentence(picked, spaceCnt, maxWidth)
		text = append(text, sentence)
	}
	return text
}

func pickOneLineWords(words []string, maxWidth int) []string {
	remainWidth := maxWidth
	pos := 0
	for pos < len(words) && len(words[pos]) <= remainWidth {
		remainWidth -= len(words[pos])
		pos++
		remainWidth--
	}
	return words[:pos]
}

func getSpaceCntAfterEachWord(words []string, maxWidth int, isLastLine bool) []int {
	remainWidth := maxWidth
	for _, word := range words {
		remainWidth -= len(word)
	}

	n := len(words)
	var spaceCnt []int
	if n == 1 || isLastLine {
		spaceCnt = make([]int, n)
		for i := range spaceCnt {
			if i < n-1 {
				spaceCnt[i] = 1
				remainWidth--
			} else {
				spaceCnt[i] = remainWidth
			}
		}
	} else {
		spaceCnt = make([]int, n-1)
		avg := remainWidth / (n - 1)
		extra := remainWidth % (n - 1)
		for i := range spaceCnt {
			spaceCnt[i] = avg
			if i < extra {
				spaceCnt[i]++
			}
		}
	}

	return spaceCnt
}

func buildSentence(words []string, spaceCnt []int, maxWidth int) string {
	buf := make([]byte, 0, maxWidth)
	for i, word := range words {
		buf = append(buf, word...)
		if i < len(spaceCnt) {
			buf = append(buf, getSpaces(spaceCnt[i])...)
		}
	}
	return string(buf)
}

func getSpaces(n int) []byte {
	spaces := make([]byte, n)
	for i := range spaces {
		spaces[i] = ' '
	}
	return spaces
}

// @lc code=end
