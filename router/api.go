package router

import (
	"demo_project/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	client := r.Group("/api")
	{
		employee := client.Group("employee")
		{
			employee.GET("", controller.GetEmployees)
			employee.POST("/create", controller.CreateEmployee)
			employee.GET("/:id", controller.GetEmployeeByID)
			employee.POST("/update/:id", controller.Update)
			employee.DELETE("/delete/:id", controller.Delete)
		}
	}
	return r
}
