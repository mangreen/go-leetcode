package q505_theMazeII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maze = [][]int{
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{1, 1, 0, 1, 1},
	{0, 0, 0, 0, 0},
}

func Test_hasPath(t *testing.T) {
	assert := assert.New(t)

	ans := shortestDistance(maze, []int{0, 4}, []int{4, 4})

	assert.Equal(12, ans, "they should be equal")
}

func Test_hasPath_failed(t *testing.T) {
	assert := assert.New(t)

	ans := shortestDistance(maze, []int{0, 4}, []int{3, 2})

	assert.Equal(-1, ans, "they should be equal")
}