package bst

import "testing"

func TestIsValidBST(t *testing.T) {
	var tree0 *TreeNode
	if isValidBST(tree0) != true {
		t.Errorf("#0: expected true, got false")
	}

	tree1 := &TreeNode{
		Val: 1,
	}
	if isValidBST(tree1) != true {
		t.Errorf("#1: expected true, got false")
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	if isValidBST(tree2) != true {
		t.Errorf("#2: expected true, got false")
	}

	tree3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	if isValidBST(tree3) != false {
		t.Errorf("#3: expected false, got true")
	}

	tree4 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	if isValidBST(tree4) != false {
		t.Errorf("#4: expected false, got true")
	}

	tree5 := &TreeNode{
		Val: 34,
		Left: &TreeNode{
			Val: -6,
			Left: &TreeNode{
				Val: -21,
			},
		},
	}
	if isValidBST(tree5) != true {
		t.Errorf("#5: expected true, got false")
	}
}
