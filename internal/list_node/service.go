package listNode

import (
	"github.com/joaofilippe/americanas-desafio/internal/interfaces"
	"github.com/joaofilippe/americanas-desafio/internal/models"
)

// Service is a struct to the service
type Service struct {
	Service    interfaces.IListNodeService
	Repository interfaces.IListNodeRepository
}

// SaveListsNode save two lists in database
func (s *Service) SaveListsNode(list1, list2 *models.ListNode) (int64, error) {
	id, err := s.Repository.InsertLists(*list1, *list2)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// MergeListNode merge two sorted linked lists and return it as a sorted list.
func (s *Service) MergeListNode(id int64) (*models.ListNode,error) {
	lists, err := s.Repository.SelectLists(id)
	if err != nil {
		return &models.ListNode{}, err
	}

	mergedList := MergeListNode(lists[0], lists[1])

	return mergedList, nil
}
