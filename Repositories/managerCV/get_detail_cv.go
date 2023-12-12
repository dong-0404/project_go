package managerCV

import "demo_project/model"

func (cv *CVRepositories) GetDetailsAll() ([]model.TblCV, error) {
	var CVs []model.TblCV

	err := db.Model(&model.TblCV{}).
		Preload("PersonalInfo").
		Preload("PersonalEducation").
		Preload("PersonalProject").
		Find(&CVs).Error
	if err != nil {
		return nil, err
	}
	return CVs, nil
}
func (cv *CVRepositories) Update(id uint, data interface{}) error {
	err := db.Model(&model.TblCV{}).
		Where("id = ?", id).
		Updates(data).Error
	return err
}
