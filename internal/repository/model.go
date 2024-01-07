package repository

import "database/sql"

// ListsDB is a struct to use in the database
type ListsDB struct {
	ID     sql.NullInt64  `db:"id"`
	List1  sql.NullString `db:"list_1"`
	List2  sql.NullString `db:"list_2"`
	Merged sql.NullString `db:"merged_list"`
}
