package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "photoapi.com/ppapi/database"
	"photoapi.com/ppapi/helpers"
	models "photoapi.com/ppapi/models"
)

func Create_photo(c *gin.Context) {
	var photo models.Photo
	photo.Id = uuid.New().String()
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messsage": err.Error()})
		return
	}

	if database.DB.Create(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to add item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"photo": photo})

}

func Show_photo(c *gin.Context) {
	var photos []models.Photo
	userid, _ := c.Get("userid")

	database.DB.Where("userid = ?", userid).Find(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func Update_photo(c *gin.Context) {
	var photo models.Photo
	photo.Id = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&photo).Where("userid = ?", photo.Userid).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data has been updated"})
}
func Delete_photo(c *gin.Context) {
	var photo models.Photo
	photo.Id = c.Param("photoId")
	userid, _ := c.Get("userid")
	photo.Userid = userid.(string)

	if err := helpers.Validation(c, photo); err != nil {
		return
	}
	if err := database.DB.Where("userid = ?", photo.Userid).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Where("userid = ?", photo.Userid).Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
