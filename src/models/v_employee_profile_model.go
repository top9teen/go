package models

import (
	"config"
	"entitys"
)

func All() ([]entitys.V_employee_profile, error) {

	db := config.DBsql()
	var emp []entitys.V_employee_profile
	err := db.Order("emp_id desc").Limit(-1).Find(&emp).Error

	return emp, err
}

func Count(number int64) ([]entitys.V_employee_profile, error) {
	db := config.DBsql()
	var emp []entitys.V_employee_profile
	err := db.Order("emp_id desc").Limit(number).Find(&emp).Error

	return emp, err
}
