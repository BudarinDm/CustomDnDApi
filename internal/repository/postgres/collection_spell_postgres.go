package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type SpellPostgres struct {
	db *gorm.DB
}

func NewSpellPostgres(db *gorm.DB) *SpellPostgres {
	return &SpellPostgres{db: db}
}

func (r *SpellPostgres) Create(userId int, spell model.Spell, collectionId int) (int, error) {
	spell.CollectionId = collectionId
	result := r.db.Create(&spell)
	if result.Error != nil {
		return 0, result.Error
	}

	id := spell.Id

	return id, nil
}

func (r *SpellPostgres) GetAllSpellByCollectionId(userId int, collectionId int) ([]model.Spell, error) {
	var spells []model.Spell

	result := r.db.Where("collection_id = ?", collectionId).Find(&spells)
	if result.Error != nil {
		return nil, result.Error
	}

	return spells, nil
}

func (r *SpellPostgres) GetSpellById(userId int, id int) (model.Spell, error) {
	var spell model.Spell

	result := r.db.First(&spell, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Spell{}, result.Error
	}

	return spell, nil
}

func (r *SpellPostgres) UpdateSpell(userId int, id int, spell model.Spell) error {
	var updateSpell model.Spell
	getResult := r.db.First(&updateSpell, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if spell.Title != "" {
		updateSpell.Title = spell.Title
	}
	if spell.Description != "" {
		updateSpell.Description = spell.Description
	}
	if spell.Range != 0 {
		updateSpell.Range = spell.Range
	}
	if spell.Effect != "" {
		updateSpell.Effect = spell.Effect
	}

	updResult := r.db.Save(&updateSpell)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *SpellPostgres) DeleteSpell(userId int, id int) error {
	result := r.db.Delete(&model.Spell{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
