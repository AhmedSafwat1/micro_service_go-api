package app

import "github.com/AhmedSafwat1/micro_service_go-api/controllers"

func mapUrl()  {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)

	//router.GET("/users/search", controllers.Search)

	router.GET("/users/:user_id", controllers.GetUser)

}
