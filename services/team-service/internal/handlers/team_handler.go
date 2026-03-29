package handlers

import (
	"cloudtask/team-service/internal/models"
	"cloudtask/team-service/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TeamHandler struct {
	service services.TeamService
}

func NewTeamHandler(service services.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) CreateTeam(c *fiber.Ctx) error {
	userIDFloat := c.Locals("user_id").(float64)
	userID := uint(userIDFloat)

	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	team, err := h.service.CreateTeam(req.Name, req.Description, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(team)
}

func (h *TeamHandler) AddMember(c *fiber.Ctx) error {
	requesterIDFloat := c.Locals("user_id").(float64)
	requesterID := uint(requesterIDFloat)

	teamIDParam := c.Params("id")
	teamID, err := strconv.ParseUint(teamIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid team ID"})
	}

	type Request struct {
		UserID uint            `json:"user_id"`
		Role   models.TeamRole `json:"role"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if req.Role == "" {
		req.Role = models.RoleMember
	}

	err = h.service.AddMember(uint(teamID), requesterID, req.UserID, req.Role)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "member added successfully"})
}

func (h *TeamHandler) GetMyTeams(c *fiber.Ctx) error {
	userIDFloat := c.Locals("user_id").(float64)
	userID := uint(userIDFloat)

	teams, err := h.service.GetUserTeams(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(teams)
}
