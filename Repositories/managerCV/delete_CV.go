package managerCV

import "demo_project/model"

func (cv *CVRepositories) Delete(id uint) error {
	err := cv.UpdateStatus(id)
	if err != nil {
		return err
	}
	var CV model.TblCV
	err = db.Where("id=?", id).Delete(&CV).Error
	return err
}

func (cv *CVRepositories) UpdateStatus(id uint) error {
	updates := map[string]interface{}{"status": "Deleted"}

	err := db.Model(&model.TblCV{}).Where("id=?", id).Updates(updates).Error
	return err
}
