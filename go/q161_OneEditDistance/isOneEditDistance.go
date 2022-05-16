package q161_OneEditDistance

func isOneEditDistance(s string, t string) bool {
	for i:=0; i<min(len(s), len(t)); i++ {
		if s[i] != t [i] {
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			} else if len(s) < len(t) {
				return s[i:] == t[i+1:]
			} else {
				return s[i+1:] == t[i:]
			}
		}
	}

	return abs(len(s)-len(t)) == 1
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}