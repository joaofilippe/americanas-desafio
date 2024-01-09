package listnode

import (
	"errors"

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

// GetMergedListNode merge two sorted linked lists and return it as a sorted list.
func (s *Service) GetMergedListNode(id int64) (*models.ListNode, error) {
	lists, err := s.Repository.SelectLists(id)
	if err != nil {
		return nil, errors.New("error on select lists - " + err.Error())
	}

	if len(lists) > 2 && lists[2] != nil {
		return lists[2], nil
	}
	
	mergedList := MergeListNode(lists[0], lists[1])

	err = s.Repository.UpdateMergedList(*mergedList, id)
	if err != nil {
		return nil, errors.New("error on update merged list - " + err.Error())
	}

	return mergedList, nil
}
