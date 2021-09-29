package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name  string `json:"department_name" gorm:"unique"`
	Phone string `json:"department_phone"`
}
