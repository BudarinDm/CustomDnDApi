package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type ArmorService struct {
	repo repository.Armor
}

func NewArmorService(repo repository.Armor) *ArmorService {
	return &ArmorService{repo: repo}
}

func (s *ArmorService) Create(userId int, armor model.Armor, collectionId int) (int, error) {
	return s.repo.Create(userId, armor, collectionId)
}

func (s *ArmorService) GetAllArmorByCollectionId(userId int, collectionId int) ([]model.Armor, error) {
	return s.repo.GetAllArmorByCollectionId(userId, collectionId)
}

func (s *ArmorService) GetArmorById(userId int, id int) (model.Armor, error) {
	return s.repo.GetArmorById(userId, id)
}

func (s *ArmorService) UpdateArmor(userId int, id int, armor model.Armor) error {
	return s.repo.UpdateArmor(userId, id, armor)
}

func (s *ArmorService) DeleteArmor(userId int, id int) error {
	return s.repo.DeleteArmor(userId, id)
}
