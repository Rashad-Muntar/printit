package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rashad-muntar/printit/controllers"
)

func AuthRoute(router *gin.Engine) {
	router.GET("/auth/login", controllers.Login)
	router.GET("/auth/register", controllers.Register)
	}
