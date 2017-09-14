package Problem0114

import (
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	cur(root)
}

// cur 会 flatten root，并返回 flatten 后的 leaf 节点
func cur(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return cur(root.Right)
	}

	res := cur(root.Right)
	cur(root.Left).Right = root.Right
	root.Right = root.Left
	// 不要忘记封闭 left 节点
	root.Left = nil

	return res
}
