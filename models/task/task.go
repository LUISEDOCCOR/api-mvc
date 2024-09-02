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

func GetById(id uint, user_id uint) (models.Task, error) {
	var task models.Task
	result := database.DB.Where("id = ? AND user_id = ?", id, user_id).First(&task, id)
	return task, result.Error
}

func Delete(id uint) error {
	result := database.DB.Delete(&models.Task{}, id)
	return result.Error
}

func GetAll(user_id uint) ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Where("user_id = ?", user_id).Find(&tasks)
	return tasks, result.Error
}

func GetAllByStatus(user_id uint, isDone bool) ([]models.Task, error) {
	var tasks []models.Task
	result := database.DB.Where("user_id = ? AND is_done = ?", user_id, isDone).Find(&tasks)
	return tasks, result.Error
}

func Update(task *models.Task) error {
	result := database.DB.Save(&task)
	return result.Error
}
