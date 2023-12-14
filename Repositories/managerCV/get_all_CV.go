package managerCV

import (
	"demo_project/model"
)

func (cv *CVRepositories) GetAll() ([]model.TblCV, error) {
	var CVs []model.TblCV

	err := db.Model(&model.TblCV{}).
		Preload("PersonalInfo").
		Preload("JobType").
		Preload("PersonalEducation").
		Preload("PersonalProject").
		Find(&CVs).Error
	if err != nil {
		return nil, err
	}
	return CVs, nil
}
