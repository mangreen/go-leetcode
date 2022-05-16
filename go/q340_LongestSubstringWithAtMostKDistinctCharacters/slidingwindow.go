package q340_LongestSubstringWithAtMostKDistinctCharacters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := make(map[byte]int)
	start := 0
	maxL := 0

	for i := range s {
		m[s[i]]++

		for len(m) > k {
			m[s[start]]--

			if m[s[start]] == 0 {
				delete(m, s[i])
			}

			start++
		}

		maxL = max(maxL, i-start+1)
	}

	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}