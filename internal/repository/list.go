package repository

import (
	"github.com/joaofilippe/americanas-desafio/internal/common"
	"github.com/joaofilippe/americanas-desafio/internal/list_node"
	"github.com/joaofilippe/americanas-desafio/internal/models"
)

// CreateListNodeTable is a function to create the list_node table in the database
func (r *Repository) CreateListNodeTable() error {
	q := `
		CREATE TABLE IF NOT EXISTS list_node
		(
			id     int auto_increment,
			list_1 varchar(255) not null,
			list_2 varchar(255) not null,
			merged varchar(600) null,
			constraint list_node_pk
				primary key (id)
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
	FROM list_node.list_node
	WHERE id = ?
	`
	listsDB := []ListsDB{}
	err := r.Db.Select(&listsDB, q, id)
	if err != nil {
		return nil, err
	}

	list1 := listnode.FromStringToListNode(listsDB[0].List1.String)
	list2 := listnode.FromStringToListNode(listsDB[0].List2.String)

	return []*models.ListNode{list1, list2}, nil
}

// InsertLists is a function to insert a list into the database
func (r *Repository) InsertLists(lists []*models.ListNode) (int64, error) {
	q := `
	INSERT INTO list_node.list_node (list_1, list_2)
	VALUES (?,?);
	`

	if len(lists) != 2 {
		return 0, common.ErrInvalidNumberOfLists
	}

	var listStrings []string
	for _, list := range lists {
		listString := listnode.FromListNodeToString(list)
		listStrings = append(listStrings, listString)
	}

	res, err := r.Db.Exec(q, listStrings[0], listStrings[1])
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// UpdateMergedList is a function to update the merged list in the database
func (r *Repository) UpdateMergedList(mergedList models.ListNode, id int64) error {
	q := `
		UPDATE list_node.list_node
		SET merged = ?
		WHERE id = ?;
	`
	ml := listnode.FromListNodeToString(&mergedList)

	_, err := r.Db.Exec(q, ml, id)

	return err
}
