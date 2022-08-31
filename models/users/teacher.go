package users

import "goyo/models"

type Teacher struct {
	models.Primary
	Name string
}
