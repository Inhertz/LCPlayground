package mindepth

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	result := 0

	if root == nil {
		return result
	}

	size := 0
	var auxNode *TreeNode

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size = queue.Len()
		result++
		for i := 0; i < size; i++ {
			auxNode = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if auxNode.Left == nil && auxNode.Right == nil {
				return result
			}
			if auxNode.Left != nil {
				queue.PushBack(auxNode.Left)
			}
			if auxNode.Right != nil {
				queue.PushBack(auxNode.Right)
			}
		}
	}

	return result
}
