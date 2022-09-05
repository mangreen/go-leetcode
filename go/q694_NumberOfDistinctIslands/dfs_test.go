package q694_NumberOfDistinctIslands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_island_1(t *testing.T) {
	assert := assert.New(t)

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	ans := numDistinctIslands(grid)

	assert.Equal(1, ans, " only 1 island")
}

func Test_islands_3(t *testing.T) {
	assert := assert.New(t)

	grid := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}
	ans := numDistinctIslands(grid)

	assert.Equal(3, ans, "3 islands")
}