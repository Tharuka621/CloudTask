package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string
type TaskPriority string

const (
	StatusTodo       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"

	PriorityLow    TaskPriority = "LOW"
	PriorityMedium TaskPriority = "MEDIUM"
	PriorityHigh   TaskPriority = "HIGH"
)

type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	TeamID      uint           `gorm:"index;not null" json:"team_id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ProjectID   uint           `gorm:"index;not null" json:"project_id"`
	TeamID      uint           `gorm:"index;not null" json:"team_id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Status      TaskStatus     `gorm:"type:varchar(20);default:'TODO'" json:"status"`
	Priority    TaskPriority   `gorm:"type:varchar(20);default:'MEDIUM'" json:"priority"`
	DueDate     *time.Time     `json:"due_date"`
	AssigneeID  *uint          `json:"assignee_id"`
	ReporterID  uint           `gorm:"not null" json:"reporter_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type TaskComment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TaskID    uint           `gorm:"index;not null" json:"task_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Content   string         `gorm:"not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
