package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type ArmorPostgres struct {
	db *gorm.DB
}

func NewArmorPostgres(db *gorm.DB) *ArmorPostgres {
	return &ArmorPostgres{db: db}
}

func (r *ArmorPostgres) Create(userId int, armor model.Armor, collectionId int) (int, error) {
	armor.CollectionId = collectionId
	result := r.db.Create(&armor)
	if result.Error != nil {
		return 0, result.Error
	}

	id := armor.Id

	return id, nil
}

func (r *ArmorPostgres) GetAllArmorByCollectionId(userId int, collectionId int) ([]model.Armor, error) {
	var armors []model.Armor

	result := r.db.Where("collection_id = ?", collectionId).Find(&armors)
	if result.Error != nil {
		return nil, result.Error
	}

	return armors, nil
}

func (r *ArmorPostgres) GetArmorById(userId int, id int) (model.Armor, error) {
	var armor model.Armor

	result := r.db.First(&armor, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Armor{}, result.Error
	}

	return armor, nil
}

func (r *ArmorPostgres) UpdateArmor(userId int, id int, armor model.Armor) error {
	var updateArmor model.Armor
	getResult := r.db.First(&updateArmor, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if armor.Title != "" {
		updateArmor.Title = armor.Title
	}
	if armor.Description != "" {
		updateArmor.Description = armor.Description
	}
	if armor.Armor != 0 {
		updateArmor.Armor = armor.Armor
	}
	if armor.Effect != "" {
		updateArmor.Effect = armor.Effect
	}

	updResult := r.db.Save(&updateArmor)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *ArmorPostgres) DeleteArmor(userId int, id int) error {
	result := r.db.Delete(&model.Armor{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
