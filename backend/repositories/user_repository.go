package repositories

import (
	"yona-backend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	// Auto-migrate the schema
	db.AutoMigrate(&models.User{})
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	r.db.Find(&users)
	return users
}

func (r *UserRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) Update(id uint, updatedUser models.User) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := r.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
