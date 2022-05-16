package q291_WordPatternII

import "strings"

func wordPatternMatch(pattern string, str string) bool {
	m := make(map[byte]string)
	u := make(map[string]bool)

	return dfs(pattern, 0, str, 0, m, u)
}

func dfs(pat string, i int, str string, j int, m map[byte]string, u map[string]bool) bool {
	if i == len(pat) && j == len(str) {
		return true
	}

	if i >= len(pat) || j >= len(str) {
		return false
	}

	b := pat[i]
	if w, ok := m[b]; ok {
		if !strings.HasPrefix(str[j:], w) {
			return false
		}
	}

	for k:=j; k<len(str); k++ {
		sub := str[j:k+1]

		if u[sub] {
			continue
		}

		m[b] = sub
		u[sub] = true

		if dfs(pat, i+1, str, k+1, m, u) {
			return true
		}

		delete(m, b)
		delete(u, sub)
	}

	return false
}
