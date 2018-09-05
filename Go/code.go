 
package main

import "fmt"

//公用:2. 21.
type ListNode struct {
	Val int
	Next *ListNode
}

func main (){
	
	//1.两数之和 twoSum
	//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。 
	//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。 
	//示例: 
	//给定 nums = [2, 7, 11, 15], target = 9 
	//因为 nums[0] + nums[1] = 2 + 7 = 9
	//所以返回 [0, 1]
	//nums := []int{2, 7, 11, 15, 2}
	//result := twoSum(nums, 4)
	//fmt.Println(result)
	
	
	//2. 两数相加 addTwoNumbers
	//给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，
	//它们的每个节点只存储单个数字。将两数相加返回一个新的链表。 
	//你可以假设除了数字 0 之外，这两个数字都不会以零开头。 
	//示例： 
	//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	//输出：7 -> 0 -> 8
	//原因：342 + 465 = 807
	//l1 := &ListNode{Val : 2, Next : &ListNode{Val : 4, Next : &ListNode{Val : 3, Next: nil}}}
	//l2 := &ListNode{Val : 5, Next : &ListNode{Val : 6, Next : &ListNode{Val : 4, Next: &ListNode{Val : 3, Next: nil}}}}
	////l1 := &ListNode{Val : 5, Next :  nil}
	////l2 := &ListNode{Val : 5, Next :  nil}
	//l3 := addTwoNumbers(l1, l2) 
	//for l3 != nil{
	//	fmt.Print(l3.Val)
	//	l3 = l3.Next
	//}

	//21. 合同两个有序链表 mergeTwoLists
	//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
	//示例： 
	//输入：1->2->4, 1->3->4
	//输出：1->1->2->3->4->4
	//l1 := &ListNode{Val : 1, Next : &ListNode{Val : 2, Next : &ListNode{Val : 4, Next: nil}}}
	//l2 := &ListNode{Val : 1, Next : &ListNode{Val : 3, Next : &ListNode{Val : 4, Next: &ListNode{Val : 5, Next: nil}}}}
	var l1 *ListNode
	var l2 *ListNode
	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3)
	for l3 != nil{
		fmt.Print(l3.Val)
		l3 = l3.Next
	}

}

func twoSum(nums []int, target int) []int {
	var result []int
	maps := make(map[int]int)
	
	for k, v := range nums{
		if _, ok := maps[target - v]; ok {
			result = append(result, maps[target - v])
			result = append(result, k)
			return result
		}
		maps[v] = k
	} 
	return result
	
}


/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/ 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	tempP := l3
	tempN := 0 
	temp := 0
	for { 
		if l1 != nil{
			temp += l1.Val
			l1 = l1.Next
		} 
		if l2 != nil{
			temp += l2.Val
			l2 = l2.Next
		}
		tempP.Val = (temp + tempN) % 10
		tempN = (temp + tempN) / 10
		temp = 0
		if l1 == nil && l2 == nil && tempN == 0{
			break
		}
		tempP.Next = &ListNode{}
		tempP = tempP.Next
	}
	return l3
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	tempP := l3
	for l1 != nil || l2 != nil{
		if l1 != nil && l2 != nil{
			if l1.Val <= l2.Val {
				tempP.Next = &ListNode{Val : l1.Val}
				l1 =l1.Next
			}else{
				tempP.Next = &ListNode{Val : l2.Val}
				l2 =l2.Next
			}
		}else{
			if l1 == nil && l2 != nil {
				tempP.Next = &ListNode{Val : l2.Val}
				l2 = l2.Next
			}else{
				tempP.Next = &ListNode{Val : l1.Val}
				l1 =l1.Next
			}
		}
		//if l1 == nil && l2 == nil{
		//	break
		//}
		tempP = tempP.Next
	}
	return l3.Next
}

