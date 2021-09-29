package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Phone        string     `json:"phone"`
	CompanyId    int        `json:"company_id"`
	Passport     Passport   `json:"passport"`
	PassportId   int        `json:"passport_id"`
	Department   Department `json:"department"`
	DepartmentId int        `json:"department_id"`
}
