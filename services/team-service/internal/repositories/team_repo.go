package repositories

import (
	"cloudtask/team-service/internal/models"
	DB "cloudtask/team-service/pkg/database"
)

type TeamRepository interface {
	CreateTeam(team *models.Team) error
	GetTeamByID(id uint) (*models.Team, error)
	AddMember(member *models.TeamMember) error
	GetTeamsByUserID(userID uint) ([]models.Team, error)
}

type teamRepository struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}

func (r *teamRepository) CreateTeam(team *models.Team) error {
	return DB.DB.Create(team).Error
}

func (r *teamRepository) GetTeamByID(id uint) (*models.Team, error) {
	var team models.Team
	if err := DB.DB.Preload("Members").First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) AddMember(member *models.TeamMember) error {
	return DB.DB.Create(member).Error
}

func (r *teamRepository) GetTeamsByUserID(userID uint) ([]models.Team, error) {
	var members []models.TeamMember
	if err := DB.DB.Where("user_id = ?", userID).Find(&members).Error; err != nil {
		return nil, err
	}

	var teamIDs []uint
	for _, m := range members {
		teamIDs = append(teamIDs, m.TeamID)
	}

	var teams []models.Team
	if len(teamIDs) > 0 {
		DB.DB.Where("id IN ?", teamIDs).Find(&teams)
	}
	return teams, nil
}
