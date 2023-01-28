package avglevelsbintree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	result := []float64{}
	if root == nil {
		return result
	}

	queue := list.New()
	sum := 0
	sumSize := 0
	var nodeAux *TreeNode

	queue.PushBack(root)

	for queue.Len() > 0 {
		sum = 0
		sumSize = queue.Len()
		for i := 0; i < sumSize; i++ {
			nodeAux = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			sum += nodeAux.Val

			if nodeAux.Left != nil {
				queue.PushBack(nodeAux.Left)
			}
			if nodeAux.Right != nil {
				queue.PushBack(nodeAux.Right)
			}
		}
		result = append(result, float64(sum)/float64(sumSize))
	}

	return result
}
