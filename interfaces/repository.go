package interfaces

import "github.com/joaofilippe/americanas-desafio/models"

// IListNodeRepository is an interface to the repository
type IListNodeRepository interface {
	SelectLists(id int) ([]*models.ListNode, error)
	InsertLists(list1, list2 models.ListNode) (int64, error)
}
