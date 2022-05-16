package q114_ext_FlattenBinaryTreeToLinkedListInorder

func inorder(root *TreeNode) *TreeNode {
	return DFS(root)
}

func DFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := DFS(root.Left)
	DFS(root.Right)

	if l == nil {
		return root
	}

	c := l
	root.Left = nil

	for c != nil && c.Right != nil {
		c = c.Right
	}

	c.Right = root

	return l
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
