package task_model

import (
	"github.com/LUISEDOCCOR/api-mvc/database"
	"github.com/LUISEDOCCOR/api-mvc/models"
)

//CRUD

func Create(task *models.Task) error {
	result := database.DB.Create(task)
	return result.Error
}

func GetById(id uint) (models.Task, error) {
	var task models.Task
	result := database.DB.Select("id, title, content, isDone").First(&task, id)
	return task, result.Error
}

func Delete(id uint) error {
	result := database.DB.Delete(models.Task{}, id)
	return result.Error
}

func GetAll(user_id uint) ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Where("user_id", user_id).Find(&tasks)
	return tasks, result.Error
}

//TODO ✅ UPDATE
