package employee

import (
	"demo_project/database"
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
	db := database.GetInstance().GetDB()
	var employee model.TblEmployee
	err := db.First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return employee, nil
}