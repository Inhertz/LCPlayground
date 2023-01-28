package inordertraversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	stack := list.New()
	result := []int{}
	node := root

	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushFront(node)
			node = node.Left
		}

		node = stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		result = append(result, node.Val)
		node = node.Right
	}

	return result
}
