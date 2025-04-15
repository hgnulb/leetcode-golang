package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 排序奇升偶降链表
func sortOddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 拆分奇偶链表
	odd, even := head, head.Next
	oddHead, evenHead := odd, even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next

		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = nil // 非常重要

	// 反转偶数链表
	var pre *ListNode
	cur := evenHead
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	evenHead = pre

	// 合并两个链表
	dummy := &ListNode{}
	cur = dummy
	l1, l2 := oddHead, evenHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

// 打印链表
func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " -> ")
		cur = cur.Next
	}
	fmt.Println("NULL")
}

func main() {
	// 示例链表
	head := &ListNode{1, &ListNode{8, &ListNode{3, &ListNode{6, &ListNode{5, &ListNode{4, &ListNode{7, &ListNode{2, nil}}}}}}}}
	fmt.Println("原链表:")
	printList(head)

	sortedHead := sortOddEvenList(head)
	fmt.Println("排序后的链表:")
	printList(sortedHead)
}
