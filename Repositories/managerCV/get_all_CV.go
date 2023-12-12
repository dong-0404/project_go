package managerCV

import (
	"gorm.io/gorm"
)

type TblCV struct {
	*gorm.Model
	FullName string `gorm:"column:full_name" json:"full_name"`
	JobId    int8   `gorm:"column:job_id" json:"job_id"`
	Status   string `gorm:"column:status" json:"status"`
	//PersonalInfo      PersonalInfo      `gorm:"foreignKey: CvId"`
	//PersonalEducation PersonalEducation `gorm:"foreignKey: CvId"`
	//PersonalProject   PersonalProject   `gorm:"foreignKey: CvId"`
}

func (TblCV) TableName() string {
	return "tbl_cvs"
}

func (cv *CVRepositories) GetAll() ([]TblCV, error) {
	var CVs []TblCV

	err := db.Model(&TblCV{}).Find(&CVs).Error
	if err != nil {
		return nil, err
	}
	return CVs, nil
}
