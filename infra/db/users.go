package db

import (
	"github.com/nandobas/twitter/domain/user"
)

type UserRepositoryMock struct {
	CreateFunc      func(u user.User) error
	GetUserByIDFunc func(userID string) (user.User, error)
}

func (ur *UserRepositoryMock) Create(u user.User) error {
	if ur.CreateFunc != nil {
		return ur.CreateFunc(u)
	}
	return nil
}

func (ur *UserRepositoryMock) GetUserByID(userID string) (user.User, error) {
	if ur.GetUserByIDFunc != nil {
		return ur.GetUserByIDFunc(userID)
	}
	return user.User{}, nil
}
