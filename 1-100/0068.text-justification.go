/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */
func fullJustify(words []string, maxWidth int) []string {
	text := make([]string, 0, 8)
	for len(words) > 0 {
		oneLineWords := pickOneLineWords(words, maxWidth)
		words = words[len(oneLineWords):]
		spaces := getSpaceCntAfterEachWord(oneLineWords, maxWidth, len(words) == 0)
		text = append(text, formOneLine(oneLineWords, maxWidth, spaces))
	}
	return text
}

func pickOneLineWords(words []string, maxWidth int) []string {
	cnt := 0
	maxWidth++
	for _, word := range words {
		if len(word)+1 <= maxWidth {
			cnt++
			maxWidth -= len(word) + 1
		} else {
			break
		}
	}
	return words[:cnt]
}

func formOneLine(words []string, maxWidth int, spaces []int) string {
	buf := getAllSpaceByteSlice(maxWidth)
	start := 0
	for i, word := range words {
		copy(buf[start:], word)
		start += len(word) + spaces[i]
	}
	return string(buf)
}

func getSpaceCntAfterEachWord(words []string, maxWidth int, isLastLine bool) []int {
	totalSpaceCnt := maxWidth
	for _, word := range words {
		totalSpaceCnt -= len(word)
	}
	n := len(words)
	if n == 1 {
		return []int{totalSpaceCnt}
	}
	spaces := make([]int, n)
	if isLastLine {
		for i := 0; i < n-1; i++ {
			spaces[i] = 1
		}
		spaces[n-1] = totalSpaceCnt - n + 1
	} else {
		averageSpaceCnt := totalSpaceCnt / (n - 1)
		extraSpaceCnt := totalSpaceCnt - averageSpaceCnt*(n-1)
		for i := 0; i < n-1; i++ {
			spaces[i] = averageSpaceCnt
			if i < extraSpaceCnt {
				spaces[i]++
			}
		}
	}
	return spaces
}

func getAllSpaceByteSlice(maxWidth int) []byte {
	buf := make([]byte, maxWidth)
	for i := range buf {
		buf[i] = ' '
	}
	return buf
}
