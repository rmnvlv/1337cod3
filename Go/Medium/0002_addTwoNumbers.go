package Medium

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Firt try Runtime: 4 - 12 ms, Memory Usage: 4.6 MB
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	myNode := ListNode{}
	myValue := &myNode
	sum, rest := 0, 0
	for l1 != nil || l2 != nil || rest != 0 {
		sum = rest
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		myValue.Next = &ListNode{Val: sum % 10, Next: nil}
		myValue = myValue.Next
		rest = sum / 10
	}
	return myNode.Next
}

// Second try Runtime: 4 - 33 ms, Memory Usage: 4.5 MB
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	myNode := ListNode{}
	myValue := &myNode
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		myValue.Next = &ListNode{Val: sum % 10, Next: nil}
		myValue = myValue.Next
		sum = sum / 10
	}
	return myNode.Next
}
