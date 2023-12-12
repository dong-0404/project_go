package model

import (
	"gorm.io/gorm"
)

type TblEmployee struct {
	*gorm.Model
	//Id        int        `gorm:"column:id;autoIncrement" json:"id"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    int32  `gorm:"column:phone" json:"phone"`
	Address  string `gorm:"column:address" json:"address"`
	JobType  int    `gorm:"column:job_type" json:"job_type"`
	StatusId string `gorm:"column:status_id" json:"status_id"`
	//CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	EmployeeDocs *EmployeeDocs `gorm:"foreignKey: EmployeeId"`
}

func (TblEmployee) TableName() string {
	return "tbl_employee"
}

type EmployeeDocs struct {
	*gorm.Model // thay cho truong id,created_at,updated_at,deleted_at
	//ID              int         `gorm:"column:id;autoIncrement" json:"id"`
	EmployeeId      int    `gorm:"column:employee_id" json:"employee_id"`
	IdCard          int32  `gorm:"column:id_card" json:"id_card"`
	JoiningDate     string `gorm:"column:joining_date" json:"joining_date"`
	BankAccount     int64  `gorm:"column:bank_account" json:"bank_account"`
	BankAccountInfo string `gorm:"column:bank_account_info" json:"bank_account_info"`
}

func (EmployeeDocs) TableName() string {
	return "Employee_docs"
}

type jobType struct {
	Int  int    `json:"id"`
	Name string `json:"name"`
}

type statusJob struct {
	Int  int    `json:"id"`
	Name string `json:"name"`
}
