package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"photoapi.com/ppapi/app"
	database "photoapi.com/ppapi/database"
	helpers "photoapi.com/ppapi/helpers"
	models "photoapi.com/ppapi/models"
)

func Register_user(c *gin.Context) {
	var user_reg app.AuthRegister
	user_reg.Id = uuid.New().String()
	if err := c.ShouldBindJSON(&user_reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validation
	if err := helpers.Validation(c, user_reg); err != nil {
		return
	}
	//hash password
	hashedPass := helpers.Encrypt_password(c, user_reg.Password)
	if hashedPass == "" {
		return
	}

	user := models.User{
		Id:        user_reg.Id,
		Username:  user_reg.Username,
		Email:     user_reg.Email,
		Password:  hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user. please change the email or the password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully registered",
	})

}

func Login_user(c *gin.Context) {
	var user_login app.AuthLogin
	var user models.User
	//bind user_login to context
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validation
	if err := helpers.Validation(c, user_login); err != nil {
		return
	}
	database.DB.First(&user, "email = ?", user_login.Email)
	//checking invalid email and password
	if user.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}
	//checking password
	if err := helpers.Check_password(user.Password, user_login.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}
	//token
	tokenStr, err := helpers.Initialize_token(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//set token and userid to cookie
	expTime := 60 //seconds
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, expTime, "", "", false, true)
}

func Update_user(c *gin.Context) {
	type Input_user struct {
		Id        string    `valid:"required" json:"id"`
		Username  string    `json:"username"`
		Email     string    `valid:"email" json:"email"`
		Password  string    `valid:"minstringlength(6)" json:"password"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	var input_user Input_user
	var user models.User
	input_user.Id = c.Param("userId")
	input_user.UpdatedAt = time.Now()

	if err := c.ShouldBindJSON(&input_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := helpers.Validation(c, input_user); err != nil {
		return
	}
	if input_user.Password != "" {
		hashedPass := helpers.Encrypt_password(c, input_user.Password)
		if hashedPass == "" {
			return
		}
		input_user.Password = hashedPass
	}
	if database.DB.Model(&user).Where("id = ?", input_user.Id).Updates(&input_user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update. please change the email or the password "})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})
}

func Delete_user(c *gin.Context) {
	var user models.User
	var photo models.Photo
	var deletedPhotos bool
	userid := c.Param("userId")

	if err := database.DB.Where("userid = ?", userid).First(&photo).Error; err == nil {
		database.DB.Delete(&models.Photo{}, "userid = ?", userid)

		deletedPhotos = !deletedPhotos
	}
	if err := database.DB.Where("id = ?", userid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	database.DB.Where("id = ?", userid).Delete(&user)

	if deletedPhotos {
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully deleted photos and account",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})

}
