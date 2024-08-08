package repository

import (
	"github.com/task_manager/Domain"
)

type UserRepository interface{
	CreateUser(data *domain.User) error
	ReadUser(username string) (domain.User, bool)
	ChangeRoleToAdmin(username string) error
	UpdateUser(username string, data *domain.User) error
	NumberOfUsers() (int64, error)
}

