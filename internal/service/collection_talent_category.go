package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type TalentCategoryService struct {
	repo repository.TalentCategory
}

func NewTalentCategoryService(repo repository.TalentCategory) *TalentCategoryService {
	return &TalentCategoryService{repo: repo}
}

func (s *TalentCategoryService) Create(userId int, talentCategory model.TalentCategory, collectionId int) (int, error) {
	return s.repo.Create(userId, talentCategory, collectionId)
}

func (s *TalentCategoryService) GetAllTalentCategoryByCollectionId(userId int, collectionId int) ([]model.TalentCategory, error) {
	return s.repo.GetAllTalentCategoryByCollectionId(userId, collectionId)
}

func (s *TalentCategoryService) GetTalentCategoryById(userId int, id int) (model.TalentCategory, error) {
	return s.repo.GetTalentCategoryById(userId, id)
}

func (s *TalentCategoryService) UpdateTalentCategory(userId int, id int, talentCategory model.TalentCategory) error {
	return s.repo.UpdateTalentCategory(userId, id, talentCategory)
}

func (s *TalentCategoryService) DeleteTalentCategory(userId int, id int) error {
	return s.repo.DeleteTalentCategory(userId, id)
}
