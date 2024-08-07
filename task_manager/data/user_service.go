package service

import (
	"task_manager/models"
    "errors"
)


func (s *Services) GetUser(username *string) (models.User, error) {
	user, ok := s.Connection.ReadUser(*username)

    if !ok{
        return models.User{}, errors.New("the user not found")
    }

    return user, nil
}


func (s *Services) PromoteUser(username string) error{
    return s.Connection.ChangeRoleToAdmin(username)
}

func (s *Services) AddUser(user *models.User) error{
    currentUsers, err := s.Connection.NumberOfUsers()
    if err != nil{
        return err
    }

    if currentUsers == 0{
        user.Role = "Admin"
    }
	return s.Connection.CreateUser(user)
}


func (s *Services) UpdateUser(id *string, updatedUser *models.User) error {
	return s.Connection.UpdateUser(*id, updatedUser)
}