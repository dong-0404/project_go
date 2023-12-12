package employee

import (
	"demo_project/database"
)

type EmployeeRepositories struct{}

var db = database.GetInstance().GetDB()

func (er *EmployeeRepositories) Create(data interface{}) error {
	err := db.Create(data).Error
	return err
}

//type CreateEmployee struct {
//	FullName string `gorm:"column:full_name" json:"full_name"`
//	JobType  int    `gorm:"column:job_type" json:"job_type"`
//	StatusId int    `gorm:"column:status_id" json:"status_id"`
//	Email    string `gorm:"column:email" json:"email"`
//	Phone    int32  `gorm:"column:phone" json:"phone"`
//	Address  string `gorm:"column:address" json:"address"`
//}

//func (CreateEmployee) TableName() string {
//	return "tbl_employees"
//}
