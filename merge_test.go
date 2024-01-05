package main

import (
	"testing"
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