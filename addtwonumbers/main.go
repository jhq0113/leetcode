package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例 1：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：
1.每个链表中的节点数在范围 [1, 100] 内
2.0 <= Node.val <= 9
3.题目数据保证列表表示的数字不含前导零
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		next       uint8
		resultNode *ListNode
		result     = &ListNode{}
	)

	resultNode = result
	for {
		resultNode.Val = int(next)

		if l1 != nil {
			resultNode.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			resultNode.Val += l2.Val
			l2 = l2.Next
		}

		resultNode.Val, next = resultNode.Val%10, uint8(resultNode.Val/10)

		if l1 == nil && l2 == nil {
			if next > 0 {
				resultNode.Next = &ListNode{
					Val: int(next),
				}
			}
			break
		}

		//提前结束循环
		if l1 == nil {
			if next < 1 {
				resultNode.Next = l2
				break
			}
		}
		if l2 == nil {
			if next < 1 {
				resultNode.Next = l1
				break
			}
		}

		resultNode.Next = &ListNode{}
		resultNode = resultNode.Next
	}

	return result
}

func printList(l *ListNode) {
	res := ""
	for l != nil {
		res += strconv.Itoa(l.Val)
		l = l.Next
	}
	fmt.Println(res)
}

func main() {
	l1 := &ListNode{
		Val: rand.Intn(10),
	}

	l2 := &ListNode{
		Val: rand.Intn(10),
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(99)

	node1 := l1
	node2 := l2

	for index := 0; index < num; index++ {
		nod := &ListNode{
			Val: 1 + rand.Intn(9),
		}
		node1.Next = nod
		node1 = nod

		node := &ListNode{
			Val: 1 + rand.Intn(9),
		}
		node2.Next = node
		node2 = node
	}

	fmt.Print("l1:")
	printList(l1)
	fmt.Print("l2:")
	printList(l2)
	result := addTwoNumbers(l1, l2)
	fmt.Print("re:")
	printList(result)
}
