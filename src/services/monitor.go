package services

import (
	"log"
	"net/http"
	"time"

	"github.com/me-dev-pulse/backend/src/database"
	"github.com/me-dev-pulse/backend/src/models"
)

func StartMonitor(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			checkAllProjects()
		}
	}()
}

func checkAllProjects() {
	var projects []models.Project
	err := database.DB.Select(&projects, "SELECT * FROM projects WHERE enabled = true")
	if err != nil {
		log.Printf("Error getting projects: %v", err)
		return
	}

	for _, p := range projects {
		go pingAndSave(p)
	}
}

func pingAndSave(p models.Project) {
	start := time.Now()
	
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(p.URL)
	latency := int(time.Since(start).Milliseconds())

	isUp := false
	statusCode := 0

	if err == nil {
		statusCode = resp.StatusCode
		if statusCode >= 200 && statusCode < 300 {
			isUp = true
		}
		resp.Body.Close()
	}

	_, dbErr := database.DB.Exec(
		"INSERT INTO checks (project_id, status_code, latency_ms, is_up) VALUES ($1, $2, $3, $4)",
		p.ID, statusCode, latency, isUp,
	)

	if dbErr != nil {
		log.Printf("❌ Error saving check for %s: %v", p.Name, dbErr)
	} else {
		statusIcon := "✅"
		if !isUp {
			statusIcon = "❌"
		}
		log.Printf("%s %s - %d ms (Status: %d)", statusIcon, p.Name, latency, statusCode)
	}
}