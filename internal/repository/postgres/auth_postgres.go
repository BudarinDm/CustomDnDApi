package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db     *gorm.DB
	DBName string
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r AuthPostgres) CreateUser(user model.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	id := user.Id

	return id, nil
}

func (r AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User

	result := r.db.First(&user, "user_name =? AND password = ?", username, password)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.User{}, result.Error
	}

	return user, nil
}
