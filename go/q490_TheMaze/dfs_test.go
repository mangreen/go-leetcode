package q490_TheMaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maze2 = [][]int{
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{1, 1, 0, 1, 1},
	{0, 0, 0, 0, 0},
}

func Test_hasPath2(t *testing.T) {
	assert := assert.New(t)

	ans := hasPath2(maze2, []int{0, 4}, []int{4, 4})

	assert.Equal(true, ans, "they should be equal")
}

func Test_hasPath2_failed(t *testing.T) {
	assert := assert.New(t)

	ans := hasPath2(maze2, []int{0, 4}, []int{3, 2})

	assert.Equal(false, ans, "they should be equal")
}
