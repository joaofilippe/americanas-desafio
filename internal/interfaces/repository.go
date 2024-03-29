package interfaces

import "github.com/joaofilippe/americanas-desafio/internal/models"

// IListNodeRepository is an interface to the repository
type IListNodeRepository interface {
	SelectLists(id int64) ([]*models.ListNode, error)
	InsertLists(lists []*models.ListNode) (int64, error)
	UpdateMergedList(mergedList models.ListNode, id int64) error
}
