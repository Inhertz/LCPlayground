package levelorderbsf

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	node := root

	var smallSliceSize int
	var smallSlice []int

	queue.PushBack(node)

	for queue.Len() != 0 {
		smallSlice = []int{}
		smallSliceSize = queue.Len()
		for i := 0; i < smallSliceSize; i++ {
			node = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			smallSlice = append(smallSlice, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, smallSlice)
	}
	return result
}
