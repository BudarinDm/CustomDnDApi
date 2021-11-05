package model

type User struct {
	Id       int    `json:"id" gorm:" primaryKey" gorm:"<-:create"`
	Name     string `json:"name"`
	UserName string `json:"username" gorm:"<-:create" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	PhotoUrl string `json:"photo_url"`
}
