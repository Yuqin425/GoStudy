package api

import (
	"goPractice/Week_08/api/middleware"
	"goPractice/Week_08/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.Run(":8080")
}
