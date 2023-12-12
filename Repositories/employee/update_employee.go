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
