package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type TalentCategoryPostgres struct {
	db *gorm.DB
}

func NewTalentCategoryPostgres(db *gorm.DB) *TalentCategoryPostgres {
	return &TalentCategoryPostgres{db: db}
}

func (r *TalentCategoryPostgres) Create(userId int, talentCategory model.TalentCategory, collectionId int) (int, error) {
	talentCategory.CollectionId = collectionId
	result := r.db.Create(&talentCategory)
	if result.Error != nil {
		return 0, result.Error
	}

	id := talentCategory.Id

	return id, nil
}

func (r *TalentCategoryPostgres) GetAllTalentCategoryByCollectionId(userId int, collectionId int) ([]model.TalentCategory, error) {
	var talentCategories []model.TalentCategory

	result := r.db.Where("collection_id = ?", collectionId).Find(&talentCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	return talentCategories, nil
}

func (r *TalentCategoryPostgres) GetTalentCategoryById(userId int, id int) (model.TalentCategory, error) {
	var talentCategory model.TalentCategory

	result := r.db.First(&talentCategory, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.TalentCategory{}, result.Error
	}

	return talentCategory, nil
}

func (r *TalentCategoryPostgres) UpdateTalentCategory(userId int, id int, talentCategory model.TalentCategory) error {
	var updateTalentCategory model.TalentCategory
	getResult := r.db.First(&updateTalentCategory, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if talentCategory.Title != "" {
		updateTalentCategory.Title = talentCategory.Title
	}
	if talentCategory.Description != "" {
		updateTalentCategory.Description = talentCategory.Description
	}

	updResult := r.db.Save(&updateTalentCategory)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *TalentCategoryPostgres) DeleteTalentCategory(userId int, id int) error {
	result := r.db.Delete(&model.TalentCategory{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
