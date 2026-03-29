package handlers

import (
	"cloudtask/task-service/internal/models"
	"cloudtask/task-service/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	reporterID := uint(c.Locals("user_id").(float64))

	type Request struct {
		TeamID      uint                `json:"team_id"`
		ProjectID   uint                `json:"project_id"`
		Title       string              `json:"title"`
		Description string              `json:"description"`
		Priority    models.TaskPriority `json:"priority"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	task, err := h.service.CreateTask(req.TeamID, req.ProjectID, reporterID, req.Title, req.Description, req.Priority)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	teamID, _ := strconv.ParseUint(c.Query("team_id"), 10, 32)
	status := c.Query("status")

	tasks, err := h.service.ListTasks(uint(teamID), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}
