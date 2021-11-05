package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type ClassService struct {
	repo repository.Class
}

func NewClassService(repo repository.Class) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) Create(userId int, сlass model.Class, collectionId int) (int, error) {
	return s.repo.Create(userId, сlass, collectionId)
}

func (s *ClassService) GetAllClassByCollectionId(userId int, collectionId int) ([]model.Class, error) {
	return s.repo.GetAllClassByCollectionId(userId, collectionId)
}

func (s *ClassService) GetClassById(userId int, id int) (model.Class, error) {
	return s.repo.GetClassById(userId, id)
}

func (s *ClassService) UpdateClass(userId int, id int, сlass model.Class) error {
	return s.repo.UpdateClass(userId, id, сlass)
}

func (s *ClassService) DeleteClass(userId int, id int) error {
	return s.repo.DeleteClass(userId, id)
}
