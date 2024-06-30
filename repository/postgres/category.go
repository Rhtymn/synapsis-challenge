package repository

import "database/sql"

type categoryRepositoryPostgres struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepositoryPostgres {
	return &categoryRepositoryPostgres{
		db: db,
	}
}
