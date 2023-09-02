package main

import (
	"github.com/gin-gonic/gin"
	controllers "photoapi.com/ppapi/controllers"
	database "photoapi.com/ppapi/database"
	middleware "photoapi.com/ppapi/middleware"
)

func main() {
	var r = gin.Default()
	database.ConnectDb()
	r.POST("/users/register", controllers.Register_user)
	r.POST("/users/login", controllers.Login_user)
	r.PUT("/users/:userId", middleware.Require_Auth, controllers.Update_user)
	r.DELETE("/users/:userId", middleware.Require_Auth, controllers.Delete_user)

	r.POST("/photos", middleware.Require_Auth, controllers.Create_photo)
	r.GET("/photos", middleware.Require_Auth, controllers.Show_photo)
	r.PUT("/photos/:photoId", middleware.Require_Auth, controllers.Update_photo)
	r.DELETE("/photos/:photoId", middleware.Require_Auth, controllers.Delete_photo)

	r.Run()
}
