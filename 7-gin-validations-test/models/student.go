package models

import (
	"gopkg.in/validator.v2"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"max=11"`
	RG   string `json:"rg" validate:"max=9""`
}

func ValidateStudentData(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
}
