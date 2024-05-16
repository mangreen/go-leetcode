package q366_FindLeavesOfBinaryTree

func findLeaves(root *TreeNode) [][]int{
	ans := [][]int{}
	DFS(root, &ans)
	return ans
}

func DFS(n *TreeNode, ans *[][]int) int {
	if n == nil {
		return 0
	}

	depthL := DFS(n.Left, ans)
	depthR := DFS(n.Right, ans)
	depth := 1 + max(depthL, depthR)

	if len(*ans) < depth {
        *ans = append(*ans, []int{})
	}
    
    (*ans)[depth-1] = append((*ans)[depth-1], n.Val)

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}