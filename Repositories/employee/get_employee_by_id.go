package employee

import (
	"demo_project/model"
)

//	func (er *EmployeeRepositories) GetByID(id uint) (interface{}, error) {
//		db := database.GetInstance().GetDB()
//		var employee model.TblEmployee
//		err := db.First(&employee).Error
//		if err != nil {
//			return nil, err
//		}
//		return employee, nil
//	}
func (er *EmployeeRepositories) GetByID(id uint) (interface{}, error) {
	var employee model.TblEmployee
	err := db.Model(&model.TblEmployee{}).
		Preload("JobType").
		Preload("EmployeeDocs").
		Where("id = ?", id).
		First(&employee).Error
	if err != nil {
		return nil, err
	}
	return employee, nil
}
