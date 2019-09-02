package main

import "fmt"

func main() {
	var l11 ListNode = ListNode{
		Val:  9,
		Next: nil,
	}
	var l12 ListNode = ListNode{
		Val:  8,
		Next: nil,
	}
	l11.Next = &l12
	var l21 ListNode = ListNode{
		Val:  1,
		Next: nil,
	}
	res := addTwoNumbers(&l11, &l21)
	fmt.Println(res.Val)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode = ListNode{
		Val:  0,
		Next: nil,
	}
	cur := &head
	flag := 0
	for ;l1!=nil && l2!=nil;{
		var next ListNode
		next.Val = (l1.Val + l2.Val + flag) % 10
		flag = (l1.Val + l2.Val + flag) / 10
		cur.Next = &next
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1==nil{
		for ;l2!=nil;{
			t := (l2.Val + flag) / 10
			l2.Val = (l2.Val + flag) % 10
			flag = t
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l2==nil{
		for ;l1!=nil;{
			t := (l1.Val + flag) / 10
			l1.Val = (l1.Val + flag) % 10
			flag = t
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	if flag==1{
		var tail ListNode = ListNode{
			Val:  1,
			Next: nil,
		}
		cur.Next = &tail
	}
	return head.Next
}

