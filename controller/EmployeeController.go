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
	if err := c.Bind(&dataInsert); err != nil {
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
	employee, err := ec.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.Bind(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = ec.Update(uint(id), &employee)

	c.JSON(http.StatusOK, gin.H{
		"message": true,
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
