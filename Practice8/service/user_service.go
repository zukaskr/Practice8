package service

import (
	"errors"
	"practice8/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) UpdateUserName(id int, name string) error {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Name = name
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	if id == 1 {
		return errors.New("it is not allowed to delete admin user")
	}
	return s.repo.DeleteUser(id)
}
