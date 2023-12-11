package controller

import (
	"demo_project/Repositories/employee"
	"demo_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

func CreateEmployee(c *gin.Context) {
	var dataInsert model.TblEmployee
	if err := c.ShouldBindJSON(&dataInsert); err != nil {
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
