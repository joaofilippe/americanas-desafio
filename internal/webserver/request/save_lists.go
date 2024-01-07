package request

import "github.com/joaofilippe/americanas-desafio/models"

// SaveListsRequest is a request to save two lists
type SaveListsRequest struct {
	Type  string           `json:"type"`
	List1 *models.ListNode `json:"list1"`
	List2 *models.ListNode `json:"list2"`
}
