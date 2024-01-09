package listnode

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/joaofilippe/americanas-desafio/internal/common"
	"github.com/joaofilippe/americanas-desafio/internal/mocks"
	"github.com/joaofilippe/americanas-desafio/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_SaveListsNode(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return error when list1 is not sorted", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		list1 := &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 2,
					Next: &models.ListNode{
						Val: 4,
					},
				},
			},
		}

		list2 := &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 5,
					Next: &models.ListNode{
						Val: 6,
					},
				},
			},
		}

		_, err := service.SaveListsNode(list1, list2)
		assert.Equal(common.ErrListNotSorted, err, "should return error when list1 is not sorted")
	})

	t.Run("should return error when list2 is not sorted", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		list1 := &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 3,
					Next: &models.ListNode{
						Val: 4,
					},
				},
			},
		}

		list2 := &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 1,
					Next: &models.ListNode{
						Val: 6,
					},
				},
			},
		}

		_, err := service.SaveListsNode(list1, list2)
		assert.Equal(common.ErrListNotSorted, err, "should return error when list2 is not sorted")
	})

	t.Run("should return error when repository returns error", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		// TODO
		repoMock.On("InsertLists",
			mock.Anything).
			Return(int64(0), sql.ErrNoRows)

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		list1 := &models.ListNode{}
		list2 := &models.ListNode{}

		id, err := service.SaveListsNode(list1, list2)
		assert.Zero(id, "should return error when repository returns error")
		assert.Equal(sql.ErrNoRows, err, "should return error when repository returns error")
	})

	t.Run("should return id when repository returns id", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		repoMock.On("InsertLists",
			mock.Anything).
			Return(int64(1), nil)

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		list1 := &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 3,
				},
			},
		}

		list2 := &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 4,
				},
			},
		}

		id, err := service.SaveListsNode(list1, list2)
		assert.Equal(int64(1), id, "should return id when repository returns id")
		assert.Nil(err, "should not return error when repository returns id")
	})
}

func Test_MergeListNode(t *testing.T){
	assert := assert.New(t)

	t.Run("should return error when repository cannot select lists", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		repoMock.On("SelectLists",
			mock.Anything).
			Return([]*models.ListNode{}, errors.New("error on select lists - err.Error()"))

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		id, err := service.GetMergedListNode(1)
		assert.Zero(id, "should 0 error when repository returns error")
		assert.ErrorContains(err, "error on select lists", "should return error when repository returns error")
	})

	t.Run("should return error when repository cannot save merged list", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		repoMock.On("SelectLists",
			mock.Anything).
			Return([]*models.ListNode{}, errors.New("error on update merged list - err.Error()"))

		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		id, err := service.GetMergedListNode(1)
		assert.Zero(id, "should 0 error when repository returns error")
		assert.ErrorContains(err, "error on update merged list", "should return error when repository returns error")
	})

	t.Run("should return merged list", func(t *testing.T) {
		repoMock := &mocks.IListNodeRepository{}
		serviceMock := &mocks.IListNodeService{}

		list1 := &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 3,
				},
			},
		}

		list2 := &models.ListNode{
			Val: 2,
			Next: &models.ListNode{
				Val: 3,
				Next: &models.ListNode{
					Val: 4,
				},
			},
		}

		mergedList:= &models.ListNode{
			Val: 1,
			Next: &models.ListNode{
				Val: 2,
				Next: &models.ListNode{
					Val: 3,
					Next: &models.ListNode{
						Val: 3,
						Next: &models.ListNode{
							Val: 4,
							Next: &models.ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		}
		
		serviceMock.On("GetMergedListNode", mock.Anything).Return(mergedList, nil)

		repoMock.On("SelectLists",mock.Anything).Return([]*models.ListNode{list1, list2}, nil)
		repoMock.On("UpdateMergedList", mock.Anything, mock.Anything).Return(nil)
		
		service := Service{
			Service:    serviceMock,
			Repository: repoMock,
		}

		mergedList, err := service.GetMergedListNode(1)
		assert.NotNil(mergedList, "should return merged list")
		assert.Nil(err, "should not return error when repository returns id")
	})
}