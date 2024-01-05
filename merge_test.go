package main

import (
	"testing"
)

func Test_FromArrayToListNode(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	list := FromArrayToListNode(a)

	if list.Val != 1 {
		t.Errorf("Expected 1, got %d", list.Val)
	}

	if list.Next.Val != 2 {
		t.Errorf("Expected 2, got %d", list.Next.Val)
	}

	if list.Next.Next.Val != 3 {
		t.Errorf("Expected 3, got %d", list.Next.Next.Val)
	}

	if list.Next.Next.Next.Val != 4 {
		t.Errorf("Expected 4, got %d", list.Next.Next.Next.Val)
	}

	if list.Next.Next.Next.Next.Val != 5 {
		t.Errorf("Expected 5, got %d", list.Next.Next.Next.Next.Val)
	}
}