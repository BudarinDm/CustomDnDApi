package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type ClassPostgres struct {
	db *gorm.DB
}

func NewClassPostgres(db *gorm.DB) *ClassPostgres {
	return &ClassPostgres{db: db}
}

func (r *ClassPostgres) Create(userId int, class model.Class, collectionId int) (int, error) {
	class.CollectionId = collectionId
	result := r.db.Create(&class)
	if result.Error != nil {
		return 0, result.Error
	}

	id := class.Id

	return id, nil
}

func (r *ClassPostgres) GetAllClassByCollectionId(userId int, collectionId int) ([]model.Class, error) {
	var classes []model.Class

	result := r.db.Where("collection_id = ?", collectionId).Find(&classes)
	if result.Error != nil {
		return nil, result.Error
	}

	return classes, nil
}

func (r *ClassPostgres) GetClassById(userId int, id int) (model.Class, error) {
	var class model.Class

	result := r.db.First(&class, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Class{}, result.Error
	}

	return class, nil
}

func (r *ClassPostgres) UpdateClass(userId int, id int, class model.Class) error {
	var updateClass model.Class
	getResult := r.db.First(&updateClass, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if class.Title != "" {
		updateClass.Title = class.Title
	}
	if class.Description != "" {
		updateClass.Description = class.Description
	}
	if class.Bonus != "" {
		updateClass.Bonus = class.Bonus
	}

	updResult := r.db.Save(&updateClass)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *ClassPostgres) DeleteClass(userId int, id int) error {
	result := r.db.Delete(&model.Class{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
