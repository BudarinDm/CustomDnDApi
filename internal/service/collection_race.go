package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type RaceService struct {
	repo repository.Race
}

func NewRaceService(repo repository.Race) *RaceService {
	return &RaceService{repo: repo}
}

func (s *RaceService) Create(userId int, race model.Race, collectionId int) (int, error) {
	return s.repo.Create(userId, race, collectionId)
}

func (s *RaceService) GetAllRaceByCollectionId(userId int, collectionId int) ([]model.Race, error) {
	return s.repo.GetAllRaceByCollectionId(userId, collectionId)
}

func (s *RaceService) GetRaceById(userId int, id int) (model.Race, error) {
	return s.repo.GetRaceById(userId, id)
}

func (s *RaceService) UpdateRace(userId int, id int, race model.Race) error {
	return s.repo.UpdateRace(userId, id, race)
}

func (s *RaceService) DeleteRace(userId int, id int) error {
	return s.repo.DeleteRace(userId, id)
}
