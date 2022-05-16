package q114_ext_FlattenBinaryTreeToLinkedListInorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var n6 = &TreeNode{Val: 6}
var n5 = &TreeNode{Val: 5, Right: n6}
var n4 = &TreeNode{Val: 4}
var n3 = &TreeNode{Val: 3}
var n2 = &TreeNode{Val: 2, Left: n3, Right: n4}
var n1 = &TreeNode{Val: 1, Left: n2, Right: n5}

func Test_inorder(t *testing.T) {
	assert := assert.New(t)

	cur := inorder(n1)

	/*
			   1
			 /   \
			2	  5
		   / \     \
		  3	  4		6

		3 > 2 > 4 > 1 > 5 > 6
	*/
	assert.Equal(3, cur.Val, "should be 3")
	cur = cur.Right
	assert.Equal(2, cur.Val, "should be 2")
	cur = cur.Right
	assert.Equal(4, cur.Val, "should be 4")
	cur = cur.Right
	assert.Equal(1, cur.Val, "should be 1")
	cur = cur.Right
	assert.Equal(5, cur.Val, "should be 5")
	cur = cur.Right
	assert.Equal(6, cur.Val, "should be 6")
}
