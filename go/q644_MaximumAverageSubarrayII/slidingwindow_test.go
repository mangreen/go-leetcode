package q644_MaximumAverageSubarrayII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slidewindow_output(t *testing.T) {
	assert := assert.New(t)

	input := []int{1,12,-5,-6,50,3}
	ans := findMaxAverage(input, 4)

	assert.Equal(12.75, ans, "[12,-5,-6,50] has Max average 12.75")
}
