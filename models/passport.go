package models

import "gorm.io/gorm"

type Passport struct {
	gorm.Model
	Type   string `json:"passport_type"`
	Number string `json:"passport_number"`
}
