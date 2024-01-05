package main

import (
	"fmt"

	"github.com/joaofilippe/americanas-desafio/models"
)

func main() {
	fmt.Println("Hello World")

	a := &models.ListNode{
		Val: 5,
		Next: &models.ListNode{
			Val: 6,
			Next: &models.ListNode{
				Val: 8,
			},
		}}

	b := &models.ListNode{
		Val: 2,
		Next: &models.ListNode{
			Val: 5,
			Next: &models.ListNode{
				Val: 6,
				Next: &models.ListNode{
					Val:10,
				},
			},
		},
	}

	list := MergeListNode(a, b)

	fmt.Println(list)

}


