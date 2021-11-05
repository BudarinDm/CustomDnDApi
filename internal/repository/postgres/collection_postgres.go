package postgres

import (
	"DnDApi/internal/model"
	"errors"
	"gorm.io/gorm"
)

type CollectionPostgres struct {
	db *gorm.DB
}

func NewCollectionPostgres(db *gorm.DB) *CollectionPostgres {
	return &CollectionPostgres{db: db}
}

func (r *CollectionPostgres) Create(userId int, collection model.Collection) (int, error) {
	collection.AuthorId = userId
	result := r.db.Create(&collection)
	if result.Error != nil {
		return 0, result.Error
	}

	id := collection.Id

	return id, nil
}

func (r *CollectionPostgres) GetAllTheType(userId int, collectionType string) ([]model.Collection, error) {
	var collections []model.Collection

	result := r.db.Where("collection_type_id = ?", collectionType).Find(&collections)
	if result.Error != nil {
		return nil, result.Error
	}

	return collections, nil
}

func (r *CollectionPostgres) GetCollectionById(userId int, id int) (model.Collection, error) {
	var collection model.Collection

	result := r.db.First(&collection, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.Collection{}, result.Error
	}

	return collection, nil
}

func (r *CollectionPostgres) UpdateCollection(userId int, id int, collection model.Collection) error {
	var updateCollection model.Collection
	getResult := r.db.First(&updateCollection, id)
	if getResult.Error != nil {
		return getResult.Error
	}

	if collection.AuthorId != 0 {
		updateCollection.AuthorId = collection.AuthorId
	}
	if collection.Title != "" {
		updateCollection.Title = collection.Title
	}
	if collection.Description != "" {
		updateCollection.Description = collection.Description
	}

	updResult := r.db.Save(&updateCollection)
	if updResult.Error != nil {
		return updResult.Error
	}

	return nil
}

func (r *CollectionPostgres) DeleteCollection(userId int, id int) error {
	result := r.db.Delete(&model.Collection{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
