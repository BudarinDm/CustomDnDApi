package postgres

import (
	"DnDApi/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Collection{}, &model.CollectionType{}, &model.Weapon{}, &model.WeaponType{},
		&model.Spell{}, &model.SpellType{}, &model.Armor{}, &model.ArmorType{}, &model.Race{}, &model.Class{}, &model.TalentCategory{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
