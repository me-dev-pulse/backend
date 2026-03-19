package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/me-dev-pulse/backend/src/database"
	"github.com/me-dev-pulse/backend/src/models"
	"github.com/me-dev-pulse/backend/src/services"
)

func GetProjects(c *fiber.Ctx) error {
	var projects []models.Project
	err := database.DB.Select(&projects, "SELECT * FROM projects ORDER BY created_at DESC")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(projects)
}

func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var project models.Project
	err := database.DB.Get(&project, "SELECT * FROM projects WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var summary models.ProjectSummary
	
	var uptime float64
	query := `
		SELECT (COUNT(CASE WHEN is_up = true THEN 1 END) * 100.0 / COUNT(*))
		FROM checks 
		WHERE project_id = $1 AND created_at > NOW() - INTERVAL '24 hours'`
	
	database.DB.Get(&uptime, query, id)

	var lastStatus bool
	database.DB.Get(&lastStatus, "SELECT is_up FROM checks WHERE project_id = $1 ORDER BY created_at DESC LIMIT 1", id)

	summary = models.ProjectSummary{
		Project:       project,
		CurrentStatus: lastStatus,
		Uptime24h:     uptime,
		SSLExpiryDays: services.GetSSLExpiryDays(project.URL),
	}


	return c.JSON(summary)
}

func GetProjectStats(c *fiber.Ctx) error {
	id := c.Params("id")
	
	var checks []models.Check
	query := `SELECT * FROM checks WHERE project_id = $1 ORDER BY created_at DESC LIMIT 60`
	err := database.DB.Select(&checks, query, id)
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(checks)
}

func CreateProjectHandler(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	p, err := database.CreateProject(req.Name, req.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create"})
	}
	return c.Status(201).JSON(p)
}

func DeleteProjectHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete"})
	}
	return c.SendStatus(204)
}

func GetProjectsSummary(c *fiber.Ctx) error {
	var projects []models.Project
	database.DB.Select(&projects, "SELECT * FROM projects")

	var summaries []models.ProjectSummary

	for _, p := range projects {
		var uptime float64
		query := `
			SELECT (COUNT(CASE WHEN is_up = true THEN 1 END) * 100.0 / COUNT(*))
			FROM checks 
			WHERE project_id = $1 AND created_at > NOW() - INTERVAL '24 hours'`
		
		database.DB.Get(&uptime, query, p.ID)

		var lastStatus bool
		database.DB.Get(&lastStatus, "SELECT is_up FROM checks WHERE project_id = $1 ORDER BY created_at DESC LIMIT 1", p.ID)

		summaries = append(summaries, models.ProjectSummary{
			Project:       p,
			CurrentStatus: lastStatus,
			Uptime24h:     uptime,
			SSLExpiryDays: services.GetSSLExpiryDays(p.URL),
		})
	}

	return c.JSON(summaries)
}