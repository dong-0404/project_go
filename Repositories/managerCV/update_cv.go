package managerCV

import "demo_project/model"

func (cv *CVRepositories) Update(id uint, data interface{}) error {
	err := db.Model(&model.TblCV{}).
		Where("id = ?", id).
		Updates(data).Error
	return err
}
