package problem124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a int, ints ...int) int {
	res := a
	for _, b := range ints {
		if b > res {
			res = b
		}
	}
	return res
}

func maxPathSumInternal(node *TreeNode) (int, int, int) {
	leftPathSum := node.Val
	maxLeft := node.Val
	if node.Left != nil {
		l, r, max := maxPathSumInternal(node.Left)
		leftPathSum = maxInt(node.Val, node.Val+l, node.Val+r)
		if max > maxLeft {
			maxLeft = max
		}
	}

	rightPathSum := node.Val
	maxRight := node.Val
	if node.Right != nil {
		l, r, max := maxPathSumInternal(node.Right)
		rightPathSum = maxInt(node.Val, node.Val+l, node.Val+r)
		if max > maxRight {
			maxRight = max
		}
	}

	maxPathSum := maxInt(leftPathSum+rightPathSum-node.Val, maxLeft, maxRight)

	return leftPathSum, rightPathSum, maxPathSum
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, _, res := maxPathSumInternal(root)
	return res
}
