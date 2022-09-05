package q644_MaximumAverageSubarrayII

func findMaxAverage(nums []int, k int) float64 {
	preSum := 0
	for i := 0; i < k; i++ {
		preSum += nums[i]
	}

	maxAvg := float64(preSum) / float64(k)

	for i := k; i < len(nums); i++ {
		preSum += nums[i]

		sum := preSum
		avg := float64(sum) / float64(i+1)
		if avg > maxAvg {
			maxAvg = avg
		}

		for j := 0; j <= i-k; j++ {
			sum -= nums[j]

			avg := float64(sum) / float64(i-j)
			if avg > maxAvg {
				maxAvg = avg
			}
		}
	}

	return maxAvg
}
/*
n[1, 12, -5, -6, 50, 3] k=4

i-k
 |               i
[1, 12, -5, -6, 50, 3]
 j
*/
