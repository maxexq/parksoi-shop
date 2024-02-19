package usersUsecases

import (
	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/users"
	"github.com/maxexq/parksoi-shop/modules/users/usersRepositories"
)

type IUsersUsecase interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
}

type usersUsecase struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UsersUsecase(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}

func (u *usersUsecase) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}

	result, err := u.usersRepository.InsertUser(req, false)

	if err != nil {
		return nil, err
	}

	return result, nil
}
