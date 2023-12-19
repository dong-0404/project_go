package managerCV

import "demo_project/model"

func (cv *CVRepositories) AccessCv(id uint) error {
	updates := map[string]interface{}{"status": "Access"}

	err := db.Model(&model.TblCV{}).Where("id = ?", id).Updates(updates).Error

	return err
}
