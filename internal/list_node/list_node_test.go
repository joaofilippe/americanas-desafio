package listnode

import (
	"testing"

	"github.com/joaofilippe/americanas-desafio/internal/models"
)

func Test_FromArrayToListNode(t *testing.T) {
	a := []int{1, 2, 3}

	b := []int{0, 6, 8}
	
	list := FromArrayToListNode(a)
	list2 := FromArrayToListNode(b)

	if list.Val != 1 {
		t.Errorf("Expected 1, got %d", list.Val)
	}

	if list.Next.Val != 2 {
		t.Errorf("Expected 2, got %d", list.Next.Val)
	}

	if list.Next.Next.Val != 3 {
		t.Errorf("Expected 3, got %d", list.Next.Next.Val)
	}

	if list2.Val != 0 {
		t.Errorf("Expected 0, got %d", list2.Val)
	}

	if list2.Next.Val != 6 {
		t.Errorf("Expected 6, got %d", list2.Next.Val)
	}

	if list2.Next.Next.Val != 8 {
		t.Errorf("Expected 8, got %d", list2.Next.Next.Val)
	}

}

func Test_FromListNodeToArray(t *testing.T) {
	a := &models.ListNode{
		Val: 5,
		Next: &models.ListNode{
			Val: 6,
			Next: &models.ListNode{
				Val: 8,
			},
		}}

	array := FromListNodeToArray(a)

	if array[0] != 5 {
		t.Errorf("Expected 5, got %d", array[0])
	}

	if array[1] != 6 {
		t.Errorf("Expected 6, got %d", array[1])
	}

	if array[2] != 8 {
		t.Errorf("Expected 8, got %d", array[2])
	}

}

func Test_MergeList(t *testing.T) {
	a := []int{2, 3, 4}
	b := []int{1, 4, 6}

	c := []int{1, 2, 3, 4, 4, 6}

	listA := FromArrayToListNode(a)
	listB := FromArrayToListNode(b)

	merged := MergeListNode(listA, listB)

	array := FromListNodeToArray(merged)

	for i, v := range array {
		if v != c[i] {
			t.Errorf("Expected %d, got %d", c[i], v)
		}
	}
}

func Test_FromListNodeToString(t *testing.T) {
	a := []int{1, 2, 3}

	list := FromArrayToListNode(a)

	stringList := FromListNodeToString(list)

	if stringList != "1,2,3" {
		t.Errorf("Expected '1,2,3', got %s", stringList)
	}
}

func Test_FromStringToListNode(t *testing.T) {
	a := "{1,2,3}"

	list := FromStringToListNode(a)

	if list.Val != 1 {
		t.Errorf("Expected 1, got %d", list.Val)
	}

	if list.Next.Val != 2 {
		t.Errorf("Expected 2, got %d", list.Next.Val)
	}

	if list.Next.Next.Val != 3 {
		t.Errorf("Expected 3, got %d", list.Next.Next.Val)
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