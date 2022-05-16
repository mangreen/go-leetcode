package q159_LongestSubstringWithAtMost2DistinctCharacters

func lengthOfLongestSubstringTwoDistinct2(s string) int {
	longest := 0

	start := 0
	end := -1

	for i:=1; i<len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}

		if end >= 0 && s[end] != s[i] {
			longest = max2(longest, i-start)
			start = end + 1
		}

		end = i - 1
	}

	return max2(len(s)-start, longest)
}

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}
