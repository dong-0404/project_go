package employee

import (
	"demo_project/database"
	"demo_project/model"
)

func (er *EmployeeRepositories) Update(id uint, data interface{}) error {
	db := database.GetInstance().GetDB()
	err := db.Model(&model.TblEmployee{}).Where("id = ?", id).Updates(data).Error
	return err
}
