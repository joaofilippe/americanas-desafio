package listNode

import (
	"github.com/joaofilippe/americanas-desafio/interfaces"
	"github.com/joaofilippe/americanas-desafio/models"
)

// ListNode is a struct to the service
type ListNodeService struct {
	Service    interfaces.IListNodeService
	Repository interfaces.IListNodeRepository
}

// SaveListsNode save two lists in database
func (l *ListNodeService) SaveListsNode(list1, list2 *models.ListNode) *models.ListNode {
	return nil
}

// MergeListNode merge two sorted linked lists and return it as a sorted list.
func (l *ListNodeService) MergeListNode(list1, list2 *models.ListNode) *models.ListNode {
	return nil
}
