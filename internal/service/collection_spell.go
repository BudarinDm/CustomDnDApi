package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type SpellService struct {
	repo repository.Spell
}

func NewSpellService(repo repository.Spell) *SpellService {
	return &SpellService{repo: repo}
}

func (s *SpellService) Create(userId int, spell model.Spell, collectionId int) (int, error) {
	return s.repo.Create(userId, spell, collectionId)
}

func (s *SpellService) GetAllSpellByCollectionId(userId int, collectionId int) ([]model.Spell, error) {
	return s.repo.GetAllSpellByCollectionId(userId, collectionId)
}

func (s *SpellService) GetSpellById(userId int, id int) (model.Spell, error) {
	return s.repo.GetSpellById(userId, id)
}

func (s *SpellService) UpdateSpell(userId int, id int, spell model.Spell) error {
	return s.repo.UpdateSpell(userId, id, spell)
}

func (s *SpellService) DeleteSpell(userId int, id int) error {
	return s.repo.DeleteSpell(userId, id)
}
