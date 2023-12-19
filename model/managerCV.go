package model

import "gorm.io/gorm"

type TblCV struct {
	*gorm.Model
	FullName string `gorm:"column:full_name" json:"full_name"`
	JobId    int8   `gorm:"column:job_id" json:"job_id"`
	Status   string `gorm:"column:status" json:"status"`
	//JobTypeId int    `gorm:"column:job_id" json:"job_id"`
	JobType           JobType           `gorm:"foreignKey:job_id;references:id"`
	PersonalInfo      PersonalInfo      `gorm:"foreignKey: CvId"`
	PersonalEducation PersonalEducation `gorm:"foreignKey: CvId"`
	PersonalProject   PersonalProject   `gorm:"foreignKey: CvId"`
}

type TblCV1 struct {
	*gorm.Model
	FullName string `gorm:"column:full_name" json:"full_name"`
	JobId    int8   `gorm:"column:job_id" json:"job_id"`
	//JobTypeId int    `gorm:"column:job_id" json:"job_id"`
	//JobType           JobType           `gorm:"foreignKey:job_id;references:id"`
	PersonalInfo      PersonalInfo      `gorm:"foreignKey: CvId"`
	PersonalEducation PersonalEducation `gorm:"foreignKey: CvId"`
	PersonalProject   PersonalProject   `gorm:"foreignKey: CvId"`
}

func (TblCV1) TableName() string {
	return "tbl_cvs"
}
func (TblCV) TableName() string {
	return "tbl_cvs"
}

type PersonalInfo struct {
	*gorm.Model
	CvId    int    `gorm:"column:cv_id" json:"cv_id"`
	Email   string `gorm:"column:email" json:"email"`
	Phone   int64  `gorm:"column:phone" json:"phone"`
	Address string `gorm:"column:address" json:"address"`
	//PE      *PersonalEducation `gorm:"foreignKey: CvId"`
	//PP      *PersonalProject   `gorm:"foreignKey: CvId"`
}

func (PersonalInfo) TableName() string {
	return "personal_infos"
}

type PersonalEducation struct {
	*gorm.Model
	CvId       int    `gorm:"column:cv_id" json:"cv_id"`
	University string `gorm:"column:university" json:"university"`
	Language   string `gorm:"column:language" json:"language"`
	Grade      string `gorm:"column:grade" json:"grade"`
}

func (PersonalEducation) TableName() string {
	return "personal_educations"
}

type PersonalProject struct {
	*gorm.Model
	CvId   int    `gorm:"column:cv_id" json:"cv_id"`
	Skill  string `gorm:"column:skill" json:"skill"`
	Source string `gorm:"column:source" json:"source"`
}

func (PersonalProject) TableName() string {
	return "personal_projects"
}
