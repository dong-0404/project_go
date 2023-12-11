package employee

import (
	"demo_project/database"
	"demo_project/model"
)

func (cre *EmployeeRepositories) GetAll() ([]model.TblEmployee, error) {
	db := database.GetInstance().GetDB()
	var employees []model.TblEmployee
	err := db.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}
