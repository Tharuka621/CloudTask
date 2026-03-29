package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamRole string

const (
	RoleOwner  TeamRole = "Owner"
	RoleAdmin  TeamRole = "Admin"
	RoleMember TeamRole = "Member"
)

type Team struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	OwnerID     uint           `gorm:"not null" json:"owner_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Members     []TeamMember   `json:"members,omitempty"`
}

type TeamMember struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TeamID    uint           `gorm:"index;not null" json:"team_id"`
	UserID    uint           `gorm:"index;not null" json:"user_id"`
	Role      TeamRole       `gorm:"type:varchar(20);default:'Member'" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
