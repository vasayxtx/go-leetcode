package problem103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(list []int) {
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		l := len(queue)
		var line []int
		for i := 0; i < l; i++ {
			node := queue[i]
			line = append(line, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		if level%2 == 1 {
			reverse(line)
		}
		res = append(res, line)
	}
	return res
}
