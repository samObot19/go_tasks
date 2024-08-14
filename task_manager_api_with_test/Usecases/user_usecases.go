package usecase

import (
	"github.com/task_manager/Domain"
	"github.com/task_manager/Repositories"
	"errors"
)

type UserUsecase struct{
	db repository.UserRepository
}


func NewUserUsecase(con repository.UserRepository) *UserUsecase{
	return &UserUsecase{
		db : con,
	}
}

func (s *UserUsecase) GetUser(username *string) (domain.User, error) {
	user, ok := s.db.ReadUser(*username)

	if !ok{
		return domain.User{}, errors.New("the user not found")
	}
	return user, nil
}

func (s *UserUsecase) PromoteUser(username string) error{
	return s.db.ChangeRoleToAdmin(username)
}

func (s *UserUsecase) AddUser(user *domain.User) error{
	currentUsers, err := s.db.NumberOfUsers()

	if err != nil{
		return err
    }

    if currentUsers == 0{
        user.Role = "Admin"
    }
	
	return s.db.CreateUser(user)
}

func (s *UserUsecase) UpdateUser(id *string, updatedUser *domain.User) error {
	return s.db.UpdateUser(*id, updatedUser)
}





