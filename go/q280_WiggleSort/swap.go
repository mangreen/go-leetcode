package q280_WiggleSort

func wiggleSort(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}

	for i := 1; i < len(nums); i++ {
		if (i%2 == 1 && nums[i-1] > nums[i]) || (i%2 == 0 && nums[i-1] < nums[i]) {
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
	}

	return nums
}
