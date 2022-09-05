package q325_MaximumSizeSubarraySumEqualsk

func maxSubArrayLen(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = -1

	preSum := 0

	maxL := 0

	for i, v := range nums {
		preSum += v

		if idx, ok := m[preSum-k]; ok {
			maxL = max(maxL, i-idx)
		}

		// 因為求maxSubAry, 只記錄最初的餘數idx(最左邊界)
		if _, ok := m[preSum]; !ok {
			m[preSum] = i
		}
	}

	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
/*
ex1.
       0   1  2   3  4
nums: [1, -1, 5, -2, 3], k = 3
sums: [1,  0, 5,  3, 6]
       <- sum=k  ->

ans = i+1 = 3+1 = 4


ex2.
        0   1   2  3
nums: [-2, -1,  2, 1], k = 1
sums: [-2, -3, -1, 0]
            <-k->
        <- sum -> 

ans = max(ans=0, i-m[sum-k]) = 2 - (m[-1-1]) = 2-0 = 2


ex3.
       0  1   2
nums: [1, 0, -1], k = -1
sums: [1, 1,  0]

ans = max(i-m[sum-k]) = 2 - (m[0+1]) = 2-0 = 2
*/
