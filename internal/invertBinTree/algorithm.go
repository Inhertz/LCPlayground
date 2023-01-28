package invertbintree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	queue := list.New()
	var currentNode *TreeNode
	var bubbleNode *TreeNode

	queue.PushBack(root)

	for queue.Len() > 0 {

		levelSize := queue.Len()

		for i := 0; i < levelSize; i++ {
			currentNode = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			bubbleNode = currentNode.Left
			currentNode.Left = currentNode.Right
			currentNode.Right = bubbleNode
			if currentNode.Left != nil {
				queue.PushBack(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.PushBack(currentNode.Right)
			}
		}
	}

	return root
}
