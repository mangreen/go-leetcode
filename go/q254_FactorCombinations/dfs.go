package q254_FactorCombinations

func getFactors(n int) [][]int {
	ans := [][]int{}

	dfs(n, 2, []int{}, &ans)

	return ans
}

func dfs(n int, idx int, set []int, ans *[][]int) {
	if n == 1 {
		// check len(factors) to make sure we don't include factors for numbers like 37 which factors is [37]
		// which we shouldn't include
		if len(set) > 1 {
			*ans = append(*ans, append([]int{}, set...))
		}

		return
	}

	// note: i<= n instead of i < n, the prime number such as 37 is handled in the base case above
	// since i starts with the start, the number in the factors will be in increasing order, so we don't have duplicates
	// like [2,2,3] and [3,2,2]
	for i:=idx; i<=n; i++ {
		// note: only make sense to add to factors when the remainder is 0
		if n%i == 0 {
			dfs(n/i, i, append(set, i), ans)
		}
	}
}