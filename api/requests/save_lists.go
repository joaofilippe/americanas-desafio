package requests

// ListNodeRequest is a node in a singly-linked list.
type ListNodeRequest struct {
	Val  int              `json:"val"`
	Next *ListNodeRequest `json:"next"`
}

// SaveListsRequest is a request to save two lists
type SaveListsRequest struct {
	Type  string           `json:"type"`
	List1 *ListNodeRequest `json:"list1"`
	List2 *ListNodeRequest `json:"list2"`
}
