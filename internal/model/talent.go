package model

type TalentCategory struct {
	Id           int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CollectionId int    `json:"collection_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description" type:"text"`
}

type Talent struct {
	Id          int    `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	CategoryId  int    `json:"category_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" type:"text"`
	Condition   string `json:"condition" type:"text"`
	Benefit     string `json:"benefit" type:"text"`
}
