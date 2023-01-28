package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{Val: 0, Next: nil}
	looper := out
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = (sum - sum%10) / 10
		looper.Val = sum % 10

		if l1 != nil {
			if l1.Next != nil {
				looper.Next = &ListNode{Val: 0, Next: nil}
			}
		}

		if l2 != nil {
			if l2.Next != nil {
				looper.Next = &ListNode{Val: 0, Next: nil}
			}
		}

		if carry != 0 {
			looper.Next = &ListNode{Val: 0, Next: nil}
		}

		looper = looper.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return out
}
