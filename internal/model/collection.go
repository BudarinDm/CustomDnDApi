package model

type Collection struct {
	Id               int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	AuthorId         int    `json:"author_id" gorm:"not null"`
	Title            string `json:"title" gorm:"not null"`
	Description      string `json:"description" type:"text"`
	CollectionTypeId int    `json:"collection_type_id" gorm:"not null"`
}

type CollectionType struct {
	Id    int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	Title string `json:"title" gorm:"not null"`
}

type Weapon struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	WeaponTypeId int    `json:"weapon_type_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
	Damage       int    `json:"damage"`
	Quality      int    `json:"quality"`
	Range        int    `json:"range"`
	Bonus        string `json:"bonus" type:"text"`
}

type WeaponType struct {
	Id    int    `json:"id" gorm:" primaryKey" gorm:"<-:create"`
	Title string `json:"title" gorm:"not null"`
}

type Spell struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	SpellTypeId  int    `json:"weapon_type_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
	Range        int    `json:"range"`
	Effect       string `json:"effect" type:"text"`
}

type SpellType struct {
	Id    int    `json:"id" gorm:" primaryKey" gorm:"<-:create"`
	Title string `json:"title" gorm:"not null"`
}

type Armor struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	ArmorTypeId  int    `json:"weapon_type_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
	Armor        int    `json:"armor"`
	Effect       string `json:"effect" type:"text"`
}

type ArmorType struct {
	Id    int    `json:"id" gorm:" primaryKey" gorm:"<-:create"`
	Title string `json:"title" gorm:"not null"`
}

type Race struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
	History      string `json:"history" type:"text"`
	Bonus        string `json:"bonus" type:"text"`
}

type Class struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
	Bonus        string `json:"bonus" type:"text"`
}
