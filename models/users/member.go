package users

import "goyo/models"

type Member struct {
	models.Primary
	Name string
}
