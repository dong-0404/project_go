package employee

import (
	"demo_project/model"
)

func (er *EmployeeRepositories) Update(id uint, data interface{}) error {
	err := db.Model(&model.TblEmployee{}).
		Where("id = ?", id).
		Updates(data).Error
	return err
}

func (er *EmployeeRepositories) UpdateEmployeeInfo(id int, employeeUpdate *model.TblEmployee) error {
	return db.Model(model.TblEmployee{}).Where("id = ?", id).Updates(employeeUpdate).Error
}

func (er *EmployeeRepositories) UpdateEmployeeDocs(id int, employeeDocsUpdate *model.EmployeeDocs) error {
	// Đặt giá trị của employee_id
	employeeDocsUpdate.EmployeeId = id
	return db.Model(model.EmployeeDocs{}).Where("employee_id = ?", id).Updates(employeeDocsUpdate).Error
}
