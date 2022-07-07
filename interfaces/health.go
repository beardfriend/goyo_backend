package interfaces

type HealthResult struct {
	Model
	Status string `json:"status"`
}

func (HealthResult) TableName() string {
	return "healths"
}
