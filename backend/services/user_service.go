package services

import (
	"yona-backend/models"
	"yona-backend/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []models.User {
	return s.repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(user models.User) models.User {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id int, user models.User) (models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
