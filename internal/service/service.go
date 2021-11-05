package service

import (
	"DnDApi/internal/model"
	"DnDApi/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Collection interface {
	Create(userId int, collection model.Collection) (int, error)
	GetAllTheType(userId int, collectionType string) ([]model.Collection, error)
	GetCollectionById(userId int, id int) (model.Collection, error)
	UpdateCollection(userId int, id int, collection model.Collection) error
	DeleteCollection(userId int, id int) error
}

type Weapon interface {
	Create(userId int, weapon model.Weapon, collectionId int) (int, error)
	GetAllWeaponByCollectionId(userId int, collectionId int) ([]model.Weapon, error)
	GetWeaponById(userId int, id int) (model.Weapon, error)
	UpdateWeapon(userId int, id int, weapon model.Weapon) error
	DeleteWeapon(userId int, id int) error
}

type Spell interface {
	Create(userId int, spell model.Spell, collectionId int) (int, error)
	GetAllSpellByCollectionId(userId int, collectionId int) ([]model.Spell, error)
	GetSpellById(userId int, id int) (model.Spell, error)
	UpdateSpell(userId int, id int, spell model.Spell) error
	DeleteSpell(userId int, id int) error
}

type Armor interface {
	Create(userId int, armor model.Armor, collectionId int) (int, error)
	GetAllArmorByCollectionId(userId int, collectionId int) ([]model.Armor, error)
	GetArmorById(userId int, id int) (model.Armor, error)
	UpdateArmor(userId int, id int, armor model.Armor) error
	DeleteArmor(userId int, id int) error
}

type TalentCategory interface {
	Create(userId int, talentCategory model.TalentCategory, collectionId int) (int, error)
	GetAllTalentCategoryByCollectionId(userId int, collectionId int) ([]model.TalentCategory, error)
	GetTalentCategoryById(userId int, id int) (model.TalentCategory, error)
	UpdateTalentCategory(userId int, id int, talentCategory model.TalentCategory) error
	DeleteTalentCategory(userId int, id int) error
}

type Race interface {
	Create(userId int, race model.Race, collectionId int) (int, error)
	GetAllRaceByCollectionId(userId int, collectionId int) ([]model.Race, error)
	GetRaceById(userId int, id int) (model.Race, error)
	UpdateRace(userId int, id int, race model.Race) error
	DeleteRace(userId int, id int) error
}

type Class interface {
	Create(userId int, class model.Class, collectionId int) (int, error)
	GetAllClassByCollectionId(userId int, collectionId int) ([]model.Class, error)
	GetClassById(userId int, id int) (model.Class, error)
	UpdateClass(userId int, id int, class model.Class) error
	DeleteClass(userId int, id int) error
}

type Talent interface {
	Create(userId int, talent model.Talent, categoryId int) (int, error)
	GetAllTalentByCategoryId(userId int, categoryId int) ([]model.Talent, error)
	GetTalentById(userId int, id int) (model.Talent, error)
	UpdateTalent(userId int, id int, talent model.Talent) error
	DeleteTalent(userId int, id int) error
}

type Service struct {
	Authorization
	Collection
	Weapon
	Spell
	Armor
	TalentCategory
	Race
	Class
	Talent
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		Collection:     NewCollectionService(repos.Collection),
		Weapon:         NewWeaponService(repos.Weapon),
		Spell:          NewSpellService(repos.Spell),
		Armor:          NewArmorService(repos.Armor),
		TalentCategory: NewTalentCategoryService(repos.TalentCategory),
		Race:           NewRaceService(repos.Race),
		Class:          NewClassService(repos.Class),
		Talent:         NewTalentService(repos.Talent),
	}
}
