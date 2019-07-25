/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */
func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0, 10)
	for len(words) > 0 {
		pw := pickWords(words, maxWidth)
		words = words[len(pw):]
		var sentence string
		if len(words) == 0 {
			sentence = getLastSentence(pw, maxWidth)
		} else {
			spaceCnt := getSpaceCntAfterEachWord(pw, maxWidth)
			sentence = getSentence(pw, spaceCnt, maxWidth)
		}
		ans = append(ans, sentence)
	}
	return ans
}

func pickWords(words []string, width int) []string {
	end := 1
	width -= len(words[0])
	for i := 1; i < len(words); i++ {
		if width >= 1+len(words[i]) {
			end++
			width -= 1 + len(words[i])
		} else {
			break
		}
	}
	return words[:end]
}

func getSpaceCntAfterEachWord(words []string, width int) []int {
	for _, word := range words {
		width -= len(word)
	}
	if len(words) == 1 {
		return []int{width}
	}
	IntervalCnt := len(words) - 1
	ret := make([]int, IntervalCnt)
	for i := range ret {
		ret[i] = width / IntervalCnt
	}
	width %= IntervalCnt
	for i := 0; i < width; i++ {
		ret[i]++
	}
	return append(ret, 0)
}

func getSentence(words []string, spaceCnt []int, width int) string {
	buf := make([]byte, 0, width)
	for i := 0; i < len(words); i++ {
		buf = append(buf, []byte(words[i])...)
		buf = appendSpace(buf, spaceCnt[i])
	}
	return string(buf)
}

func getLastSentence(words []string, width int) string {
	buf := make([]byte, 0, width)
	for i := 0; i < len(words); i++ {
		buf = append(buf, []byte(words[i])...)
		width -= len(words[i])
		if i != len(words)-1 {
			buf = append(buf, ' ')
			width--
		}
	}
	buf = appendSpace(buf, width)
	return string(buf)
}

func appendSpace(buf []byte, n int) []byte {
	for i := 0; i < n; i++ {
		buf = append(buf, ' ')
	}
	return buf
}
