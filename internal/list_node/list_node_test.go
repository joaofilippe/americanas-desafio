package listnode

import (
	"testing"

	"github.com/joaofilippe/americanas-desafio/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_FromArrayToListNode(t *testing.T) {
	assert := assert.New(t)

	slices := [][]int{
		{1, 2, 3},
		{5, 6, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1},
		{10, 12, 11},
	}

	for _, list := range slices {
		listNode := FromArrayToListNode(list)
		CheckListNodeFromSlice(assert, listNode, list)
	}
}

func Test_FromListNodeToArray(t *testing.T) {
	assert := assert.New(t)
	listNodes := []*models.ListNode{
		&models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 2,
				Next: &models.ListNode{
					Val: 3,
				},
			}},
		&models.ListNode{
			Val: 5,
			Next: &models.ListNode{
				Val: 6,
				Next: &models.ListNode{
					Val: 8,
				},
			}},
		&models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 2,
				Next: &models.ListNode{
					Val: 3,
					Next: &models.ListNode{
						Val: 4,
						Next: &models.ListNode{
							Val: 5,
							Next: &models.ListNode{
								Val: 6,
							},
						},
					},
				},
			}},
		&models.ListNode{
			Val: 1,
		},
		&models.ListNode{
			Val: 10,
			Next: &models.ListNode{
				Val: 12,
				Next: &models.ListNode{
					Val: 11,
				},
			}},
	}

	for _, list := range listNodes {
		slices := FromListNodeToArray(list)
		CheckSliceFromListNode(assert, slices, list)
	}
}

func Test_MergeList(t *testing.T) {
	tests := []struct {
		listA    []int
		listB    []int
		expected []int
	}{
		{
			listA:    []int{1, 2, 3},
			listB:    []int{5, 6, 8},
			expected: []int{1, 2, 3, 5, 6, 8},
		},
		{
			listA:    []int{1, 2, 8, 9},
			listB:    []int{5, 6, 10},
			expected: []int{1, 2, 5, 6, 8, 9, 10},
		},
		{
			listA:    []int{11, 40, 100},
			listB:    []int{5, 6, 10},
			expected: []int{5, 6, 10, 11, 40, 100},
		},
	}

	for _, test := range tests {
		listNodeA := FromArrayToListNode(test.listA)
		listNodeB := FromArrayToListNode(test.listB)

		merged := MergeListNode(listNodeA, listNodeB)
		slice := FromListNodeToArray(merged)

		for i, v := range slice {
			if v != test.expected[i] {
				t.Errorf("Expected %d, got %d", test.expected[i], v)
			}
		}
	}

}

func Test_FromListNodeToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		slice    []int
		expected string
	}{
		{slice: []int{1, 2, 3}, expected: "1,2,3"},
		{slice: []int{5, 6, 8}, expected: "5,6,8"},
		{slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: "1,2,3,4,5,6,7,8,9"},
		{slice: []int{1}, expected: "1"},
		{slice: []int{10, 12, 11}, expected: "10,12,11"},
	}

	for _, test := range tests {
		listNode := FromArrayToListNode(test.slice)
		stringList := FromListNodeToString(listNode)

		assert.Equal(test.expected, stringList, "Expected %s, got %s", test.expected, stringList)
	}
}

func Test_FromStringToListNode(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		str      string
		expected []int
	}{
		{str: "{1,2,3}", expected: []int{1, 2, 3}},
		{str: "{5,6,8}", expected: []int{5, 6, 8}},
		{str: "{1,2,5,7,8,9}", expected: []int{1, 2, 5, 7, 8, 9}},
		{str: "{1}", expected: []int{1}},
		{str: "{10,12,11}", expected: []int{10, 12, 11}},
	}

	for _, test := range tests {
		listNode := FromStringToListNode(test.str)
		slice := FromListNodeToArray(listNode)

		for i, v := range slice {
			assert.Equal(test.expected[i], v, "Expected %d, got %d", test.expected[i], v)
		}
	}
}

func Test_ValidateSorted(t *testing.T) {
	a := []int{1, 2, 3}

	list := FromArrayToListNode(a)

	valid := ValidateSorted(list)

	if !valid {
		t.Errorf("Expected true, got %t", valid)
	}
}

func CheckListNodeFromSlice(assert *assert.Assertions, list *models.ListNode, expected []int) {
	for _, v := range expected {
		assert.Equal(v, list.Val, "Expected %d, got %d", v, list.Val)

		if list.Next != nil {
			list = list.Next
		} else {
			break
		}
	}

}

func CheckSliceFromListNode(assert *assert.Assertions, slice []int, expected *models.ListNode) {
	for _, v := range slice {
		assert.Equal(expected.Val, v, "Expected %d, got %d", expected.Val, v)

		if expected.Next != nil {
			expected = expected.Next
		} else {
			break
		}
	}

}
