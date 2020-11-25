package app

import (
	"github.com/rest_api_master/controllers"
)

func mapUrls() {

	router.GET("/api/ping", controllers.Ping)
	router.GET("/api/allUsers", controllers.GetAllUsers)
	router.POST("/api/createUser", controllers.CreateUser)
	router.PUT("/api/updateUser/:userId", controllers.UpdateUser)
	router.DELETE("/api/deleteUser/:userId", controllers.DeleteUser)
	router.GET("/api/getUserById/:userId", controllers.GetUserById)

}
