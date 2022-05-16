package q159_LongestSubstringWithAtMost2DistinctCharacters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := make(map[byte]int)
	start := 0
	longest := 0

	for i := range s {
		m[s[i]]++

		for len(m)>2 {
			if m[s[start]]--; m[s[start]]==0 {
				delete(m, s[start])
			}

			start++
		}

		longest = max(longest, i-start+1)
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
