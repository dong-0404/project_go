package model

import (
	"gorm.io/gorm"
	"time"
)

type TblEmployee struct {
	*gorm.Model
	//Id        int        `gorm:"column:id;autoIncrement" json:"id"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    int32  `gorm:"column:phone" json:"phone"`
	Address  string `gorm:"column:address" json:"address"`
	JobType  int    `gorm:"column:job_type" json:"job_type"`
	StatusId int    `gorm:"column:status_id" json:"status_id"`
	//CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (TblEmployee) TableName() string {
	return "tbl_employee"
}

type EmployeeDocs struct {
	Id              int         `gorm:"column:id;autoIncrement" json:"id"`
	EmployeeId1     int         `gorm:"column:employee_id1" json:"employee_id1"`
	IdCard          int32       `gorm:"column:id_card" json:"id_card"`
	JoiningDate     time.Time   `gorm:"column:joining_date" json:"joining_date"`
	BankAccount     int64       `gorm:"column:bank_account" json:"bank_account"`
	BankAccountInfo string      `gorm:"column:bank_account_info" json:"bank_account_info"`
	TblEmployee     TblEmployee `gorm:"foreignKey:EmployeeId1"`
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
