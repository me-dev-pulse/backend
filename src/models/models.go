package models

import "time"

type Project struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	URL       string    `db:"url" json:"url"`
	Enabled   bool      `db:"enabled" json:"enabled"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Check struct {
	ID         int       `db:"id" json:"id"`
	ProjectID  int       `db:"project_id" json:"project_id"`
	StatusCode int       `db:"status_code" json:"status_code"`
	LatencyMs  int       `db:"latency_ms" json:"latency_ms"`
	IsUp       bool      `db:"is_up" json:"is_up"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type ProjectSummary struct {
	Project
	CurrentStatus bool    `json:"current_status"`
	Uptime24h     float64 `json:"uptime_24h"`
	SSLExpiryDays int     `json:"ssl_expiry_days"`
}