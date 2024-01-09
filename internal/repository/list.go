package repository

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/joaofilippe/americanas-desafio/internal/common"
	"github.com/joaofilippe/americanas-desafio/internal/list_node"
	"github.com/joaofilippe/americanas-desafio/internal/models"
)

// CreateListNodeTable is a function to create the list_node table in the database
func (r *Repository) CreateListNodeTable() error {
	q := `
		CREATE TABLE IF NOT EXISTS public.list_node
		(
		    id     serial   not null
		        constraint list_node_pk
		            primary key,
		    list_1 bigint[] not null,
		    list_2 bigint[] not null,
		    merged bigint[]
		);
	`

	_, err := r.Db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

// SelectLists is a function to select all lists from the database based on the id
func (r *Repository) SelectLists(id int64) ([]*models.ListNode, error) {
	q := `
	SELECT *
	FROM public.list_node
	WHERE id = $1
	`
	listsDB := []ListsDB{}
	err := r.Db.Select(&listsDB, q, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if len(listsDB) == 0 {
		return nil, common.ErrListsNotFound
	}
	
	list1 := listnode.FromStringToListNode(listsDB[0].List1.String)
	list2 := listnode.FromStringToListNode(listsDB[0].List2.String)
	
	if listsDB[0].Merged.Valid {
		merged := listnode.FromStringToListNode(listsDB[0].Merged.String)
		return []*models.ListNode{list1, list2, merged}, nil
	}

	return []*models.ListNode{list1, list2}, nil
}

// InsertLists is a function to insert a list into the database
func (r *Repository) InsertLists(lists []*models.ListNode) (int64, error) {
	q := `
	INSERT INTO public.list_node(
		list_1, list_2)
		VALUES ($1, $2);
	`

	if len(lists) != 2 {
		return 0, common.ErrInvalidNumberOfLists
	}

	var arrays [][]int
	for _, list := range lists {
		array := listnode.FromListNodeToSlice(list)
		arrays = append(arrays, array)
	}

	_, err := r.Db.Exec(q, pq.Array(arrays[0]), pq.Array(arrays[1]))
	if err != nil {
		return 0, err
	}

	var id int64
	qLastID := `
		SELECT id
		FROM public.list_node
		ORDER BY id DESC
		LIMIT 1
	`

	err = r.Db.Get(&id, qLastID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateMergedList is a function to update the merged list in the database
func (r *Repository) UpdateMergedList(mergedList models.ListNode, id int64) error {
	q := `
		UPDATE public.list_node
		SET merged = $1
		WHERE id = $2;
	`
	ml := listnode.FromListNodeToSlice(&mergedList)

	_, err := r.Db.Exec(q, pq.Array(ml), id)
	if err != nil {
		return err
	}

	return nil
}
