package employee

import (
	"demo_project/model"
)

//type GetEmployee struct {
//	*gorm.Model
//	//Id       int    `gorm:"column:id;autoIncrement" json:"id"`
//	FullName string `gorm:"column:full_name" json:"full_name"`
//	Email    string `gorm:"column:email" json:"email"`
//	Phone    int32  `gorm:"column:phone" json:"phone"`
//	Address  string `gorm:"column:address" json:"address"`
//	JobType  int    `gorm:"column:job_type" json:"job_type"`
//	StatusId string `gorm:"column:status_id" json:"status_id"`
//	//CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
//	//UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
//	//EmployeeDocs *EmployeeDocs `gorm:"foreignKey: EmployeeId"`
//}

//func (GetEmployee) TableName() string {
//	return "tbl_employee"
//}

func (er *EmployeeRepositories) GetAll() ([]model.TblEmployee1, error) {
	var employees []model.TblEmployee1
	err := db.Model(&model.TblEmployee1{}).Preload("JobType").Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}
