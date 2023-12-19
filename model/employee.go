package model

import (
	"gorm.io/gorm"
)

type TblEmployee struct {
	*gorm.Model
	//Id        int        `gorm:"column:id;autoIncrement" json:"id"`
	FullName     string        `gorm:"column:full_name" json:"full_name"`
	Email        string        `gorm:"column:email" json:"email"`
	Phone        int           `gorm:"column:phone" json:"phone"`
	Address      string        `gorm:"column:address" json:"address"`
	JobTypeId    int           `gorm:"column:job_type" json:"job_type"`
	JobType      JobType       `gorm:"foreignKey:job_type;references:id"`
	EmployeeDocs *EmployeeDocs `gorm:"foreignKey: EmployeeId"`
}

func (TblEmployee) TableName() string {
	return "tbl_employee"
}

type TblEmployee1 struct {
	*gorm.Model
	//Id        int        `gorm:"column:id;autoIncrement" json:"id"`
	FullName  string  `gorm:"column:full_name" json:"full_name"`
	Email     string  `gorm:"column:email" json:"email"`
	Phone     int     `gorm:"column:phone" json:"phone"`
	Address   string  `gorm:"column:address" json:"address"`
	JobTypeId int     `gorm:"column:job_type" json:"job_type"`
	JobType   JobType `gorm:"foreignKey:job_type;references:id"`
	StatusId  string  `gorm:"column:status_id" json:"status_id"`
	//CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	EmployeeDocs *EmployeeDocs `gorm:"foreignKey: employee_id"`
}

func (TblEmployee1) TableName() string {
	return "tbl_employee"
}

type JobType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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

func (JobType) TableName() string {
	return "tbl_job_type"
}

//type statusJob struct {
//	Int  int    `json:"id"`
//	Name string `json:"name"`
//}
