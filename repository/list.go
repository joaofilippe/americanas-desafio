package repository

import (
	"github.com/joaofilippe/americanas-desafio/list_node"
	"github.com/joaofilippe/americanas-desafio/models"
)

// SelectLists is a function to select all lists from the database based on the id
func (r *Repository)SelectLists(id int) ([]*models.ListNode, error){
	q := `
	SELECT *
	FROM public.list_node
	WHERE id = $1
	`
	listsDB := []ListsDB{}
	err := r.Db.Select(&listsDB, q, id)
	if err != nil {
		return nil, err
	}

	list1 := list_node.FromStringToListNode(listsDB[0].List1.String)
	list2 := list_node.FromStringToListNode(listsDB[0].List2.String)

	return []*models.ListNode{list1, list2}, nil
}

// InsertList is a function to insert a list into the database
func (r *Repository)InsertLists(list1, list2 models.ListNode) (int64, error){
	q := `
	INSERT INTO public.list_node(
		list_1, list_2)
		VALUES ('{$1}', '{$2}');
	`
	l1 := list_node.FromListNodeToString(&list1)
	l2 := list_node.FromListNodeToString(&list2)

	result, err := r.Db.Exec( q, l1, l2)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}