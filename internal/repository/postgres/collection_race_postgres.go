package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type RacePostgres struct {
	db *gorm.DB
}

func NewRacePostgres(db *gorm.DB) *RacePostgres {
	return &RacePostgres{db: db}
}

func (r *RacePostgres) Create(userId int, race model.Race, collectionId int) (int, error) {
	race.CollectionId = collectionId
	result := r.db.Create(&race)
	if result.Error != nil {
		return 0, result.Error
	}

	id := race.Id

	return id, nil
}

func (r *RacePostgres) GetAllRaceByCollectionId(userId int, collectionId int) ([]model.Race, error) {
	var races []model.Race

	result := r.db.Where("collection_id = ?", collectionId).Find(&races)
	if result.Error != nil {
		return nil, result.Error
	}

	return races, nil
}

func (r *RacePostgres) GetRaceById(userId int, id int) (model.Race, error) {
	var race model.Race

	result := r.db.First(&race, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Race{}, result.Error
	}

	return race, nil
}

func (r *RacePostgres) UpdateRace(userId int, id int, race model.Race) error {
	var updateRace model.Race
	getResult := r.db.First(&updateRace, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if race.Title != "" {
		updateRace.Title = race.Title
	}
	if race.Description != "" {
		updateRace.Description = race.Description
	}
	if race.History != "" {
		updateRace.History = race.History
	}
	if race.Bonus != "" {
		updateRace.Bonus = race.Bonus
	}

	updResult := r.db.Save(&updateRace)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *RacePostgres) DeleteRace(userId int, id int) error {
	result := r.db.Delete(&model.Race{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
