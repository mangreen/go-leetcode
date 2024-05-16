package q366_FindLeavesOfBinaryTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLeaves(t *testing.T) {
	assert := assert.New(t)

	n5 := &TreeNode{5, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{2, n4, n5}
	n1 := &TreeNode{1, n2, n3}

	ans := findLeaves(n1)

	for _, row := range ans {
		for _, v := range row {
			fmt.Println("v: ", v)
		}
	}

	mockAns := [][]int{{4, 5, 3}, {2}, {1}}

	assert.Equal(mockAns, ans, "Should be equal to [[4, 5, 3], [2], [1]]")
}
