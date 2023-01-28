package maxwidthbintree

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeWithPos struct {
	node *TreeNode
	pos  int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(&TreeWithPos{node: root, pos: 0})
	res := 1

	for queue.Len() > 0 {
		minPos, maxPos := math.MaxInt, math.MinInt
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			element := queue.Front()
			queue.Remove(element)
			curr := element.Value.(*TreeWithPos)
			minPos, maxPos = min(minPos, curr.pos), max(maxPos, curr.pos)
			if curr.node.Left != nil {
				queue.PushBack(&TreeWithPos{node: curr.node.Left, pos: curr.pos * 2})
			}
			if curr.node.Right != nil {
				queue.PushBack(&TreeWithPos{node: curr.node.Right, pos: curr.pos*2 + 1})
			}
		}
		res = max(res, maxPos-minPos+1)
	}
	return res
}
