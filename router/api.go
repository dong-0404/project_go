package router

import (
	"demo_project/controller"
	"demo_project/middleware"
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

	r.POST("/SignUp", controller.SignUp)
	r.POST("/Login", controller.Login)
	Employee := r.Group("/v1/api", middleware.RequireAuth)
	{
		employee := Employee.Group("employee")
		{
			employee.GET("", controller.GetEmployees)
			employee.POST("/create", controller.CreateEmployee)
			employee.GET("/:id", controller.GetEmployeeByID)
			employee.POST("/update/:id", controller.Update)
			employee.DELETE("/delete/:id", controller.Delete)
		}
	}
	ManagerCv := r.Group("/v1/api")
	{
		Cv := ManagerCv.Group("cv")
		{
			Cv.GET("", controller.GetCVs)
			Cv.POST("/create", controller.CreateCV)
			Cv.GET("/:id", controller.GetCVByID)
			Cv.POST("/update/:id", controller.UpdateCv)
			Cv.DELETE("/delete/:id", controller.DeleteCv)
		}
	}
	return r
}
