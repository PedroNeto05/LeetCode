package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3

	list1 := node1

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}

	node4.Next = node5
	node5.Next = node6

	list2 := node4

	mergeTwoLists(list1, list2)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	cur := newHead
	for list1 != nil || list2 != nil {
		if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
			cur.Next = list1
			list1 = list1.Next
		} else if list2 != nil {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	return newHead.Next
}
