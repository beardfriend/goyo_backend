package health

import "goyo/models"

type HealthResult struct {
	models.Primary
	Status string `json:"status"`
	models.TimeWithDeleted
}
