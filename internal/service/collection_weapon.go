package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type WeaponService struct {
	repo repository.Weapon
}

func NewWeaponService(repo repository.Weapon) *WeaponService {
	return &WeaponService{repo: repo}
}

func (s *WeaponService) Create(userId int, weapon model.Weapon, collectionId int) (int, error) {
	return s.repo.Create(userId, weapon, collectionId)
}

func (s *WeaponService) GetAllWeaponByCollectionId(userId int, collectionId int) ([]model.Weapon, error) {
	return s.repo.GetAllWeaponByCollectionId(userId, collectionId)
}

func (s *WeaponService) GetWeaponById(userId int, id int) (model.Weapon, error) {
	return s.repo.GetWeaponById(userId, id)
}

func (s *WeaponService) UpdateWeapon(userId int, id int, weapon model.Weapon) error {
	return s.repo.UpdateWeapon(userId, id, weapon)
}

func (s *WeaponService) DeleteWeapon(userId int, id int) error {
	return s.repo.DeleteWeapon(userId, id)
}
