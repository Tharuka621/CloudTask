package repositories

import (
	"cloudtask/task-service/internal/models"
	DB "cloudtask/task-service/pkg/database"
)

type TaskRepository interface {
	CreateProject(project *models.Project) error
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	GetTasksByTeam(teamID uint, status string) ([]models.Task, error)
	DeleteTask(id uint) error
	UpdateTaskStatus(taskID uint, status models.TaskStatus) error
}

type taskRepository struct{}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (r *taskRepository) CreateProject(project *models.Project) error {
	return DB.DB.Create(project).Error
}

func (r *taskRepository) CreateTask(task *models.Task) error {
	return DB.DB.Create(task).Error
}

func (r *taskRepository) UpdateTask(task *models.Task) error {
	return DB.DB.Save(task).Error
}

func (r *taskRepository) GetTasksByTeam(teamID uint, status string) ([]models.Task, error) {
	var tasks []models.Task
	query := DB.DB.Where("team_id = ?", teamID)
	
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) DeleteTask(id uint) error {
	return DB.DB.Delete(&models.Task{}, id).Error
}

func (r *taskRepository) UpdateTaskStatus(taskID uint, status models.TaskStatus) error {
	return DB.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("status", status).Error
}
