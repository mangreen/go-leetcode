package q325_MaximumSizeSubarraySumEqualsk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_output1(t *testing.T) {
	assert := assert.New(t)

	ans := maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3)

	assert.Equal(4, ans, "[1, -1, 5, -2] length is 4")
}

func Test_output2(t *testing.T) {
	assert := assert.New(t)

	ans := maxSubArrayLen([]int{-2, -1, 2, 1}, 1)

	assert.Equal(2, ans, "[-1, 2] length is 2")
}

func Test_output3(t *testing.T) {
	assert := assert.New(t)

	ans := maxSubArrayLen([]int{1, 0, -1}, -1)

	assert.Equal(2, ans, "[0, -1] length is 2")
}

func Test_output4(t *testing.T) {
	assert := assert.New(t)

	ans := maxSubArrayLen([]int{1, 1, 3}, 2)

	assert.Equal(2, ans, "[1, 1] length is 2")
}