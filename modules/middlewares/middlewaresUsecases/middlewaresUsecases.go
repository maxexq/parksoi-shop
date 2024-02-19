package middlewaresUsecases

import (
	middlewaresrepositories "github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecases interface {
}

type middlewaresUsecases struct {
	middlewaresrepository middlewaresrepositories.IMiddlewaresrepositories
}

func MiddlewaresRepository(middlewaresrepository middlewaresrepositories.IMiddlewaresrepositories) IMiddlewaresUsecases {
	return &middlewaresUsecases{
		middlewaresrepository: middlewaresrepository,
	}
}
