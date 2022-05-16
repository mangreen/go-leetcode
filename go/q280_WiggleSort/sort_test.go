package q280_WiggleSort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sort_true(t *testing.T) {
	assert := assert.New(t)

	ans := wiggleSort2([]int{3, 5, 2, 1, 6, 4})

	assert.Equal(1, ans[0], "Output is 1")
	assert.Equal(3, ans[1], "Output is 3")
	assert.Equal(2, ans[2], "Output is 2")
	assert.Equal(5, ans[3], "Output is 5")
	assert.Equal(4, ans[4], "Output is 4")
	assert.Equal(6, ans[5], "Output is 6")
}

func Test_sort_sort(t *testing.T) {
	assert := assert.New(t)

	ans := wiggleSort([]int{1, 2, 3, 4, 5})

	for _, n := range ans {
		fmt.Print(n, " ")
	}

	assert.Equal(true, ans[0] <= ans[1], "Output is true")
	assert.Equal(true, ans[1] >= ans[2], "Output is true")
	assert.Equal(true, ans[2] <= ans[3], "Output is true")
	assert.Equal(true, ans[3] >= ans[4], "Output is true")
}
