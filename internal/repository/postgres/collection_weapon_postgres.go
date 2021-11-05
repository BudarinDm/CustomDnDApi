package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type WeaponPostgres struct {
	db *gorm.DB
}

func NewWeaponPostgres(db *gorm.DB) *WeaponPostgres {
	return &WeaponPostgres{db: db}
}

func (r *WeaponPostgres) Create(userId int, weapon model.Weapon, collectionId int) (int, error) {
	weapon.CollectionId = collectionId
	result := r.db.Create(&weapon)
	if result.Error != nil {
		return 0, result.Error
	}

	id := weapon.Id

	return id, nil
}

func (r *WeaponPostgres) GetAllWeaponByCollectionId(userId int, collectionId int) ([]model.Weapon, error) {
	var weapons []model.Weapon

	result := r.db.Where("collection_id = ?", collectionId).Find(&weapons)
	if result.Error != nil {
		return nil, result.Error
	}

	return weapons, nil
}

func (r *WeaponPostgres) GetWeaponById(userId int, id int) (model.Weapon, error) {
	var weapon model.Weapon

	result := r.db.First(&weapon, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Weapon{}, result.Error
	}

	return weapon, nil
}

func (r *WeaponPostgres) UpdateWeapon(userId int, id int, weapon model.Weapon) error {
	var updateWeapon model.Weapon
	getResult := r.db.First(&updateWeapon, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if weapon.Title != "" {
		updateWeapon.Title = weapon.Title
	}
	if weapon.Description != "" {
		updateWeapon.Description = weapon.Description
	}
	if weapon.Damage != 0 {
		updateWeapon.Damage = weapon.Damage
	}
	if weapon.Quality != 0 {
		updateWeapon.Quality = weapon.Quality
	}
	if weapon.Range != 0 {
		updateWeapon.Range = weapon.Range
	}
	if weapon.Bonus != "" {
		updateWeapon.Bonus = weapon.Bonus
	}

	updResult := r.db.Save(&updateWeapon)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *WeaponPostgres) DeleteWeapon(userId int, id int) error {
	result := r.db.Delete(&model.Weapon{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
