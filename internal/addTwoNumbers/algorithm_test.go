package addtwonumbers

import (
	"reflect"
	"testing"
)

type TestCase struct {
	l1   *ListNode
	l2   *ListNode
	sum  *ListNode
	name string
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := []TestCase{
		{&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}, "Case 1"},
		{&ListNode{5, nil}, &ListNode{5, nil}, &ListNode{0, &ListNode{1, nil}}, "Case 2"},
		{&ListNode{9, &ListNode{9, nil}}, &ListNode{1, nil}, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}, "Case 3"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := addTwoNumbers(tc.l1, tc.l2)
			if !reflect.DeepEqual(sum, tc.sum) {
				t.Errorf("Expected sum %v but got %v", tc.sum, sum)
			}
		})
	}
}
