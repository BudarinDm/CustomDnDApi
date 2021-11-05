package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type TalentPostgres struct {
	db *gorm.DB
}

func NewTalentPostgres(db *gorm.DB) *TalentPostgres {
	return &TalentPostgres{db: db}
}

func (r *TalentPostgres) Create(userId int, talent model.Talent, categoryId int) (int, error) {
	talent.CategoryId = categoryId
	result := r.db.Create(&talent)
	if result.Error != nil {
		return 0, result.Error
	}

	id := talent.Id

	return id, nil
}

func (r *TalentPostgres) GetAllTalentByCategoryId(userId int, categoryId int) ([]model.Talent, error) {
	var talents []model.Talent

	result := r.db.Where("category_id = ?", categoryId).Find(&talents)
	if result.Error != nil {
		return nil, result.Error
	}

	return talents, nil
}

func (r *TalentPostgres) GetTalentById(userId int, id int) (model.Talent, error) {
	var talent model.Talent

	result := r.db.First(&talent, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Talent{}, result.Error
	}

	return talent, nil
}

func (r *TalentPostgres) UpdateTalent(userId int, id int, talent model.Talent) error {
	var updateTalent model.Talent
	getResult := r.db.First(&updateTalent, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if talent.Title != "" {
		updateTalent.Title = talent.Title
	}
	if talent.Description != "" {
		updateTalent.Description = talent.Description
	}
	if talent.Condition != "" {
		updateTalent.Condition = talent.Condition
	}
	if talent.Benefit != "" {
		updateTalent.Benefit = talent.Benefit
	}

	updResult := r.db.Save(&updateTalent)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *TalentPostgres) DeleteTalent(userId int, id int) error {
	result := r.db.Delete(&model.Talent{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
