package q323_NumberOfConnectedComponentsInAnUndirectedGraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_map_output2(t *testing.T) {
	assert := assert.New(t)

	ans := countComponents3(5, [][]int{{0, 1}, {1, 2}, {3, 4}})

	assert.Equal(2, ans, "Output is 2")
}

func Test_map_output1(t *testing.T) {
	assert := assert.New(t)

	ans := countComponents3(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}})

	assert.Equal(1, ans, "Output is 1")
}
