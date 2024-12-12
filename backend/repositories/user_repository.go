package repositories

import (
	"errors"
	"yona-backend/models"
)

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
		},
	}
}

func (r *UserRepository) GetAll() []models.User {
	return r.users
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *UserRepository) Create(user models.User) models.User {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user
}

func (r *UserRepository) Update(id int, updatedUser models.User) (models.User, error) {
	for i, user := range r.users {
		if user.ID == id {
			r.users[i] = updatedUser
			return updatedUser, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *UserRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
