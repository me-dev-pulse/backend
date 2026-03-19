package database

import (
	"github.com/me-dev-pulse/backend/src/models"
)

func CreateProject(name, url string) (*models.Project, error) {
	query := `INSERT INTO projects (name, url) VALUES ($1, $2) RETURNING id, name, url, enabled, created_at`

	var p models.Project
	err := DB.QueryRowx(query, name, url).StructScan(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}