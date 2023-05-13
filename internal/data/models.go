package data

import "database/sql"

type Models struct {
	Stories StoryModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Stories: StoryModel{DB: db},
	}
}
