package middlewaresrepositories

import "github.com/jmoiron/sqlx"

type IMiddlewaresrepositories interface {
}

type middlewaresrepositories struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMiddlewaresrepositories {
	return &middlewaresrepositories{
		db: db,
	}
}
