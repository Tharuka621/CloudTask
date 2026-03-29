package services

import (
	"errors"

	"cloudtask/team-service/internal/models"
	"cloudtask/team-service/internal/repositories"
)

type TeamService interface {
	CreateTeam(name, description string, ownerID uint) (*models.Team, error)
	AddMember(teamID, requesterID, newMemberID uint, role models.TeamRole) error
	GetUserTeams(userID uint) ([]models.Team, error)
}

type teamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(repo repositories.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) CreateTeam(name, description string, ownerID uint) (*models.Team, error) {
	team := &models.Team{
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}

	if err := s.repo.CreateTeam(team); err != nil {
		return nil, errors.New("failed to create team")
	}

	// Add owner as a team member
	member := &models.TeamMember{
		TeamID: team.ID,
		UserID: ownerID,
		Role:   models.RoleOwner,
	}
	_ = s.repo.AddMember(member)

	return team, nil
}

func (s *teamService) AddMember(teamID, requesterID, newMemberID uint, role models.TeamRole) error {
	team, err := s.repo.GetTeamByID(teamID)
	if err != nil {
		return errors.New("team not found")
	}

	if team.OwnerID != requesterID {
		isOwnerOrAdmin := false
		for _, m := range team.Members {
			if m.UserID == requesterID && (m.Role == models.RoleOwner || m.Role == models.RoleAdmin) {
				isOwnerOrAdmin = true
				break
			}
		}
		if !isOwnerOrAdmin {
			return errors.New("unauthorized to add members")
		}
	}

	member := &models.TeamMember{
		TeamID: teamID,
		UserID: newMemberID,
		Role:   role,
	}

	return s.repo.AddMember(member)
}

func (s *teamService) GetUserTeams(userID uint) ([]models.Team, error) {
	return s.repo.GetTeamsByUserID(userID)
}
