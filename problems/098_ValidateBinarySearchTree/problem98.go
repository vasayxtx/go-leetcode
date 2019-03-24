package problem98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func exploreBSTNode(node *TreeNode) (bool, int, int, int) {
	if node == nil {
		return true, 0, 0, 0
	}
	leftOk, leftLen, leftMin, leftMax := exploreBSTNode(node.Left)
	if !leftOk {
		return false, 0, 0, 0
	}
	rightOk, rightLen, rightMin, rightMax := exploreBSTNode(node.Right)
	if !rightOk {
		return false, 0, 0, 0
	}
	min := node.Val
	if leftLen > 0 {
		if node.Val <= leftMax {
			return false, 0, 0, 0
		}
		min = leftMin
	}
	max := node.Val
	if rightLen > 0 {
		if node.Val >= rightMin {
			return false, 0, 0, 0
		}
		max = rightMax
	}

	return true, maxInt(leftLen, rightLen) + 1, min, max

}

func isValidBST(node *TreeNode) bool {
	ok, _, _, _ := exploreBSTNode(node)
	return ok
}
