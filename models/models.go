package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	Task     []Task `gorm:"foreignKey:UserId" json:"tasks"`
}

type Task struct {
	gorm.Model

	Title   string `gorm:"not null" json:"title"`
	Content string `json:"content"`
	IsDone  bool   `gorm:"not null" json:"isDone"`
	UserId  uint   `gorm:"not null" json:"user_id"`
}
