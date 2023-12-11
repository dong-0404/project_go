package employee

import (
	"demo_project/database"
	"time"
)

type EmployeeRepositories struct{}

func (cre *EmployeeRepositories) Create(data interface{}) error {
	db := database.GetInstance().GetDB()
	err := db.Create(data).Error
	return err
}

type CreateEmployee struct {
	FullName        string     `gorm:"column:full_name" json:"full_name"`
	JobType         int        `gorm:"column:job_type" json:"job_type"`
	StatusId        int        `gorm:"column:status_id" json:"status_id"`
	Email           string     `gorm:"column:email" json:"email"`
	Phone           int32      `gorm:"column:phone" json:"phone"`
	Address         string     `gorm:"column:address" json:"address"`
	IdCard          int32      `gorm:"column:id_card" json:"id_card"`
	JoiningDate     *time.Time `gorm:"column:joining_date" json:"joining_date"`
	BankAccount     int64      `gorm:"column:bank_account" json:"bank_account"`
	BankAccountInfo string     `gorm:"column:bank_account_info" json:"bank_account_info"`
}

func (CreateEmployee) TableName() string {
	return "tbl_employees"
}
