package interfaces

import "github.com/joaofilippe/americanas-desafio/models"

// IListNodeService is an interface to the service
type IListNodeService interface {
	SaveListsNode(list1, list2 *models.ListNode) (int64, error)
	MergeListNode(id int64) (*models.ListNode, error)
}