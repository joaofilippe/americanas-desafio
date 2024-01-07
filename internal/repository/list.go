package repository

import (
	"github.com/lib/pq"

	"github.com/joaofilippe/americanas-desafio/internal/list_node"
	"github.com/joaofilippe/americanas-desafio/internal/models"
)

// CreateListNodeTable is a function to create the list_node table in the database
func (r *Repository) CreateListNodeTable() error {
	q := `
	CREATE TABLE IF NOT EXISTS public.table_name
	(
		id     bigint[] not null
			constraint id_table_name_pk
				primary key,
		list_1 integer,
		list_2 bigint[]
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
	if err != nil {
		return nil, err
	}

	list1 := listNode.FromStringToListNode(listsDB[0].List1.String)
	list2 := listNode.FromStringToListNode(listsDB[0].List2.String)

	return []*models.ListNode{list1, list2}, nil
}

// InsertLists is a function to insert a list into the database
func (r *Repository) InsertLists(list1, list2 models.ListNode) (int64, error) {
	q := `
	INSERT INTO public.list_node(
		list_1, list_2)
		VALUES ($1, $2);
	`
	l1 := listNode.FromListNodeToArray(&list1)
	l2 := listNode.FromListNodeToArray(&list2)

	_, err := r.Db.Exec(q, pq.Array(l1), pq.Array(l2))
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
