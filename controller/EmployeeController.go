package controller

import (
	"demo_project/Repositories/employee"
	"demo_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ec = employee.EmployeeRepositories{}

func GetEmployees(c *gin.Context) {
	employees, err := ec.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var dataInsert model.TblEmployee
	if err := c.BindJSON(&dataInsert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err := ec.Create(&dataInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dataInsert)
}

func GetEmployeeByID(c *gin.Context) {
	// Params.ByName ("Truy xuất tham số URL theo tên")
	// strconv.Atoi (chuyển từ string->int)
	id, err := strconv.Atoi(c.Params.ByName("id"))
	employee, err := ec.GetByID(uint(id))
	//employee, err := ec.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	type dataUpdate struct {
		dataEmployee     model.TblEmployee  `json:"dataEmployee"`
		dataEmployeeDocs model.EmployeeDocs `json:"dataEmployeeDocs"`
	}

	var UpdateEmployee dataUpdate = dataUpdate{}

	if err := c.ShouldBindJSON(&UpdateEmployee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin cho bảng tbl_employee
	//requestData.EmployeeInfo.ID = uint(id)
	err = ec.UpdateEmployeeInfo(id, &UpdateEmployee.dataEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cập nhật thông tin cho bảng employeeDocs
	err = ec.UpdateEmployeeDocs(id, &UpdateEmployee.dataEmployeeDocs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    UpdateEmployee,
	})
}

//var employee model.TblEmployee

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ec.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
