package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hagios2/gopharma/controllers"
)

func Routes(route *gin.Engine) {
	authRoutes := route.Group("/auth")
	authRoutes.POST("/login", controllers.Login)
	authRoutes.POST("/logout", controllers.Login)

}
