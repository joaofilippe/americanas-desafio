package main

import "github.com/joaofilippe/americanas-desafio/models"

// MergeListNode merge two sorted linked lists and return it as a sorted list.
func MergeListNode(list1, list2 *models.ListNode) *models.ListNode {
	list := new(models.ListNode)

	temp := list

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}

	if list1 != nil {
		temp.Next = list1
	} else {
		temp.Next = list2
	}

	return list.Next
}

// FromArrayToListNode convert an array to a ListNode
func FromArrayToListNode(array []int) *models.ListNode {
	list := new(models.ListNode)
	temp := list

	for _, value := range array {
		temp.Next = &models.ListNode{Val: value}
		temp = temp.Next
	}

	return list.Next
}

// FromListNodeToArray convert a ListNode to an array
func FromListNodeToArray(list *models.ListNode) []int {
	array := make([]int, 0)

	for list != nil {
		array = append(array, list.Val)
		list = list.Next
	}

	return array
}