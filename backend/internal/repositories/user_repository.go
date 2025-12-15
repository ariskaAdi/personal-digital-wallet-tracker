package repositories

import (
	"ariskaAdi/personal-digital-wallet/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.Users) (entity.Users, error)
    Update(ctx context.Context, user entity.Users) (entity.Users, error)
    Delete(ctx context.Context, id int) error
    FindAll(ctx context.Context) ([]entity.Users, error)
    FindById(ctx context.Context, id int) (entity.Users, error)
	FindByEmail(ctx context.Context, email string) (entity.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

// Composition Root
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user entity.Users) (entity.Users, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	return  user, err
}

func (r *userRepository) Update(ctx context.Context, user entity.Users) (entity.Users, error) {
	res := r.db.WithContext(ctx).Model(&entity.Users{}).Where("id = ?", user.ID).Updates(user)

	if res.Error != nil {
		return entity.Users{}, res.Error
	}

	if res.RowsAffected == 0 {
		return entity.Users{}, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	res := r.db.WithContext(ctx).Delete(&entity.Users{}, id)
		if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.WithContext(ctx).Find(&users).Error
	return  users, err
}

func (r *userRepository) FindById(ctx context.Context, id int) (entity.Users, error) {
	var user entity.Users
	err := r.db.WithContext(ctx).First(&user, id).Error
	return  user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (entity.Users, error) {
	var user entity.Users
	err := r.db.WithContext(ctx).Where("email = ?", email ).First(&user).Error
	return  user, err
}