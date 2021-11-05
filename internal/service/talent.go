package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type TalentService struct {
	repo repository.Talent
}

func NewTalentService(repo repository.Talent) *TalentService {
	return &TalentService{repo: repo}
}

func (s *TalentService) Create(userId int, talent model.Talent, categoryId int) (int, error) {
	return s.repo.Create(userId, talent, categoryId)
}

func (s *TalentService) GetAllTalentByCategoryId(userId int, categoryId int) ([]model.Talent, error) {
	return s.repo.GetAllTalentByCategoryId(userId, categoryId)
}

func (s *TalentService) GetTalentById(userId int, id int) (model.Talent, error) {
	return s.repo.GetTalentById(userId, id)
}

func (s *TalentService) UpdateTalent(userId int, id int, talent model.Talent) error {
	return s.repo.UpdateTalent(userId, id, talent)
}

func (s *TalentService) DeleteTalent(userId int, id int) error {
	return s.repo.DeleteTalent(userId, id)
}
