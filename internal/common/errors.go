package common

import "errors"

var (

	// ErrListNotSorted is the error code for list not sorted
	ErrListNotSorted        = errors.New("list not sorted")
	
	// ErrInvalidNumberOfLists is the error code for invalid number of lists
	ErrInvalidNumberOfLists = errors.New("invalid number of lists")

	// ErrListNotFound is the error code for lists not found
	ErrListsNotFound		 = errors.New("lists not found")
)

const (
	// ErrMsgContentNotBinded is the error message for content not binded 
	ErrMsgContentNotBinded = "server could not bind the content"
	
	// ErrMsgCannotSaveLists is the error message for cannot save lists
	ErrMsgCannotSaveLists = "server could not save the lists"

	// ErrMsgCannotMergeLists is the error message for cannot merge lists
	ErrMsgCannotMergeLists = "server could not merge the lists"
)