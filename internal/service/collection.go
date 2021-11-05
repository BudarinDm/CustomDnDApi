package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type CollectionService struct {
	repo repository.Collection
}

func NewCollectionService(repo repository.Collection) *CollectionService {
	return &CollectionService{repo: repo}
}

func (s *CollectionService) Create(userId int, collection model.Collection) (int, error) {
	return s.repo.Create(userId, collection)
}

func (s *CollectionService) GetAllTheType(userId int, collectionType string) ([]model.Collection, error) {
	return s.repo.GetAllTheType(userId, collectionType)
}

func (s *CollectionService) GetCollectionById(userId int, id int) (model.Collection, error) {
	return s.repo.GetCollectionById(userId, id)
}

func (s *CollectionService) UpdateCollection(userId int, id int, collection model.Collection) error {
	return s.repo.UpdateCollection(userId, id, collection)
}

func (s *CollectionService) DeleteCollection(userId int, id int) error {
	return s.repo.DeleteCollection(userId, id)
}
