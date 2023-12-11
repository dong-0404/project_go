package employee

import (
	"demo_project/database"
	"demo_project/model"
)

func (er *EmployeeRepositories) Delete(id uint) error {
	// cập nhật trạng thái trong db
	err := er.Delete1(id)
	if err != nil {
		return err
	}

	db := database.GetInstance().GetDB()

	// Xoá bản ghi
	var employee model.TblEmployee
	err = db.Where("id = ?", id).Delete(&employee).Error
	return err
}

func (er *EmployeeRepositories) Delete1(id uint) error {
	db := database.GetInstance().GetDB()

	// Sử dụng map để chỉ cập nhật cột 'status_id'
	updates := map[string]interface{}{"status_id": "Deleted"}

	// Sử dụng phương thức Updates để thực hiện truy vấn UPDATE
	err := db.Model(&model.TblEmployee{}).Where("id = ?", id).Updates(updates).Error

	return err
}
