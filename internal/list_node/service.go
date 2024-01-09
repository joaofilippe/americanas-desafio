package listnode

import (
	"github.com/joaofilippe/americanas-desafio/internal/common"
	"github.com/joaofilippe/americanas-desafio/internal/interfaces"
	"github.com/joaofilippe/americanas-desafio/internal/models"
)

// Service is a struct to be used as a service
type Service struct {
	Service    interfaces.IListNodeService
	Repository interfaces.IListNodeRepository
}

// SaveListsNode save two lists in database
func (s *Service) SaveListsNode(list1, list2 *models.ListNode) (int64, error) {
	validList1 := ValidateSorted(list1)
	if !validList1 {
		return 0, common.ErrListNotSorted
	}

	validList2 := ValidateSorted(list2)
	if !validList2 {
		return 0, common.ErrListNotSorted
	}

	id, err := s.Repository.InsertLists([]*models.ListNode{list1, list2})
	if err != nil {
		return 0, err
	}

	return id, nil
}

// MergeListNode merge two sorted linked lists and return it as a sorted list.
func (s *Service) MergeListNode(id int64) (*models.ListNode, error) {
	lists, err := s.Repository.SelectLists(id)
	if err != nil {
		return nil, err
	}

	mergedList := MergeListNode(lists[0], lists[1])

	err = s.Repository.UpdateMergedList(*mergedList, id)
	if err != nil {
		return nil, err
	}

	return mergedList, nil
}
