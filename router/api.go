package router

import (
	"demo_project/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS	"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Xử lý yêu cầu OPTIONS
	r.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	public := r.Group("/public")
	{
		public.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	v1api := r.Group("/v1/api")
	{
		employee := v1api.Group("/employee")
		{
			employee.GET("", controller.GetEmployees)
			employee.POST("/create", controller.CreateEmployee)
			employee.GET("/:id", controller.GetEmployeeByID)
			employee.POST("/update/:id", controller.Update)
			employee.DELETE("/delete/:id", controller.Delete)
		}

		cv := v1api.Group("/cv")
		{
			cv.GET("", controller.GetCVs)
			cv.POST("/create", controller.CreateCV)
			cv.GET("/:id", controller.GetCVByID)
			cv.POST("/update/:id", controller.UpdateCv)
			cv.DELETE("/delete/:id", controller.DeleteCv)
		}
	}

	return r

}
