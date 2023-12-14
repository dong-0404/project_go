package managerCV

import "demo_project/model"

func (cv *CVRepositories) GetByID(id uint) (interface{}, error) {
	var CV model.TblCV
	err := db.Model(&model.TblCV{}).
		Preload("JobType").
		Preload("PersonalInfo").
		Preload("PersonalEducation").
		Preload("PersonalProject").
		Where("id = ?", id).
		First(&CV).Error
	if err != nil {
		return nil, err
	}
	return CV, nil
}
