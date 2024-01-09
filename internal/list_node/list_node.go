package listnode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaofilippe/americanas-desafio/internal/models"
)


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

// FromSliceToListNode convert an slice to a ListNode
func FromSliceToListNode(array []int) *models.ListNode {
	list := new(models.ListNode)
	temp := list

	for _, value := range array {
		temp.Next = &models.ListNode{Val: value}
		temp = temp.Next
	}

	return list.Next
}

// FromListNodeToSlice convert a ListNode to an array
func FromListNodeToSlice(list *models.ListNode) []int {
	array := make([]int, 0)

	for list != nil {
		array = append(array, list.Val)
		list = list.Next
	}

	return array
}

// FromListNodeToString convert a ListNode to a string
func FromListNodeToString(list *models.ListNode) string {
	array := make([]int, 0)

	for list != nil {
		array = append(array, list.Val)
		list = list.Next
	}
	
	var stringArray string
	for i, v := range array {
		if i == len(array)-1 {
			stringArray += fmt.Sprintf("%d", v)
			break
		}

		stringArray += fmt.Sprintf("%d,", v)
	}

	return stringArray
}

// FromStringToListNode convert a string to a ListNode
func FromStringToListNode(stringArray string) *models.ListNode {
	array := make([]int, 0)

	stringArray = strings.Trim(stringArray, "[{}]")
	arrayString := strings.Split(stringArray, ",")

	for _, value := range arrayString {
		v, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		array = append(array, v)
	}

	list := new(models.ListNode)
	temp := list

	for _, value := range array {
		temp.Next = &models.ListNode{Val: value}
		temp = temp.Next
	}

	return list.Next
}

// ValidateSorted validate if a ListNode is sorted
func ValidateSorted(list *models.ListNode) bool {
	for list != nil && list.Next != nil {
		if list.Val > list.Next.Val {
			return false
		}
		list = list.Next
	}

	return true
}