package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	Task     []Task `json:"tasks"`
}

type Task struct {
	gorm.Model

	Title   string `gorm:"not null" json:"title"`
	Contetn string `json:"content"`
	IsDone  bool   `gorm:"not null" json:"isDone"`
	UserId  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserId;not null" json:"user"`
}
