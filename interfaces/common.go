package interfaces

import (
	"goyo/libs/types"
)

type Model struct {
	ID        uint             `json:"id"`
	CreatedAt types.TimeString `json:"created_at"`
	UpdatedAt types.TimeString `json:"updated_at"`
}
