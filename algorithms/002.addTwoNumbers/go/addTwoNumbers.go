package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
        Val  int
        Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
        p1 := l1
        p2 := l2
        carry := 0
        header := &ListNode{}
        p := header
        for {
                if p1 == nil && p2 == nil {
                        if 0 == carry {
                                break
                        } else {
                                p.Next = &ListNode{Val: carry}
                                carry = 0
                        }
                } else if p1 == nil && p2 != nil {
                        v := (carry + p2.Val) % 10
                        carry = (carry + p2.Val) / 10
                        p.Next = &ListNode{Val: v}
                        p = p.Next
                        p2 = p2.Next
                } else if p1 != nil && p2 == nil {
                        v := (carry + p1.Val) % 10
                        carry = (carry + p1.Val) / 10
                        p.Next = &ListNode{Val: v}
                        p = p.Next
                        p1 = p1.Next
                } else {
                        v := (carry + p1.Val + p2.Val) % 10
                        carry = (carry + p1.Val + p2.Val) / 10
                        p.Next = &ListNode{Val: v}
                        p = p.Next
                        p1 = p1.Next
                        p2 = p2.Next
                }
        }
        return header.Next
}

func main() {
        l1 := &ListNode{Val: 2}
        l1.Next = &ListNode{Val: 4}
        l1.Next.Next = &ListNode{Val: 3}

        l2 := &ListNode{Val: 5}
        l2.Next = &ListNode{Val: 6}
        l2.Next.Next = &ListNode{Val: 4}

        l := addTwoNumbers(l1, l2)
        for {
                if l != nil {
                        fmt.Println(l.Val)
                        l = l.Next
                } else {
                        break
                }
        }
}
