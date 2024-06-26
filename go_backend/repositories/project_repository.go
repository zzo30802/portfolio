package repositories

import (
	"go_backend/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func (r *ProjectRepository) GetAll() ([]models.Project, error) {
	var projects []models.Project
	result := r.DB.Find(&projects)
	return projects, result.Error
}

func (r *ProjectRepository) GetByID(id uint) (models.Project, error) {
	var project models.Project
	result := r.DB.First(&project, id)
	return project, result.Error
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return r.DB.Save(project).Error
}

func (r *ProjectRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Project{}, id).Error
}
