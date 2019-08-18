package problem114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToList(root *TreeNode) *TreeNode {
	last := root
	if root.Left != nil {
		last = treeToList(root.Left)
		last.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if last.Right != nil {
		last = treeToList(last.Right)
	}
	return last
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	treeToList(root)
}
