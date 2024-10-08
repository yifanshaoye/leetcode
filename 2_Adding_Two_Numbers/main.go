package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	p, q := l1, l2
	res := &ListNode{}
	r := res

	sum, val, carry := 0, 0, 0
	for p != nil && q != nil {
		sum = p.Val + q.Val + carry
		val = sum % 10
		carry = sum / 10
		r.Next = &ListNode{
			Val: val,
		}
		r = r.Next
		p = p.Next
		q = q.Next
	}

	for p != nil {
		sum = p.Val + carry
		val = sum % 10
		carry = sum / 10
		r.Next = &ListNode{
			Val: val,
		}
		r = r.Next
		p = p.Next
	}

	for q != nil {
		sum = q.Val + carry
		val = sum % 10
		carry = sum / 10
		r.Next = &ListNode{
			Val: val,
		}
		r = r.Next
		q = q.Next
	}

	if carry > 0 {
		r.Next = &ListNode{
			Val: carry,
		}
	}

	return res.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	pre, p, q := l1, l1, l2

	sum, carry := 0, 0
	for p != nil && q != nil {
		sum = p.Val + q.Val + carry
		p.Val = sum % 10
		carry = sum / 10

		pre = p
		p = p.Next
		q = q.Next
	}

	if p == nil {
		pre.Next = q
		p = q
	}

	for p != nil {
		sum = p.Val + carry
		p.Val = sum % 10
		carry = sum / 10

		pre = p
		p = p.Next
	}

	if carry > 0 {
		pre.Next = &ListNode{
			Val: carry,
		}
	}

	return l1
}

func main() {
	l1 := transferNumberToList(98765)
	l2 := transferNumberToList(654)
	fmt.Println(transferListToNumber(addTwoNumbers1(l1, l2)))
	fmt.Println(transferListToNumber(addTwoNumbers2(l1, l2)))
}

func transferNumberToList(num int) *ListNode {
	head := &ListNode{}
	if num == 0 {
		return head
	}
	p := head
	for num > 0 {
		p.Next = &ListNode{
			Val: num % 10,
		}
		num /= 10
		p = p.Next
	}
	return head.Next
}

func transferListToNumber(head *ListNode) int {
	num, pow := 0, 1
	for head != nil {
		num += head.Val * pow
		head = head.Next
		pow *= 10
	}
	return num
}

func addTwoNumbersPattern2(nums1, nums2 []int) int {

	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	for i < m && j < n {
		// deal with common section

		i++
		j++
	}

	for i < m {
		// nums1 is longer, deal with the remaining

		i++
	}

	for j < n {
		// nums2 is longer, deal with the remaining

		j++
	}

	return 0

}

func addTwoNumbersPattern1(l1 *ListNode, l2 *ListNode) *ListNode {

	p, q := l1, l2
	for p != nil && q != nil {
		// deal with common section

		p = p.Next
		q = q.Next
	}

	for p != nil {
		// l1 is longer, deal with the remaining

		p = p.Next
	}

	for q != nil {
		// l2 is longer, deal with the remaining

		q = q.Next
	}

	return l1

}
