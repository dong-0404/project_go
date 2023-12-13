package router

import (
	"demo_project/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Adjust to the origins you want to allow
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.POST("/SignUp", controller.SignUp)
	r.POST("/Login", controller.Login)
	Employee := r.Group("/v1/api")
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
