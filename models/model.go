package models

import "gorm.io/gorm"

type SocioGram struct {
	gorm.Model
	Post  string
	Image string
}