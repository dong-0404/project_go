package employee

import "demo_project/model"

func (er *EmployeeRepositories) GetDetailsEmployee() ([]model.TblEmployee, error) {
	var employees []model.TblEmployee
	err := db.
		Model(&model.TblEmployee{}).
		Preload("EmployeeDocs").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}
