package health

import "goyo/models"

type HealthResult struct {
	models.Model
	Status string `json:"status"`
}
