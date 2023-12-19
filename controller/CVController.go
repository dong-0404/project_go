package controller

import (
	"demo_project/Repositories/managerCV"
	"demo_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var cv = managerCV.CVRepositories{}

func CreateCV(c *gin.Context) {
	var dataInsert model.TblCV1
	if err := c.Bind(&dataInsert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := cv.Create(&dataInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dataInsert)
}

func GetCVs(c *gin.Context) {
	CVs, err := cv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CVs)
}

func GetCVByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	cv, err := cv.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": cv,
	})
}

func UpdateCv(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	var CvUpdate model.TblCV1
	if err := c.ShouldBindJSON(&CvUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = cv.Update(uint(id), &CvUpdate)

	c.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    CvUpdate,
	})
}

func DeleteCv(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = cv.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
func AccessCv(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = cv.AccessCv(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"message": "Access successfully",
	})
}
