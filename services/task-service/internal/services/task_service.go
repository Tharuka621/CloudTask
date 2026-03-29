package services

import (
	"cloudtask/task-service/internal/models"
	"cloudtask/task-service/internal/repositories"
)

type TaskService interface {
	CreateProject(teamID uint, name, description string) (*models.Project, error)
	CreateTask(teamID, projectID, reporterID uint, title, description string, priority models.TaskPriority) (*models.Task, error)
	ListTasks(teamID uint, status string) ([]models.Task, error)
	UpdateTaskStatus(taskID uint, status models.TaskStatus) error
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateProject(teamID uint, name, description string) (*models.Project, error) {
	project := &models.Project{
		TeamID:      teamID,
		Name:        name,
		Description: description,
	}
	err := s.repo.CreateProject(project)
	return project, err
}

func (s *taskService) CreateTask(teamID, projectID, reporterID uint, title, description string, priority models.TaskPriority) (*models.Task, error) {
	if priority == "" {
		priority = models.PriorityMedium
	}
	task := &models.Task{
		TeamID:      teamID,
		ProjectID:   projectID,
		ReporterID:  reporterID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Status:      models.StatusTodo,
	}
	err := s.repo.CreateTask(task)
	return task, err
}

func (s *taskService) ListTasks(teamID uint, status string) ([]models.Task, error) {
	return s.repo.GetTasksByTeam(teamID, status)
}

func (s *taskService) UpdateTaskStatus(taskID uint, status models.TaskStatus) error {
	return s.repo.UpdateTaskStatus(taskID, status)
}
