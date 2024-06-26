package services

import (
	"go_backend/models"
	"go_backend/repositories"
)

type ProjectService struct {
	repo *repositories.ProjectRepository
}

func NewProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	return s.repo.GetAll()
}

func (s *ProjectService) GetProjectByID(id uint) (models.Project, error) {
	return s.repo.GetByID(id)
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	return s.repo.Create(project)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
	return s.repo.Update(project)
}

func (s *ProjectService) DeleteProject(id uint) error {
	return s.repo.Delete(id)
}
