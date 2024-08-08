package repository

import (
	"github.com/task_manager/Domain"
)

type UserRepository interface{
	CreateUser(data *models.User) error
	ReadUser(username string) (models.User, bool)
	ChangeRoleToAdmin(username string) error
	UpdateUser(username string, data *models.User) error
	NumberOfUsers() (int64, error)
}

